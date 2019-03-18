/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// GetAvailabilityZone gets an Availability Zone object given a zone name.
func (ac *client) GetAvailabilityZone(zone string) (*ec2.DescribeAvailabilityZonesOutput, error) {
	log.Info(fmt.Sprintf("Getting Zone `%s`", zone))

	az, err := ac.ec2Service.DescribeAvailabilityZonesRequest(&ec2.DescribeAvailabilityZonesInput{
		ZoneNames: []string{zone},
	}).Send()
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting Zone `%s`", zone))
		return nil, err
	}

	return az, nil
}
