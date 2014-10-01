package hex_test

import (
	. "github.com/piderman314/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HexGrid Tests", func() {
	var (
		grid Grid
	)

	Describe("HexGrid Get and Put", func() {
		BeforeEach(func() {
			grid = NewGrid()
			grid.Put(&hex{}, 1, 1, 1)
		})

		Context("Getting from a location with a hex", func() {
			It("should not be nil", func() {
				Expect(grid.Get(1, 1, 1)).NotTo(BeNil())
			})
		})

		Context("Getting from a location without a hex", func() {
			It("should be nil", func() {
				Expect(grid.Get(2, 2, 2)).To(BeNil())
			})
		})
	})

	Describe("HexGrid GetAllNeighbours", func() {
		Context("Getting from a location without neighbours", func() {
			c := hex{}
			grid = NewGrid()
			grid.Put(&c, 1, 1, 1)
			grid.Put(&hex{}, 3, 3, 3)

			It("should be an empty slice", func() {
				Expect(len(grid.GetAllNeighbours(&c))).To(Equal(0))
			})
		})

		Context("Getting from a location with 6 neighbours", func() {
			c := hex{}
			grid := NewGrid()
			grid.Put(&c, 1, 1, 1)
			grid.Put(&hex{}, 2, 0, 1)
			grid.Put(&hex{}, 0, 2, 1)
			grid.Put(&hex{}, 1, 2, 0)
			grid.Put(&hex{}, 1, 0, 2)
			grid.Put(&hex{}, 2, 1, 0)
			grid.Put(&hex{}, 0, 1, 2)

			It("should be a slice if size 6", func() {
				Expect(len(grid.GetAllNeighbours(&c))).To(Equal(6))
			})
		})
	})
})
