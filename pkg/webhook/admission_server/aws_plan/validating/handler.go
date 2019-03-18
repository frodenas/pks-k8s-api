/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package validating

import (
	"context"
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
	webhookName := "awsplan-validating-webhook"
	if HandlerMap[webhookName] == nil {
		HandlerMap[webhookName] = []admission.Handler{}
	}
	HandlerMap[webhookName] = append(HandlerMap[webhookName], &Handler{})
}

// Handler handles AWS Plan
type Handler struct {
	Client client.Client

	// Decoder decodes objects
	Decoder types.Decoder
}

func (h *Handler) validatingAWSPlanFn(ctx context.Context, awsPlan *pksv1alpha1.AWSPlan) (bool, string, error) {
	log.Info(fmt.Sprintf("Validating AWS Plan `%s/%s`", awsPlan.Namespace, awsPlan.Name))

	// Validate the Provider Credentials Secret.
	providerCredentialsSecret := &corev1.Secret{}
	if err := h.Client.Get(context.TODO(), apitypes.NamespacedName{Name: awsPlan.Spec.ProviderSpec.CredentialsSecretRef.Name, Namespace: awsPlan.Spec.ProviderSpec.CredentialsSecretRef.Namespace}, providerCredentialsSecret); err != nil {
		log.Error(err, fmt.Sprintf("Error validating Provider Credentials Secret `%s/%s` for AWS Plan `%s/%s`", awsPlan.Spec.ProviderSpec.CredentialsSecretRef.Namespace, awsPlan.Spec.ProviderSpec.CredentialsSecretRef.Name, awsPlan.Namespace, awsPlan.Name))
		return false, "error validating awsplan", fmt.Errorf("Error validating Provider Credentials Secret `%s/%s` for AWS Plan `%s/%s`: %v", awsPlan.Spec.ProviderSpec.CredentialsSecretRef.Namespace, awsPlan.Spec.ProviderSpec.CredentialsSecretRef.Name, awsPlan.Namespace, awsPlan.Name, err)

	}

	return true, "allowed to be admitted", nil
}

var _ admission.Handler = &Handler{}

// Handle handles admission requests.
func (h *Handler) Handle(ctx context.Context, req types.Request) types.Response {
	obj := &pksv1alpha1.AWSPlan{}

	err := h.Decoder.Decode(req, obj)
	if err != nil {
		return admission.ErrorResponse(http.StatusBadRequest, err)
	}

	allowed, reason, err := h.validatingAWSPlanFn(ctx, obj)
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
