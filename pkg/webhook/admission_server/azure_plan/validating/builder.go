/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package validating

import (
	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/builder"
)

var (
	log = logf.Log.WithName("webhook.azureplan.validating")

	// Builders contain admission webhook builders
	Builders = map[string]*builder.WebhookBuilder{}
	// HandlerMap contains admission webhook handlers
	HandlerMap = map[string][]admission.Handler{}
)

func init() {
	builderName := "azureplan-validating-webhook"
	Builders[builderName] = builder.
		NewWebhookBuilder().
		Name(builderName+".pks.vcna.io").
		Path("/"+builderName).
		Validating().
		Operations(admissionregistrationv1beta1.Create, admissionregistrationv1beta1.Update).
		FailurePolicy(admissionregistrationv1beta1.Fail).
		ForType(&pksv1alpha1.AzurePlan{})
}
