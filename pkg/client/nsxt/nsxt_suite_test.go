package nsxt_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNSXTClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NSX-T Client Suite")
}
