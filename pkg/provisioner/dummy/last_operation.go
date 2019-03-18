/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package dummyprovisioner

import (
	"fmt"
	"strconv"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	provisionertypes "github.com/frodenas/pks-k8s-api/pkg/provisioner/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// NumIterationsToSuceed is the number of iterations for an operation to consider it suceed.
	NumIterationsToSuceed = 10
)

// LastOperation polls the last operation for a Dummy Cluster.
func (p *Provisioner) LastOperation(instance *pksv1alpha1.Cluster, provisionerLastOperation provisionertypes.ProvisionerLastOperation) (*provisionertypes.ProvisionerLastOperation, error) {
	log.Info(fmt.Sprintf("Polling Last Operation fo Dummy Cluster `%s/%s`", instance.Namespace, instance.Name))

	iteration, err := strconv.Atoi(provisionerLastOperation.ProvisionerData)
	if err != nil {
		return nil, fmt.Errorf("error getting iteration from provisioner last operation data for Dummy Cluster `%s/%s`: %v", instance.Namespace, instance.Name, err)
	}

	iteration = iteration + 1

	newProvisionerLastOperation := &provisionertypes.ProvisionerLastOperation{
		Description:     "operation in progress",
		LastUpdated:     metav1.NewTime(time.Now()),
		State:           provisionertypes.ProvisionerOperationStateInProgress,
		Type:            provisionerLastOperation.Type,
		ProvisionerData: strconv.Itoa(iteration),
	}

	if iteration > NumIterationsToSuceed {
		newProvisionerLastOperation.Description = "operation succeeded"
		newProvisionerLastOperation.State = provisionertypes.ProvisionerOperationStateSucceeded
		newProvisionerLastOperation.ProvisionerData = ""
	}

	return newProvisionerLastOperation, nil
}
