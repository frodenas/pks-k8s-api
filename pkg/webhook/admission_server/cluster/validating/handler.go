/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package validating

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apitypes "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/types"
)

func init() {
	webhookName := "cluster-validating-webhook"
	if HandlerMap[webhookName] == nil {
		HandlerMap[webhookName] = []admission.Handler{}
	}
	HandlerMap[webhookName] = append(HandlerMap[webhookName], &Handler{})
}

// Handler handles Cluster
type Handler struct {
	Client client.Client

	// Decoder decodes objects
	Decoder types.Decoder
}

func (h *Handler) validatingClusterFn(ctx context.Context, cluster *pksv1alpha1.Cluster) (bool, string, error) {
	log.Info(fmt.Sprintf("Validating Cluster `%s/%s`", cluster.Namespace, cluster.Name))

	// Validate the Provisioner Credentials secret.
	if err := h.validateProvisionerCredentialsSecret(ctx, cluster); err != nil {
		log.Error(err, fmt.Sprintf("Error validating Provisioner Credentials Secret `%s/%s` for Cluster `%s/%s`", cluster.Spec.ProvisionerSpec.CredentialsSecretRef.Namespace, cluster.Spec.ProvisionerSpec.CredentialsSecretRef.Name, cluster.Namespace, cluster.Name))
		return false, "error validating cluster", fmt.Errorf("Error validating Provisioner Credentials Secret `%s/%s` for Cluster `%s/%s`: %v", cluster.Spec.ProvisionerSpec.CredentialsSecretRef.Namespace, cluster.Spec.ProvisionerSpec.CredentialsSecretRef.Name, cluster.Namespace, cluster.Name, err)
	}

	// Validate the Plan reference.
	if err := h.validatePlanRef(ctx, cluster); err != nil {
		log.Error(err, fmt.Sprintf("Error validating %s `%s/%s` for Cluster `%s/%s`", cluster.Spec.PlanRef.Kind, cluster.Spec.PlanRef.Namespace, cluster.Spec.PlanRef.Name, cluster.Namespace, cluster.Name))
		return false, "error validating cluster", fmt.Errorf("Error validating %s `%s/%s` for Cluster `%s/%s`: %v", cluster.Spec.PlanRef.Kind, cluster.Spec.PlanRef.Namespace, cluster.Spec.PlanRef.Name, cluster.Namespace, cluster.Name, err)
	}

	// Validate the Kubernetes Profile reference.
	if err := h.validateKubernetesProfileRef(ctx, cluster); err != nil {
		log.Error(err, fmt.Sprintf("Error validating Kubernetes Profile `%s/%s` for Cluster `%s/%s`", cluster.Spec.KubernetesProfileRef.Namespace, cluster.Spec.KubernetesProfileRef.Name, cluster.Namespace, cluster.Name))
		return false, "error validating cluster", fmt.Errorf("Error validating Kubernetes Profile `%s/%s` for Cluster `%s/%s`: %v", cluster.Spec.KubernetesProfileRef.Namespace, cluster.Spec.KubernetesProfileRef.Name, cluster.Namespace, cluster.Name, err)
	}

	return true, "allowed to be admitted", nil
}

func (h *Handler) validateProvisionerCredentialsSecret(ctx context.Context, cluster *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating Provisioner Credentials Secret `%s/%s` for Cluster `%s/%s`", cluster.Spec.ProvisionerSpec.CredentialsSecretRef.Namespace, cluster.Spec.ProvisionerSpec.CredentialsSecretRef.Name, cluster.Namespace, cluster.Name))

	provisionerCredentialsSecret := &corev1.Secret{}
	if err := h.Client.Get(context.TODO(), apitypes.NamespacedName{Name: cluster.Spec.ProvisionerSpec.CredentialsSecretRef.Name, Namespace: cluster.Spec.ProvisionerSpec.CredentialsSecretRef.Namespace}, provisionerCredentialsSecret); err != nil {
		return err
	}

	return nil
}

func (h *Handler) validatePlanRef(ctx context.Context, cluster *pksv1alpha1.Cluster) error {
	switch cluster.Spec.PlanRef.Kind {
	case pksv1alpha1.AWSPlanKind:
		if err := h.validateAWSPlanRef(ctx, cluster); err != nil {
			return err
		}
	case pksv1alpha1.AzurePlanKind:
		if err := h.validateAzurePlanRef(ctx, cluster); err != nil {
			return err
		}
	case pksv1alpha1.GCPPlanKind:
		if err := h.validateGCPPlanRef(ctx, cluster); err != nil {
			return err
		}
	case pksv1alpha1.VSpherePlanKind:
		if err := h.validateVSpherePlanRef(ctx, cluster); err != nil {
			return err
		}
	case pksv1alpha1.ODBPlanKind:
		if err := h.validateODBPlanRef(ctx, cluster); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Plan Kind `%s` for Cluster `%s/%s` not supported", cluster.Spec.PlanRef.Kind, cluster.Namespace, cluster.Name)
	}

	switch cluster.Spec.ProvisionerSpec.Type {
	case pksv1alpha1.ODBProvisioner:
		if cluster.Spec.PlanRef.Kind != pksv1alpha1.ODBPlanKind {
			return errors.New("Clusters using `ODB` provisioner can only reference an `ODBPlan`")
		}
	default:
		if cluster.Spec.PlanRef.Kind == pksv1alpha1.ODBPlanKind {
			return errors.New("`ODBPlan` can only be referenced by clusters using `ODB` provisioner")
		}
	}

	return nil
}

