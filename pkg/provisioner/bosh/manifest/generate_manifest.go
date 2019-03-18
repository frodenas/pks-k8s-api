/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package boshmanifest

import (
	"fmt"

	"github.com/cppforlife/go-patch/patch"
	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

//go:generate go-bindata -pkg boshmanifest -o kubo.go -prefix ../../../../vendor/github.com/pivotal-cf/pks-kubo-deployment/manifests/ ../../../../vendor/github.com/pivotal-cf/pks-kubo-deployment/manifests/... ops-files/...

var log = logf.Log.WithName("bosh.manifest")

// ManifestGenerator is
type ManifestGenerator struct {
	deploymentName string
	instance       *pksv1alpha1.Cluster
}

// NewManifestGenerator returns a new manifest generator instance.
func NewManifestGenerator(deploymentName string, instance *pksv1alpha1.Cluster) *ManifestGenerator {
	return &ManifestGenerator{
		deploymentName: deploymentName,
		instance:       instance,
	}
}

// Generate generates a BOSH manifest given a deployment name.
func (mg *ManifestGenerator) Generate() ([]byte, error) {
	log.Info(fmt.Sprintf("Generating manifest for BOSH Cluster `%s`", mg.deploymentName))

	kuboBytes, err := Asset("cfcr.yml")
	if err != nil {
		return nil, fmt.Errorf("failed to load `cfcr.yml`: %v", err)
	}

	_, err = mg.getOps(mg.instance)
	if err != nil {
		return nil, fmt.Errorf("Could not get ops %v", err)
	}

	var boshManifest BoshManifest
	err = yaml.Unmarshal(kuboBytes, &boshManifest)
	if err != nil {
		return nil, errors.Errorf("Failed to unmarshal `cfcr.yml`: %v", err)
	}

	boshManifest.Name = mg.deploymentName
	boshManifest.Properties["cluster-name"] = fmt.Sprintf("%s-%s", mg.instance.Namespace, mg.instance.Name)

	manifest, err := yaml.Marshal(boshManifest)
	if err != nil {
		return nil, errors.Errorf("Failed to marshal manifest: %v", err)
	}

	return manifest, nil
}

func (p *ManifestGenerator) getOps(instance *pksv1alpha1.Cluster) (patch.Ops, error) {
	opsFiles, err := p.getOpsFiles(instance)
	if err != nil {
		return patch.Ops{}, err
	}
	var op []patch.Op
	for _, opFile := range opsFiles {
		opsBytes, err := Asset(opFile)
		if err != nil {
			return patch.Ops{}, errors.Errorf("Failed to load ops file %s due to %s", opFile, err.Error())
		}
		op, err = createOps(op, opsBytes, opFile)
		if err != nil {
			return patch.Ops{}, err
		}
	}
	ops := patch.Ops(op)

	return patch.Ops(ops), nil
}

func (p *ManifestGenerator) getOpsFiles(instance *pksv1alpha1.Cluster) ([]string, error) {
	opsFiles := []string{
		"ops-files/misc/bootstrap.yml",
		"ops-files/misc/first-time-deploy.yml",
		"ops-files/update-overrides.yml",
		"ops-files/use-runtime-config-bosh-dns.yml",
		"ops-files/drain-cluster-errand.yml",
		"ops-files/update-worker-overrides.yml",
		"ops-files/pks-master-aliases.yml",
		"ops-files/enable-nfs.yml",
		"ops-files/modify-audit-log.yml",
	}

	return opsFiles, nil
}

func createOps(opList []patch.Op, opsBytes []byte, opFile string) (patch.Ops, error) {
	var opDefinitions []patch.OpDefinition
	err := yaml.Unmarshal(opsBytes, &opDefinitions)
	if err != nil {
		return patch.Ops{}, errors.Errorf("Failed to unmarshal ops file %s due to %s", opFile, err.Error())
	}
	tmpOps, err := patch.NewOpsFromDefinitions(opDefinitions)
	if err != nil {
		return patch.Ops{}, errors.Errorf("Failed to create ops from definitions in %s due to %s", opFile, err.Error())
	}
	return append(opList, tmpOps), nil
}
