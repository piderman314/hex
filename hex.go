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

type IndexableHex interface {
	GetIndex() hexIndex
	SetIndex(hexIndex)
}

type HexGrid struct {
	grid map[hexIndex]IndexableHex
}

type hexIndex struct {
	x, y, z int
}

func NewGrid() HexGrid {
	return HexGrid{grid: make(map[hexIndex]IndexableHex)}
}

func (grid *HexGrid) Put(cell IndexableHex, x, y, z int) {
	cell.SetIndex(hexIndex{x, y, z})
	(*grid).grid[cell.GetIndex()] = cell
}

func (grid *HexGrid) Get(x, y, z int) IndexableHex {
	return (*grid).grid[hexIndex{x, y, z}]
}