func (h *Handler) validateAWSPlanRef(ctx context.Context, cluster *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating AWS Plan `%s/%s` for Cluster `%s/%s`", cluster.Spec.PlanRef.Namespace, cluster.Spec.PlanRef.Name, cluster.Namespace, cluster.Name))

	awsPlan := &pksv1alpha1.AWSPlan{}
	if err := h.Client.Get(context.TODO(), apitypes.NamespacedName{Name: cluster.Spec.PlanRef.Name, Namespace: cluster.Spec.PlanRef.Namespace}, awsPlan); err != nil {
		return err
	}

	return nil
}

func (h *Handler) validateAzurePlanRef(ctx context.Context, cluster *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating Azure Plan `%s/%s` for Cluster `%s/%s`", cluster.Spec.PlanRef.Namespace, cluster.Spec.PlanRef.Name, cluster.Namespace, cluster.Name))

	azurePlan := &pksv1alpha1.AzurePlan{}
	if err := h.Client.Get(context.TODO(), apitypes.NamespacedName{Name: cluster.Spec.PlanRef.Name, Namespace: cluster.Spec.PlanRef.Namespace}, azurePlan); err != nil {
		return err
	}

	return nil
}

func (h *Handler) validateGCPPlanRef(ctx context.Context, cluster *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating GCP Plan `%s/%s` for Cluster `%s/%s`", cluster.Spec.PlanRef.Namespace, cluster.Spec.PlanRef.Name, cluster.Namespace, cluster.Name))

	gcpPlan := &pksv1alpha1.GCPPlan{}
	if err := h.Client.Get(context.TODO(), apitypes.NamespacedName{Name: cluster.Spec.PlanRef.Name, Namespace: cluster.Spec.PlanRef.Namespace}, gcpPlan); err != nil {
		return err
	}

	return nil
}

func (h *Handler) validateVSpherePlanRef(ctx context.Context, cluster *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating VSphere Plan `%s/%s` for Cluster `%s/%s`", cluster.Spec.PlanRef.Namespace, cluster.Spec.PlanRef.Name, cluster.Namespace, cluster.Name))

	vSpherePlan := &pksv1alpha1.VSpherePlan{}
	if err := h.Client.Get(context.TODO(), apitypes.NamespacedName{Name: cluster.Spec.PlanRef.Name, Namespace: cluster.Spec.PlanRef.Namespace}, vSpherePlan); err != nil {
		return err
	}

	return nil
}

func (h *Handler) validateODBPlanRef(ctx context.Context, cluster *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating ODB Plan `%s/%s` for Cluster `%s/%s`", cluster.Spec.PlanRef.Namespace, cluster.Spec.PlanRef.Name, cluster.Namespace, cluster.Name))

	odbPlan := &pksv1alpha1.ODBPlan{}
	if err := h.Client.Get(context.TODO(), apitypes.NamespacedName{Name: cluster.Spec.PlanRef.Name, Namespace: cluster.Spec.PlanRef.Namespace}, odbPlan); err != nil {
		return err
	}

	return nil
}

func (h *Handler) validateKubernetesProfileRef(ctx context.Context, cluster *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating Kubernetes Profile `%s/%s` for Cluster `%s/%s`", cluster.Spec.KubernetesProfileRef.Namespace, cluster.Spec.KubernetesProfileRef.Name, cluster.Namespace, cluster.Name))

	kubernetesProfile := &pksv1alpha1.KubernetesProfile{}
	if err := h.Client.Get(context.TODO(), apitypes.NamespacedName{Name: cluster.Spec.KubernetesProfileRef.Name, Namespace: cluster.Spec.KubernetesProfileRef.Namespace}, kubernetesProfile); err != nil {
		return err
	}

	return nil
}

var _ admission.Handler = &Handler{}

// Handle handles admission requests.
func (h *Handler) Handle(ctx context.Context, req types.Request) types.Response {
	obj := &pksv1alpha1.Cluster{}

	err := h.Decoder.Decode(req, obj)
	if err != nil {
		return admission.ErrorResponse(http.StatusBadRequest, err)
	}

	allowed, reason, err := h.validatingClusterFn(ctx, obj)
	if err != nil {
		return admission.ErrorResponse(http.StatusInternalServerError, err)
	}

	return admission.ValidationResponse(allowed, reason)
}

var _ inject.Client = &Handler{}

// InjectClient injects the client into the Handler
func (h *Handler) InjectClient(c client.Client) error {
	h.Client = c
	return nil
}

var _ inject.Decoder = &Handler{}

// InjectDecoder injects the decoder into the Handler
func (h *Handler) InjectDecoder(d types.Decoder) error {
	h.Decoder = d
	return nil
}
