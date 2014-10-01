package hex

type Grid struct {
	grid map[hexIndex]IndexableHex
}

func NewGrid() Grid {
	return Grid{grid: make(map[hexIndex]IndexableHex)}
}

func (grid *Grid) Put(hex IndexableHex, x, y, z int) {
	hex.SetIndex(hexIndex{x, y, z})
	(*grid).grid[hex.GetIndex()] = hex
}

func (grid *Grid) Get(x, y, z int) IndexableHex {
	return (*grid).grid[hexIndex{x, y, z}]
}

func (grid *Grid) GetAllNeighbours(hex IndexableHex) (out []IndexableHex) {
	index := hex.GetIndex()

	for _, i := range neighbours {
		hex := grid.Get(index.x+i[0], index.y+i[1], index.z+i[2])
		if hex != nil {
			out = append(out, hex)
		}
	}

	return
}
