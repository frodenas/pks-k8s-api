/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package dummyprovisioner

import (
	"fmt"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	provisionertypes "github.com/frodenas/pks-k8s-api/pkg/provisioner/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateCluster creates a Dummy Cluster.
func (p *Provisioner) CreateCluster(instance *pksv1alpha1.Cluster) (*provisionertypes.ProvisionerLastOperation, error) {
	log.Info(fmt.Sprintf("Creating Dummy Cluster `%s/%s`", instance.Namespace, instance.Name))

	provisionerLastOperation := &provisionertypes.ProvisionerLastOperation{
		Description:     "operation in progress",
		LastUpdated:     metav1.NewTime(time.Now()),
		State:           provisionertypes.ProvisionerOperationStateInProgress,
		Type:            provisionertypes.ProvisionerOperationTypeCreate,
		ProvisionerData: "1",
	}

	return provisionerLastOperation, nil
}
