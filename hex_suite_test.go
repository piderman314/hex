package hex_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hex Suite")
}
