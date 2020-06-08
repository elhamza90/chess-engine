package board

type Board struct {
	Pieces bbRepr

	State boardState
}

// InitStandard sets the board
// to the initial position of standard chess
func (b Board) InitStandard() {
	//
}
