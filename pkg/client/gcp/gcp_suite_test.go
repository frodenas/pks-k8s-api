package gcp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGCPClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GCP Client Suite")
}
