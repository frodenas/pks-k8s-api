/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package admissionserver

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/builder"
)

var (
	log        = logf.Log.WithName("admission-server")
	builderMap = map[string]*builder.WebhookBuilder{}
	// HandlerMap contains all admission webhook handlers.
	HandlerMap = map[string][]admission.Handler{}
)

// Add adds itself to the manager
func Add(mgr manager.Manager) error {
	ns := os.Getenv("POD_NAMESPACE")
	if len(ns) == 0 {
		ns = "default"
	}
	secretName := os.Getenv("SECRET_NAME")
	if len(secretName) == 0 {
		/* #nosec */
		secretName = "webhook-server-secret"
	}

	svr, err := webhook.NewServer("pks-admission-server", mgr, webhook.ServerOptions{
		Port:    9876,
		CertDir: "/tmp/cert",
		BootstrapOptions: &webhook.BootstrapOptions{
			Secret: &types.NamespacedName{
				Namespace: ns,
				Name:      secretName,
			},

			Service: &webhook.Service{
				Namespace: ns,
				Name:      "webhook-server-service",
				// Selectors should select the pods that runs this webhook server.
				Selectors: map[string]string{
					"control-plane": "controller-manager",
				},
			},
		},
	})
	if err != nil {
		return err
	}

	var webhooks []webhook.Webhook
	for k, builder := range builderMap {
		handlers, ok := HandlerMap[k]
		if !ok {
			log.V(1).Info(fmt.Sprintf("can't find handlers for builder: %v", k))
			handlers = []admission.Handler{}
		}
		wh, err := builder.
			Handlers(handlers...).
			WithManager(mgr).
			Build()
		if err != nil {
			return err
		}
		webhooks = append(webhooks, wh)
	}

	return svr.Register(webhooks...)
}
