/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"path/filepath"

	"github.com/go-openapi/runtime"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/printer"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/util"
)

// client represents nsx client with its metadata
type client struct {
	info      *nsxManagerInfo
	transport *rc.Runtime
	client    *nsx.Nsx
	auth      runtime.ClientAuthInfoWriter
	*printer.Printer
}

// nsxManagerInfo represents info of nsx manager
type nsxManagerInfo struct {
	target             string
	username           string
	password           string
	insecureSkipVerify bool
	rootCA             string
	clientCert         string
	clientKey          string
}

// NewClient provides new nsx client with configuration
func NewClient(target string) Client {

	transport := rc.New(target, basePath, scheme)
	transport.Transport = NsxKeepAliveTransport(http.DefaultTransport)
	t := transport.Transport.(*nsxKeepAliveTransport).Transport.(*http.Transport)
	t.TLSClientConfig = &tls.Config{
		ServerName: target,
	}

	transport.Consumers[contentType] = runtime.JSONConsumer()
	return &client{
		info: &nsxManagerInfo{
			target: target,
		},
		transport: transport,
		client:    nsx.New(transport, strfmt.Default),
		// prints to stderr by default
		Printer: printer.New(os.Stderr),
	}
}

// WithBasicAuth adds basic authentication setup to nsx client
func (nc *client) WithBasicAuth(username, password string) Client {
	if username != "" {
		nc.auth = rc.BasicAuth(username, password)
	}
	nc.info.username = username
	nc.info.password = password

	return nc
}

// WithInsecure sets insecure flag of nsx client
func (nc *client) WithInsecure(insecure bool) Client {
	t := nc.transport.Transport.(*nsxKeepAliveTransport).Transport.(*http.Transport)
	t.TLSClientConfig.InsecureSkipVerify = insecure
	nc.info.insecureSkipVerify = insecure

	return nc
}

// WithDebug sets debug
func (nc *client) WithDebug() Client {
	nc.transport.SetDebug(true)
	return nc
}

// WithRootCA adds root certificate to nsx client, this requires insecure flag
// to be true
func (nc *client) WithRootCA(rootCA string) Client {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootCA))
	if !ok {
		return nc
	}

	t := nc.transport.Transport.(*nsxKeepAliveTransport).Transport.(*http.Transport)
	t.TLSClientConfig.RootCAs = roots
	nc.info.rootCA = rootCA
	return nc
}

// WithClientCert adds client certificate to nsx client
func (nc *client) WithClientCert(cert, key string) Client {
	nc.info.clientCert = cert
	nc.info.clientKey = key
	clientCert, err := tls.X509KeyPair([]byte(cert), []byte(key))
	if err != nil {
		return nc
	}
	t := nc.transport.Transport.(*nsxKeepAliveTransport).Transport.(*http.Transport)
	t.TLSClientConfig.GetClientCertificate = func(info *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		return &clientCert, nil
	}
	t.TLSClientConfig.BuildNameToCertificate()
	return nc
}

// WithClientCertFromFile adds client certificate files to nsx client
func (nc *client) WithClientCertFromFile(certFile, keyFile string) Client {
	cert, err := ioutil.ReadFile(filepath.Clean(certFile))
	if err != nil {
		return nc
	}
	key, err := ioutil.ReadFile(filepath.Clean(keyFile))
	if err != nil {
		return nc
	}
	return nc.WithClientCert(string(cert), string(key))
}

// WithOverWriteHeader sets global headers for all the requests
func (nc *client) WithOverWriteHeader() Client {
	t := nc.transport.Transport.(*nsxKeepAliveTransport)
	if t.header == nil {
		t.header = make(http.Header)
	}
	t.header.Set("X-Allow-Overwrite", "true")
	return nc
}

// Validate validates if nsx client doesnt have invalid information
func (nc *client) Validate() error {
	err := util.EnsureParams(nc.info.target)
	if err != nil {
		return err
	}

	if util.EnsureParams(nc.info.username, nc.info.password) != nil && util.EnsureParams(nc.info.clientCert, nc.info.clientKey) != nil {
		return fmt.Errorf("Client using neither basic auth nor client certificate")
	}

	if !nc.info.insecureSkipVerify && nc.info.rootCA == "" {
		return fmt.Errorf("False insecure flag must include valid certificate")
	}
	return nil
}

// type mapping function
func GetManagedResource(i interface{}) (*models.ManagedResource, error) {
	var m *models.ManagedResource
	switch obj := i.(type) {
	case *models.FirewallSection:
		m = &obj.ManagedResource
	case *models.NSGroup:
		m = &obj.ManagedResource
	case *models.LbService:
		m = &obj.ManagedResource
	case *models.LbRule:
		m = &obj.ManagedResource
	case *models.LbPool:
		m = &obj.ManagedResource
	case *models.LogicalPort:
		m = &obj.ManagedResource
	case *models.LogicalRouterPort:
		m = &obj.ManagedResource
	case *models.LbVirtualServer:
		m = &obj.ManagedResource
	case *models.LogicalRouter:
		m = &obj.ManagedResource
	case *models.LogicalSwitch:
		m = &obj.ManagedResource
	case *models.IPPool:
		m = &obj.ManagedResource
	case *models.NatRule:
		m = &obj.ManagedResource
	case *models.IPSet:
		m = &obj.ManagedResource
	case *models.LbAppProfile:
		m = &obj.ManagedResource
	case *models.BaseSwitchingProfile:
		m = &obj.ManagedResource
	case *models.Certificate:
		m = &obj.ManagedResource
	case *models.LbPersistenceProfile:
		m = &obj.ManagedResource
	case *models.ManagedResource:
		m = obj
	default:
		return nil, errors.New(fmt.Sprintf("GetManagedResource(): unrecognized type %s", obj))
	}
	return m, nil
}
