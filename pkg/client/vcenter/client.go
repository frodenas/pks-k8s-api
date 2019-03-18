/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"context"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/soap"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("client.vcenter")

type client struct {
	context context.Context

	client *govmomi.Client
	finder *find.Finder
}

// NewClient returns a new vCenter client given a vCenter server URL, user, password, and a context.
func NewClient(ctx context.Context, server string, user string, password string, insecure bool) (Client, error) {
	soapURL, err := soap.ParseURL(server)
	if err != nil {
		return nil, err
	}

	soapURL.User = url.UserPassword(user, password)
	vcClient, err := govmomi.NewClient(ctx, soapURL, insecure)
	if err != nil {
		return nil, err
	}

	finder := find.NewFinder(vcClient.Client, true)

	return &client{
		context: ctx,
		client:  vcClient,
		finder:  finder,
	}, nil
}
