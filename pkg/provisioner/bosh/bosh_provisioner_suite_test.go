package boshprovisioner_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBOSHProvisioner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BOSH Provisioner Suite")
}
