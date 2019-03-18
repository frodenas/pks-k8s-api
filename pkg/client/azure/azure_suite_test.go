package azure_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAzureClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Azure Client Suite")
}
