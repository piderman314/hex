package hex

type IndexableHex interface {
	GetIndex() hexIndex
	SetIndex(hexIndex)
}

type hexIndex struct {
	x, y, z int
}
