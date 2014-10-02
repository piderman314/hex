package hex

import (
	"errors"
)

type Grid struct {
	grid map[Index]IndexableHex
}

func NewGrid() Grid {
	return Grid{grid: make(map[Index]IndexableHex)}
}

func (grid *Grid) Put(hex IndexableHex, x, y, z int) (*Index, error) {
	if x+y+z != 0 {
		return nil, errors.New("x, y and z don't add up to 0.")
	}

	index := Index{x, y, z}
	hex.SetIndex(index)
	(*grid).grid[index] = hex
	return &index, nil
}

func (grid *Grid) GetByCoords(x, y, z int) (IndexableHex, error) {
	if x+y+z != 0 {
		return nil, errors.New("x, y and z don't add up to 0.")
	}

	return (*grid).GetByIndex(Index{x, y, z}), nil
}

func (grid *Grid) GetByIndex(index Index) IndexableHex {
	return (*grid).grid[index]
}

func (grid *Grid) GetAllNeighbours(hex IndexableHex) (out []IndexableHex) {
	index := hex.GetIndex()

	for _, i := range neighbours {
		hex, _ := grid.GetByCoords(index.x+i[0], index.y+i[1], index.z+i[2])
		if hex != nil {
			out = append(out, hex)
		}
	}

	return
}
