/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package provisioner

import (
	"fmt"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	boshprovisioner "github.com/frodenas/pks-k8s-api/pkg/provisioner/bosh"
	dummyprovisioner "github.com/frodenas/pks-k8s-api/pkg/provisioner/dummy"
	odbprovisioner "github.com/frodenas/pks-k8s-api/pkg/provisioner/odb"
	provisionertypes "github.com/frodenas/pks-k8s-api/pkg/provisioner/types"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Provisioner represents a Cluster provisioner.
//go:generate moq -out fakes/provisioner.go -pkg fakes . Provisioner
type Provisioner interface {
	CreateCluster(instance *pksv1alpha1.Cluster) (*provisionertypes.ProvisionerLastOperation, error)
	DeleteCluster(instance *pksv1alpha1.Cluster) (*provisionertypes.ProvisionerLastOperation, error)
	LastOperation(instance *pksv1alpha1.Cluster, provisionerLastOperation provisionertypes.ProvisionerLastOperation) (*provisionertypes.ProvisionerLastOperation, error)
}

// New returns a new Provisioner.
func New(provisionerType string, k8sClient client.Client, credentialsSecret *corev1.Secret) (Provisioner, error) {
	switch provisionerType {
	case pksv1alpha1.DUMMYProvisioner:
		return dummyprovisioner.NewProvisioner(k8sClient, credentialsSecret)
	case pksv1alpha1.BOSHProvisioner:
		return boshprovisioner.NewProvisioner(k8sClient, credentialsSecret)
	case pksv1alpha1.ODBProvisioner:
		return odbprovisioner.NewProvisioner(k8sClient, credentialsSecret)
	}

	return nil, fmt.Errorf("Provisioner `%s` not supported", provisionerType)
}
