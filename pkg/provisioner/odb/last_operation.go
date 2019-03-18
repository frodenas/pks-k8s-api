/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package odbprovisioner

import (
	"encoding/json"
	"fmt"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/frodenas/pks-k8s-api/pkg/provisioner/odb/utils"
	provisionertypes "github.com/frodenas/pks-k8s-api/pkg/provisioner/types"
	osb "github.com/maplain/go-open-service-broker-client/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LastOperation polls the last operation for an On-Demand-Broker Cluster.
func (p *Provisioner) LastOperation(instance *pksv1alpha1.Cluster, provisionerLastOperation provisionertypes.ProvisionerLastOperation) (*provisionertypes.ProvisionerLastOperation, error) {
	serviceInstanceName := utils.ServiceInstanceName(instance.Namespace, instance.Name)
	log.Info(fmt.Sprintf("Polling Last Operation for On-Demand-Broker Cluster `%s`", serviceInstanceName))

	// Poll last operation.
	var operationKey osb.OperationKey
	err := json.Unmarshal([]byte(provisionerLastOperation.ProvisionerData), &operationKey)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling provisioner last operation data for On-Demand-Broker Cluster `%s`: %v", serviceInstanceName, err)
	}

	lastOperationRequest := &osb.LastOperationRequest{
		InstanceID:   serviceInstanceName,
		OperationKey: &operationKey,
	}

	lastOperationResponse, err := p.osbClient.PollLastOperation(lastOperationRequest)
	if err != nil {
		return nil, fmt.Errorf("error polling last operation for On-Demand-Broker Cluster `%s`: %v", serviceInstanceName, err)
	}

	// Return last operation.
	newProvisionerLastOperation := &provisionertypes.ProvisionerLastOperation{
		Description:     *lastOperationResponse.Description,
		LastUpdated:     metav1.NewTime(time.Now()),
		Type:            provisionerLastOperation.Type,
		ProvisionerData: provisionerLastOperation.ProvisionerData,
	}

	switch lastOperationResponse.State {
	case osb.StateSucceeded:
		newProvisionerLastOperation.State = provisionertypes.ProvisionerOperationStateSucceeded
	case osb.StateInProgress:
		newProvisionerLastOperation.State = provisionertypes.ProvisionerOperationStateInProgress
	default:
		newProvisionerLastOperation.State = provisionertypes.ProvisionerOperationStateFailed
	}

	return newProvisionerLastOperation, nil
}
