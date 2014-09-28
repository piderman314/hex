package hex

type HexCell struct {
	index hexIndex
}

type hexCellable interface {
	GetIndex() hexIndex
	SetIndex(hexIndex)
}

func (cell *HexCell) GetIndex() hexIndex {
	return cell.index
}

func (cell *HexCell) SetIndex(index hexIndex) {
	cell.index = index
}

type HexGrid struct {
	grid map[hexIndex]hexCellable
}

type hexIndex struct {
	x, y, z int
}

func NewGrid() HexGrid {
	return HexGrid{grid: make(map[hexIndex]hexCellable)}
}

func (grid *HexGrid) Put(cell hexCellable, x, y, z int) {
	cell.SetIndex(hexIndex{x, y, z})
	(*grid).grid[cell.GetIndex()] = cell
}

func (grid *HexGrid) Get(x, y, z int) hexCellable {
	return (*grid).grid[hexIndex{x, y, z}]
}
