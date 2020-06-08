package board

// FromFen sets the state of the board
// to the state given by the FEN representation
func (b Board) FromFen(fen string) {
	if ok := fenIsValid(fen); !ok {
		return
	}
	// TODO
}

func fenIsValid(fen string) bool {
	return false
}
