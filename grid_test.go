package hex_test

import (
	. "github.com/piderman314/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid Tests", func() {
	var (
		grid Grid
	)

	Describe("Grid Get and Put", func() {
		BeforeEach(func() {
			grid = NewGrid()
			grid.Put(&hex{}, 1, 1, -2)
		})

		Context("Getting from a location with a hex", func() {
			It("should not be nil", func() {
				Expect(grid.GetByCoords(1, 1, -2)).NotTo(BeNil())
			})
		})

		Context("Getting from a location without a hex", func() {
			It("should be nil", func() {
				Expect(grid.GetByCoords(-1, -1, 2)).To(BeNil())
			})
		})

		Context("When coordinates don't add up to 0", func() {
			It("should produce an error", func() {
				_, err := grid.Put(&hex{}, 1, 1, 1)
				Expect(err.Error()).To(Equal("x, y and z don't add up to 0."))
			})
		})
	})

	Describe("Grid GetAllNeighbours", func() {
		Context("Getting from a location without neighbours", func() {
			c := hex{}
			grid = NewGrid()
			grid.Put(&c, 1, 1, -2)
			grid.Put(&hex{}, 3, 3, -6)

			It("should be an empty slice", func() {
				Expect(len(grid.GetAllNeighbours(&c))).To(Equal(0))
			})
		})

		Context("Getting from a location with 6 neighbours", func() {
			c := hex{}
			grid := NewGrid()
			grid.Put(&c, 1, 1, -2)
			grid.Put(&hex{}, 2, 0, -2)
			grid.Put(&hex{}, 0, 2, -2)
			grid.Put(&hex{}, 1, 2, -3)
			grid.Put(&hex{}, 1, 0, -1)
			grid.Put(&hex{}, 2, 1, -3)
			grid.Put(&hex{}, 0, 1, -1)

			It("should be a slice if size 6", func() {
				Expect(len(grid.GetAllNeighbours(&c))).To(Equal(6))
			})
		})
	})
})
