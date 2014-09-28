package hex

type HexGrid struct {
	grid map[hexIndex]IndexableHex
}

func NewGrid() HexGrid {
	return HexGrid{grid: make(map[hexIndex]IndexableHex)}
}

func (grid *HexGrid) Put(hex IndexableHex, x, y, z int) {
	hex.SetIndex(hexIndex{x, y, z})
	(*grid).grid[hex.GetIndex()] = hex
}

func (grid *HexGrid) Get(x, y, z int) IndexableHex {
	return (*grid).grid[hexIndex{x, y, z}]
}

func (grid *HexGrid) GetAllNeighbours(hex IndexableHex) (out []IndexableHex) {
	index := hex.GetIndex()

	for _, i := range neighbours {
		hex := grid.Get(index.x+i[0], index.y+i[1], index.z+i[2])
		if hex != nil {
			out = append(out, hex)
		}
	}

	return
}
