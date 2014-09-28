package hex_test

import (
	. "github.com/piderman314/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HexGrid with cell in 1, 1, 1", func() {
	var (
		grid HexGrid
	)

	BeforeEach(func() {
		grid = NewGrid()
		grid.Put(&cell{}, 1, 1, 1)
	})

	Describe("when retrieving from the HexGrid", func() {
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

type cell struct {
	HexCell
}
