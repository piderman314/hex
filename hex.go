package hex

var neighbours = [][]int{
	{+1, -1, 0}, {+1, 0, -1}, {0, +1, -1},
	{-1, +1, 0}, {-1, 0, +1}, {0, -1, +1},
}

type Hex struct {
	index hexIndex
}

func (hex *Hex) GetIndex() hexIndex {
	return hex.index
}

func (hex *Hex) SetIndex(index hexIndex) {
	hex.index = index
}

type IndexableHex interface {
	GetIndex() hexIndex
	SetIndex(hexIndex)
}

type hexIndex struct {
	x, y, z int
}
