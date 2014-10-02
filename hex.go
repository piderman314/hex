package hex

import (
	"math"
)

var neighbours = [][]int{
	{+1, -1, 0}, {+1, 0, -1}, {0, +1, -1},
	{-1, +1, 0}, {-1, 0, +1}, {0, -1, +1},
}

type Hex struct {
	index Index
}

func (hex *Hex) GetIndex() Index {
	return hex.index
}

func (hex *Hex) SetIndex(index Index) {
	hex.index = index
}

func (hex *Hex) Neighbours(other IndexableHex) bool {
	diffX := float64(hex.index.x - other.GetIndex().x)
	diffY := float64(hex.index.y - other.GetIndex().y)
	diffZ := float64(hex.index.z - other.GetIndex().z)

	return math.Abs(diffX) <= 1 &&
		math.Abs(diffY) <= 1 &&
		math.Abs(diffZ) <= 1
}

type IndexableHex interface {
	GetIndex() Index
	SetIndex(Index)
}

type Index struct {
	x, y, z int
}
