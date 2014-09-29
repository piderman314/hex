package hex

var neighbours = [][]int{
	{+1, -1, 0}, {+1, 0, -1}, {0, +1, -1},
	{-1, +1, 0}, {-1, 0, +1}, {0, -1, +1},
}

type Hex struct {
	index hexIndex
}

func (cell *Hex) GetIndex() hexIndex {
	return cell.index
}

func (cell *Hex) SetIndex(index hexIndex) {
	cell.index = index
}

type IndexableHex interface {
	GetIndex() hexIndex
	SetIndex(hexIndex)
}

type hexIndex struct {
	x, y, z int
}
