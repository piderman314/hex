package hex_test

import (
	. "github.com/piderman314/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hex Suite")
}

type hex struct {
	Hex
}
