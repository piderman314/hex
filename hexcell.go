package hex

var neighbours = [][]int{
	{+1, -1, 0}, {+1, 0, -1}, {0, +1, -1},
	{-1, +1, 0}, {-1, 0, +1}, {0, -1, +1},
}

type HexCell struct {
	index hexIndex
}

func (cell *HexCell) GetIndex() hexIndex {
	return cell.index
}

func (cell *HexCell) SetIndex(index hexIndex) {
	cell.index = index
}
