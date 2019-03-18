package bosh_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBOSHClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BOSH Client Suite")
}
