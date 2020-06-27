package board

// Type to gather possible moves of a player's pieces
type MoveMap map[Piece]uint64

// GenMoves returns a list of all legal moves
// for current player in current position
func (b Board) GenLegalMoves() (moves MoveMap) {
	var oppPlayer Player = b.State.oppPlayer() // Opponent Player
	var oppPlayerAttackSquares uint64 = 0
	moves = make(MoveMap)
	// add moves from King
	moves[KING] = b.Pieces.kingPseudoLegalMoves(b.State.currPlayer)
	oppPlayerAttackSquares |= b.Pieces.kingPseudoLegalMoves(oppPlayer)

	// add moves from Knights
	moves[KNIGHT] = b.Pieces.knightsPseudoLegalMoves(b.State.currPlayer)
	oppPlayerAttackSquares |= b.Pieces.knightsPseudoLegalMoves(oppPlayer)

	// add moves from Rooks
	moves[ROOK] = b.Pieces.rooksPseudoLegalMoves(b.State.currPlayer)
	oppPlayerAttackSquares |= b.Pieces.rooksPseudoLegalMoves(oppPlayer)

	// add moves from Bishops
	moves[BISHOP] = b.Pieces.bishopsPseudoLegalMoves(b.State.currPlayer)
	oppPlayerAttackSquares |= b.Pieces.bishopsPseudoLegalMoves(oppPlayer)

	// add moves from Queen
	moves[QUEEN] = b.Pieces.queenPseudoLegalMoves(b.State.currPlayer)
	oppPlayerAttackSquares |= b.Pieces.queenPseudoLegalMoves(oppPlayer)

	// add moves from Pawns
	moves[PAWN] = b.Pieces.pawnsPseudoLegalMoves(b.State.currPlayer)
	oppPlayerAttackSquares |= b.Pieces.pawnsCaptureMoves(oppPlayer)

	// Set inCheck if King is in check
	b.State.inCheck = false

	// Now Compute Legal Moves:
	//  - Case: King is in check / double check
	//  - Case: King can not move to squares attacked by other pieces
	//  - Case: Absolute pins (piece can't move if pinned by enemy piece)

	return moves
}
