package hex_test

import (
	. "github.com/piderman314/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hex Tests", func() {
	var (
		grid Grid
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

			grid.Put(&hex1, 1, 1, -2)
			grid.Put(&hex2, 0, 2, -2)
			grid.Put(&hex3, 2, 2, -4)
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

	Describe("DiagonalNeighbours", func() {
		BeforeEach(func() {
			grid = NewGrid()
			hex1 = hex{}
			hex2 = hex{}
			hex3 = hex{}

			grid.Put(&hex1, 1, 1, -2)
			grid.Put(&hex2, 0, 3, -3)
			grid.Put(&hex3, 2, 3, -5)
		})

		Context("Checking actual diagonal neighbours", func() {
			It("should be true", func() {
				Expect(hex1.DiagonalNeighbours(&hex2)).To(BeTrue())
			})
		})

		Context("Checking not diagonal neighbours", func() {
			It("should be false", func() {
				Expect(hex1.DiagonalNeighbours(&hex3)).To(BeFalse())
				Expect(hex2.DiagonalNeighbours(&hex3)).To(BeFalse())
			})
		})
	})
})
