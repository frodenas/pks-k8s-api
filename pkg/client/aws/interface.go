/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package aws

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Client represents an AWS Client.
//go:generate moq -out fakes/client.go -pkg fakes . Client
type Client interface {
	// GetAvailabilityZone gets an Availability Zone object given a zone name.
	GetAvailabilityZone(zone string) (*ec2.DescribeAvailabilityZonesOutput, error)

	// GetVPC gets an VPC object given a vpc id.
	GetVPC(vpcID string) (*ec2.DescribeVpcsOutput, error)
}
