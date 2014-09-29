package hex_test

import (
	. "github.com/piderman314/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hex Tests", func() {
	var (
		grid HexGrid
		hex1 hex
		hex2 hex
		hex3 hex
	)

	Describe("Neighbours", func() {
		BeforeEach(func() {
			grid = NewGrid()
			hex1 = hex{}
			hex2 = hex{}
			hex3 = hex{}

			grid.Put(&hex1, 1, 1, 1)
			grid.Put(&hex2, 0, 2, 1)
			grid.Put(&hex3, 2, 2, 2)
		})

		Context("Checking actual neighbours", func() {
			It("should be true", func() {
				Expect(hex1.Neighbours(&hex2)).To(BeTrue())
			})
		})

		Context("Checking not neighbours", func() {
			It("should be false", func() {
				Expect(hex1.Neighbours(&hex3)).To(BeFalse())
				Expect(hex2.Neighbours(&hex3)).To(BeFalse())
			})
		})
	})

})
