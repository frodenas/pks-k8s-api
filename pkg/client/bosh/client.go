/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package bosh

import (
	"fmt"

	"github.com/cloudfoundry/bosh-cli/director"
	"github.com/cloudfoundry/bosh-cli/uaa"
	"github.com/cloudfoundry/bosh-utils/logger"
)

type client struct {
	boshLogger     logger.Logger
	directorConfig director.FactoryConfig
	boshFactory    director.Factory
}

// NewClient returns a new BOSH client given a BOSH url, clientID, clientSecret, and CA certificate.
func NewClient(url string, clientID string, clientSecret string, caCert string) (Client, error) {
	boshLogger := logger.NewLogger(logger.LevelInfo)

	directorConfig, err := director.NewConfigFromURL(url)
	if err != nil {
		return nil, err
	}
	directorConfig.CACert = caCert

	anonymousDirector, err := director.NewFactory(boshLogger).New(directorConfig, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	boshInfo, err := anonymousDirector.Info()
	if err != nil {
		return nil, err
	}

	uaaURL := boshInfo.Auth.Options["url"]
	uaaURLStr, ok := uaaURL.(string)
	if !ok {
		return nil, fmt.Errorf("Expected UAA URL '%s' to be a string", uaaURL)
	}

	uaaConfig, err := uaa.NewConfigFromURL(uaaURLStr)
	if err != nil {
		return nil, err
	}

	uaaConfig.CACert = caCert
	uaaConfig.Client = clientID
	uaaConfig.ClientSecret = clientSecret

	uaaFactory := uaa.NewFactory(boshLogger)
	uaaClient, err := uaaFactory.New(uaaConfig)
	if err != nil {
		return nil, err
	}

	directorConfig.TokenFunc = uaa.NewClientTokenSession(uaaClient).TokenFunc
	boshFactory := director.NewFactory(boshLogger)

	return &client{
		boshLogger:     boshLogger,
		directorConfig: directorConfig,
		boshFactory:    boshFactory,
	}, nil
}

// Director returns a BOSH Director.
func (bc *client) Director(taskReporter director.TaskReporter) (director.Director, error) {
	return bc.boshFactory.New(bc.directorConfig, nil, taskReporter, director.NewNoopFileReporter())
}
