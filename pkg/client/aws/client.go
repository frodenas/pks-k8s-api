/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("client.aws")

type client struct {
	region string

	ec2Service *ec2.EC2
}

// NewClient returns a new AWS client given an AWS access key, a secret access key, and a region.
func NewClient(accessKey string, secretAccessKey string, region string) (Client, error) {
	cfg, err := external.LoadDefaultAWSConfig(
		external.WithCredentialsValue(aws.Credentials{AccessKeyID: accessKey, SecretAccessKey: secretAccessKey}),
		external.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating an AWS client: %v", err)
	}

	ec2Service := ec2.New(cfg)

	return &client{
		region:     region,
		ec2Service: ec2Service,
	}, nil
}
