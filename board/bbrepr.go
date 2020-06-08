package board

type bbRepr struct {
	White map[Piece]uint64
	Black map[Piece]uint64
	Empty uint64
}
