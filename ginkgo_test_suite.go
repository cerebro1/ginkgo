package example

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgoExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoExample Suite")
}
