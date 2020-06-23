package board

type Move struct {
	piece      Piece
	destSquare Square
}

// GenMoves returns a list of all pseudo legal moves
// for current player specified in Board state
func (b Board) GenMoves() []Move {
	moves := []Move{}
	return moves
}
