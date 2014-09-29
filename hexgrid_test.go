package hex_test

import (
	. "github.com/piderman314/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HexGrid", func() {
	var (
		grid HexGrid
	)

	Describe("HexGrid Get and Put", func() {
		Describe("When retrieving from the HexGrid", func() {
			BeforeEach(func() {
				grid = NewGrid()
				grid.Put(&cell{}, 1, 1, 1)
			})

			Context("from location 1, 1, 1", func() {
				It("should not be nil", func() {
					Expect(grid.Get(1, 1, 1)).NotTo(BeNil())
				})
			})

			Context("from location 2, 2, 2", func() {
				It("should be nil", func() {
					Expect(grid.Get(2, 2, 2)).To(BeNil())
				})
			})
		})
	})

	Describe("HexGrid GetAllNeighbours", func() {
		Context("no neighbours", func() {
			c := cell{}
			grid = NewGrid()
			grid.Put(&c, 1, 1, 1)
			grid.Put(&cell{}, 3, 3, 3)

			It("should be an empty slice", func() {
				Expect(len(grid.GetAllNeighbours(&c))).To(Equal(0))
			})
		})

		Context("all neighbours", func() {
			c := cell{}
			grid := NewGrid()
			grid.Put(&c, 1, 1, 1)
			grid.Put(&cell{}, 2, 0, 1)
			grid.Put(&cell{}, 0, 2, 1)
			grid.Put(&cell{}, 1, 2, 0)
			grid.Put(&cell{}, 1, 0, 2)
			grid.Put(&cell{}, 2, 1, 0)
			grid.Put(&cell{}, 0, 1, 2)

			It("should be a slice if size 6", func() {
				Expect(len(grid.GetAllNeighbours(&c))).To(Equal(6))
			})
		})
	})
})

type cell struct {
	HexCell
}
