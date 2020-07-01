package board

/************************************************************/

type PiecePositions map[Piece]Bitboard

type PlayerPiecePositions map[Player]PiecePositions

func (pp PiecePositions) All() (res Bitboard) {
	for _, pos := range pp {
		res |= pos
	}
	return res
}

func (ppp PlayerPiecePositions) Empty() (res Bitboard) {
	for _, piecePositions := range ppp {
		res |= piecePositions.All()
	}
	return ^res
}

/******************** Move Generation ****************************/

// kingPseudoLegalMoves returns a bitboard containing all possible
// moves for the King considering current/opponent pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (ppp PlayerPiecePositions) kingPseudoLegalMoves(ply Player) (res Bitboard) {
	// WIP
	// separate friendly and opponent pieces
	//var opponent Player = ply.opponent()

	return res
}

// knightsPseudoLegalMoves returns a bitboard containing all possible
// moves for the King considering current/opponent pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (ppp PlayerPiecePositions) knightsPseudoLegalMoves(ply Player) (res Bitboard) {
	// TODO
	return res
}

// rooksPseudoLegalMoves returns a bitboard containing all possible
// moves for the King considering current/opponent pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (ppp PlayerPiecePositions) rooksPseudoLegalMoves(ply Player) (res Bitboard) {
	// TODO
	return res
}

// bishopsPseudoLegalMoves returns a bitboard containing all possible
// moves for the King considering current/opponent pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (ppp PlayerPiecePositions) bishopsPseudoLegalMoves(ply Player) (res Bitboard) {
	// TODO
	return res
}

// queenPseudoLegalMoves returns a bitboard containing all possible
// moves for the Queen considering current/opponent pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (ppp PlayerPiecePositions) queenPseudoLegalMoves(ply Player) (res Bitboard) {
	// TODO
	return res
}

// pawnsCaptureMoves returns a bitboard containing all possible attacking
// moves for the pawns (side attacks) considering current/opponent pieces positions.
// There are 2 types of attacking moves:
//   - side captures
//   - en-passant captures
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (ppp PlayerPiecePositions) pawnsCaptureMoves(ply Player) (res Bitboard) {
	// TODO
	return res
}

// pawnsPseudoLegalMoves returns a bitboard containing all possible
// moves for the Pawns considering current/opponent pieces positions.
// There are 3 types of moves for pawns:
//   - One square pushes
//   - Two square pushes
//   - Capture Moves:
//		- side captures
//		- en-passant captures
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (ppp PlayerPiecePositions) pawnsPseudoLegalMoves(ply Player) (res Bitboard) {
	// First get the capture moves
	res |= ppp.pawnsCaptureMoves(ply)
	// Next get push moves
	// TODO
	return res
}
