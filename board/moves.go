package board

// Map representing the Attacking Squares of each piece in it's current position
type AttackMap map[Piece]Bitboard

func (b Board) PlayerPseudoLegalMoves(ply Player) (moves AttackMap) {
	moves = make(AttackMap)
	moves[PAWN] = b.Pieces.pawnsPseudoLegalMoves(b.State.currPlayer)
	moves[KNIGHT] = b.Pieces.knightsPseudoLegalMoves(b.State.currPlayer)
	moves[BISHOP] = b.Pieces.bishopsPseudoLegalMoves(b.State.currPlayer)
	moves[ROOK] = b.Pieces.rooksPseudoLegalMoves(b.State.currPlayer)
	moves[QUEEN] = b.Pieces.queenPseudoLegalMoves(b.State.currPlayer)
	moves[KING] = b.Pieces.kingPseudoLegalMoves(b.State.currPlayer)
	return moves
}

// GenMoves returns a list of all legal moves
// for current player in current position
func (b Board) GenLegalMoves() (moves AttackMap) {
	var oppPlayer Player = b.State.oppPlayer() // Opponent Player
	var oppPlayerAttackSquares Bitboard = 0

	// Get Pseudo-legal moves of current player
	moves = b.PlayerPseudoLegalMoves(b.State.CurrPlayer())

	// Get All pseudo-legal moves of opponent and combine them in one bitboard
	for _, mvs := range b.PlayerPseudoLegalMoves(oppPlayer) {
		oppPlayerAttackSquares |= mvs
	}

	// Set inCheck if King is in check
	kingPosBitboard := b.Pieces.Positions[b.State.currPlayer][KING]
	var kingSquare Square
	for sq := A1; sq < H8; sq++ {
		if kingPosBitboard.IsSet(sq) {
			kingSquare = Square(sq)
			break
		}
	}
	b.State.inCheck = squareAttacked(kingSquare, oppPlayerAttackSquares)

	// Now Compute Legal Moves:
	//  - Case: King is in check / double check
	//  - Case: King can not move to squares attacked by other pieces
	//  - Case: Absolute pins (piece can't move if pinned by enemy piece)
	// TODO

	// Check if Casling is possible
	// TODO

	return moves
}

func squareAttacked(sq Square, attackBitboard Bitboard) bool {
	return attackBitboard.IsSet(sq)
}
