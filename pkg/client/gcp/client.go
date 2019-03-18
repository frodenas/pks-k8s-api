/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package gcp

import (
	"fmt"

	"golang.org/x/oauth2"
	oauthgoogle "golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/storage/v1"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("client.gcp")

const (
	computeScope = compute.ComputeScope
	storageScope = storage.DevstorageFullControlScope
)

type client struct {
	project string

	computeService *compute.Service
	storageService *storage.Service
}

// NewClient returns a new GCP client given a GCP project, and a JSONKey.
func NewClient(project string, jsonKey string) (Client, error) {
	computeJwtConf, err := oauthgoogle.JWTConfigFromJSON([]byte(jsonKey), computeScope)
	if err != nil {
		return nil, fmt.Errorf("error reading Google JSON Key: %v", err)
	}
	computeClient := computeJwtConf.Client(oauth2.NoContext)

	storageJwtConf, err := oauthgoogle.JWTConfigFromJSON([]byte(jsonKey), storageScope)
	if err != nil {
		return nil, fmt.Errorf("error reading Google JSON Key: %v", err)
	}
	storageClient := storageJwtConf.Client(oauth2.NoContext)

	computeService, err := compute.New(computeClient)
	if err != nil {
		return nil, fmt.Errorf("error creating a Google Compute Service client: %v", err)
	}

	storageService, err := storage.New(storageClient)
	if err != nil {
		return nil, fmt.Errorf("error creating a Google Storage Service client: %v", err)
	}

	return &client{
		project:        project,
		computeService: computeService,
		storageService: storageService,
	}, nil
}
