package board

// Define Players
type Player byte

const (
	WHITE Player = 'W'
	BLACK Player = 'B'
)

// Define Pieces
type Piece byte

const (
	PAWN   Piece = 'P'
	KNIGHT Piece = 'N'
	BISHOP Piece = 'B'
	ROOK   Piece = 'R'
	QUEEN  Piece = 'Q'
	KING   Piece = 'K'
)
