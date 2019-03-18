/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"context"

	"github.com/vmware/govmomi/simulator"
)

//CreateVcenterMockServer creates a mock server and returns a mock server and model used
func CreateVcenterMockServer() (*simulator.Model, *simulator.Server, error) {
	// create sim
	m := simulator.VPX()
	err := m.Create()
	// create server
	server := m.Service.NewServer()
	return m, server, err
}

//CreateContext creates a context with background and returns the context and cancel function
func CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	return ctx, cancel
}

//CreateVcenterClient creates a vcenter client connected to VCURL
func CreateVcenterClient(ctx context.Context, VCURL string) (*Client, error) {
	// create client with context
	vcenterClient, err := NewClient(ctx, VCURL, true)
	return vcenterClient, err
}

//CancelContext cancels the context using the passed in cancel function
func CancelContext(cancel context.CancelFunc) {
	if cancel != nil {
		cancel()
	}
}

//DestroyVcenterMockServer closes the mock server and also removes the simulator model
func DestroyVcenterMockServer(model *simulator.Model, server *simulator.Server) {
	if server != nil {
		server.Close()
	}

	if model != nil {
		model.Remove()
	}
}
