package hex

type HexGrid struct {
	grid map[hexIndex]IndexableHex
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
