package hex

type HexCell struct {
	index hexIndex
}

func (cell *HexCell) GetIndex() hexIndex {
	return cell.index
}

func (cell *HexCell) SetIndex(index hexIndex) {
	cell.index = index
}
