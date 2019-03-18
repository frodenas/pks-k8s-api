package vcenter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVCenterClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "vCenter Client Suite")
}
