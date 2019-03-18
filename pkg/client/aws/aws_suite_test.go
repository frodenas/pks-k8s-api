package aws_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAWSClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AWS Client Suite")
}
