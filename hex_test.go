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

	Describe("Retrieving from the HexGrid", func() {
		Context("Retrieving from 1, 1, 1", func() {
			It("should not be nil", func() {
				Expect(grid.Get(1, 1, 1)).NotTo(BeNil())
			})
		})
	})

})

type cell struct {
	HexCell
}
