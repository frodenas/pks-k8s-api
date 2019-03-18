/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package boshprovisioner

import (
	"encoding/json"
	"fmt"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	boshmanifest "github.com/frodenas/pks-k8s-api/pkg/provisioner/bosh/manifest"
	"github.com/frodenas/pks-k8s-api/pkg/provisioner/bosh/utils"
	provisionertypes "github.com/frodenas/pks-k8s-api/pkg/provisioner/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateCluster creates a BOSH Cluster.
func (p *Provisioner) CreateCluster(instance *pksv1alpha1.Cluster) (*provisionertypes.ProvisionerLastOperation, error) {
	deploymentName := utils.DeploymentName(instance.Namespace, instance.Name)
	log.Info(fmt.Sprintf("Creating BOSH Cluster `%s`", deploymentName))

	manifestGenerator := boshmanifest.NewManifestGenerator(deploymentName, instance)
	manifest, err := manifestGenerator.Generate()
	if err != nil {
		return nil, fmt.Errorf("error generating manifest for bosh cluster `%s`: %v", deploymentName, err)
	}

	taskID, err := p.boshClient.Deploy(deploymentName, manifest)
	if err != nil {
		return nil, fmt.Errorf("error creating On-Demand-Broker Cluster `%s`: %v", deploymentName, err)
	}

	boshProvisionerData := BOSHProvisionerData{
		TaskID: taskID,
	}
	provisionerData, err := json.Marshal(boshProvisionerData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling provisioner last operation data for BOSH Cluster `%s`: %v", deploymentName, err)
	}

	provisionerLastOperation := &provisionertypes.ProvisionerLastOperation{
		Description:     "operation in progress",
		LastUpdated:     metav1.NewTime(time.Now()),
		State:           provisionertypes.ProvisionerOperationStateInProgress,
		Type:            provisionertypes.ProvisionerOperationTypeCreate,
		ProvisionerData: string(provisionerData),
	}

	return provisionerLastOperation, nil
}
