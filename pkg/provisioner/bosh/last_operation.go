/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package boshprovisioner

import (
	"encoding/json"
	"fmt"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/frodenas/pks-k8s-api/pkg/client/bosh"
	"github.com/frodenas/pks-k8s-api/pkg/provisioner/bosh/utils"
	provisionertypes "github.com/frodenas/pks-k8s-api/pkg/provisioner/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LastOperation polls the last operation for a BOSH Cluster.
func (p *Provisioner) LastOperation(instance *pksv1alpha1.Cluster, provisionerLastOperation provisionertypes.ProvisionerLastOperation) (*provisionertypes.ProvisionerLastOperation, error) {
	deploymentName := utils.DeploymentName(instance.Namespace, instance.Name)
	log.Info(fmt.Sprintf("Polling Last Operation for BOSH Cluster `%s`", deploymentName))

	boshProvisionerData := &BOSHProvisionerData{}
	err := json.Unmarshal([]byte(provisionerLastOperation.ProvisionerData), boshProvisionerData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling provisioner last operation data for BOSH Cluster `%s`: %v", deploymentName, err)
	}

	boshTask, err := p.boshClient.GetTask(boshProvisionerData.TaskID)
	if err != nil {
		return nil, fmt.Errorf("error getting task `%d` for BOSH Cluster `%s`: %v", boshProvisionerData.TaskID, deploymentName, err)
	}

	newBoshProvisionerData := BOSHProvisionerData{
		TaskID: boshTask.ID,
	}
	newProvisionerData, err := json.Marshal(newBoshProvisionerData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling provisioner last operation data for BOSH Cluster `%s`: %v", deploymentName, err)
	}

	newProvisionerLastOperation := &provisionertypes.ProvisionerLastOperation{
		Description:     fmt.Sprintf("%s: %s", boshTask.Description, boshTask.Result),
		LastUpdated:     metav1.NewTime(time.Now()),
		Type:            provisionerLastOperation.Type,
		ProvisionerData: string(newProvisionerData),
	}

	switch boshTask.StateType() {
	case bosh.TaskComplete:
		newProvisionerLastOperation.State = provisionertypes.ProvisionerOperationStateSucceeded
	case bosh.TaskIncomplete:
		newProvisionerLastOperation.State = provisionertypes.ProvisionerOperationStateInProgress
	default:
		newProvisionerLastOperation.State = provisionertypes.ProvisionerOperationStateFailed
	}

	return newProvisionerLastOperation, nil
}
