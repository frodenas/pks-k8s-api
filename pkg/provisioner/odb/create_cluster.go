/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package odbprovisioner

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/frodenas/pks-k8s-api/pkg/provisioner/odb/utils"
	provisionertypes "github.com/frodenas/pks-k8s-api/pkg/provisioner/types"
	osb "github.com/maplain/go-open-service-broker-client/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apitypes "k8s.io/apimachinery/pkg/types"
)

// CreateCluster creates an On-Demand-Broker Cluster.
func (p *Provisioner) CreateCluster(instance *pksv1alpha1.Cluster) (*provisionertypes.ProvisionerLastOperation, error) {
	serviceInstanceName := utils.ServiceInstanceName(instance.Namespace, instance.Name)
	log.Info(fmt.Sprintf("Creating On-Demand-Broker Cluster `%s`", serviceInstanceName))

	// Read the ODB plan.
	if instance.Spec.PlanRef.Kind != pksv1alpha1.ODBPlanKind {
		return nil, fmt.Errorf("Clusters using `ODB` provisioner can only reference an `ODBPlan`")
	}
	odbPlan := &pksv1alpha1.ODBPlan{}
	if err := p.k8sClient.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, odbPlan); err != nil {
		return nil, err
	}

	// Provision instance.
	provisionParameters := map[string]interface{}{
		"clusterName":            fmt.Sprintf("%s-%s", instance.Namespace, instance.Name),
		"kubernetes_master_host": instance.Spec.ExternalHostname,
		"kubernetes_master_port": 8443,
	}

	if instance.Spec.NumWorkerReplicas > 0 {
		provisionParameters["kubernetesWorkerInstances"] = instance.Spec.NumWorkerReplicas
	}

	provisionRequest := &osb.ProvisionRequest{
		InstanceID:        serviceInstanceName,
		AcceptsIncomplete: true,
		ServiceID:         odbPlan.Spec.ServiceID,
		PlanID:            odbPlan.Spec.PlanID,
		OrganizationGUID:  "organization_guid",
		SpaceGUID:         "space_guid",
		Parameters:        provisionParameters,
		Context:           map[string]interface{}{},
	}
	provisionResponse, err := p.osbClient.ProvisionInstance(provisionRequest)
	if err != nil {
		return nil, fmt.Errorf("error creating On-Demand-Broker Cluster `%s`: %v", serviceInstanceName, err)
	}

	// Return last operation.
	provisionerData, err := json.Marshal(provisionResponse.OperationKey)
	if err != nil {
		return nil, fmt.Errorf("error marshalling provisioner last operation data for On-Demand-Broker Cluster `%s`: %v", serviceInstanceName, err)
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
