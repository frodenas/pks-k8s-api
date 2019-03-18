package dummyprovisioner_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDummyProvisioner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dummy Provisioner Suite")
}
