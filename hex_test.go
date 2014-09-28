package hex

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHexGridPut(t *testing.T) {
	Convey("Given a HexGrid", t, func() {
		grid := NewGrid()

		Convey("When a cell is inserted", func() {
			grid.Put(&testCell{}, 1, 1, 1)

			Convey("And retrieved", func() {
				cell := grid.Get(1, 1, 1)

				Convey("It should not be nil", func() {
					So(cell, ShouldNotBeNil)
				})
			})
		})
	})
}

type testCell struct {
	HexCell
}
