/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"context"
	"errors"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/soap"
)

// Client is a client to communicate with a given vcenter server
type Client struct {
	url      string
	insecure bool
	client   *govmomi.Client
	context  context.Context
	finder   *find.Finder
}

// NewClient returns a new vcenter client given a vcenter server URL and a context
func NewClient(ctx context.Context, url string, insecure bool) (*Client, error) {
	if url == "" {
		return nil, errors.New("vcenter url is empty")
	}

	if ctx == nil {
		return nil, errors.New("context is nil")
	}

	parsedurl, err := soap.ParseURL(url)
	if err != nil {
		return nil, err
	}

	govmomiClient, govmomiErr := govmomi.NewClient(ctx, parsedurl, insecure)
	if govmomiErr != nil {
		return nil, govmomiErr
	}

	vcenterAPIFinder := find.NewFinder(govmomiClient.Client, true)
	vcenterclient := &Client{url: url, insecure: insecure, client: govmomiClient,
		context: ctx, finder: vcenterAPIFinder}

	return vcenterclient, nil
}

// GetContext returns the context associated with the vcenter client
func (vc *Client) GetContext() context.Context {
	return vc.context
}
