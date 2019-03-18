/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// GetVPC gets an VPC object given a vpc id.
func (ac *client) GetVPC(vpcID string) (*ec2.DescribeVpcsOutput, error) {
	log.Info(fmt.Sprintf("Getting VPC with ID `%s`", vpcID))

	az, err := ac.ec2Service.DescribeVpcsRequest(&ec2.DescribeVpcsInput{VpcIds: []string{vpcID}}).Send()
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting VPC with ID `%s`", vpcID))
		return nil, err
	}

	return az, nil
}
