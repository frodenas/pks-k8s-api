/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package nsxt

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/netprovisioner"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("client.nsxt")

type client struct {
	nsxtClient      nsx.Client
	nsxtProvisioner netprovisioner.NsxNetworkProvisioner
}

// NewClient returns a new NSX-T client given a NSX-T manager host, user, and password.
func NewClient(host string, user string, password string, insecure bool) (Client, error) {
	nsxtClient := nsx.NewClient(host).WithBasicAuth(user, password).WithInsecure(insecure)
	if err := nsxtClient.Validate(); err != nil {
		return nil, fmt.Errorf("error creating a NSX-T client: %v", err)
	}

	nsxtProvisioner, err := netprovisioner.NewNsxNetworkProvisioner(nsxtClient, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating a NSX-T provisioner: %v", err)
	}

	return &client{
		nsxtClient:      nsxtClient,
		nsxtProvisioner: nsxtProvisioner,
	}, nil
}
