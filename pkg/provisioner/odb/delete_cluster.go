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

// DeleteCluster deletes an On-Demand-Broker Cluster.
func (p *Provisioner) DeleteCluster(instance *pksv1alpha1.Cluster) (*provisionertypes.ProvisionerLastOperation, error) {
	serviceInstanceName := utils.ServiceInstanceName(instance.Namespace, instance.Name)
	log.Info(fmt.Sprintf("Deleting On-Demand-Broker Cluster `%s`", serviceInstanceName))

	// Read the ODB plan.
	if instance.Spec.PlanRef.Kind != pksv1alpha1.ODBPlanKind {
		return nil, fmt.Errorf("Clusters using `ODB` provisioner can only reference an `ODBPlan`")
	}
	odbPlan := &pksv1alpha1.ODBPlan{}
	if err := p.k8sClient.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, odbPlan); err != nil {
		return nil, err
	}

	// Deprovision instance.
	deprovisionRequest := &osb.DeprovisionRequest{
		InstanceID:        serviceInstanceName,
		AcceptsIncomplete: true,
		ServiceID:         odbPlan.Spec.ServiceID,
		PlanID:            odbPlan.Spec.PlanID,
	}

	deprovisionResponse, err := p.osbClient.DeprovisionInstance(deprovisionRequest)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("error deleting On-Demand-Broker Cluster `%s`: %v", serviceInstanceName, err)
		}
	}

	// Return last operation.
	provisionerData, err := json.Marshal(deprovisionResponse.OperationKey)
	if err != nil {
		return nil, fmt.Errorf("error marshalling provisioner last operation data for On-Demand-Broker Cluster `%s`: %v", serviceInstanceName, err)
	}

	provisionerLastOperation := &provisionertypes.ProvisionerLastOperation{
		Description:     "operation in progress",
		LastUpdated:     metav1.NewTime(time.Now()),
		State:           provisionertypes.ProvisionerOperationStateInProgress,
		Type:            provisionertypes.ProvisionerOperationTypeDelete,
		ProvisionerData: string(provisionerData),
	}

	return provisionerLastOperation, nil
}
