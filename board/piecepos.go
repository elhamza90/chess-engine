package board

/************************************************************/

type piecePos struct {
	Positions map[Player]map[Piece]uint64
	Empty     uint64
}

func binaryIndexIsOne(bb uint64, ix byte) bool {
	return bb&(uint64(1)<<ix) != 0
}

/******************** Move Generation ****************************/

// kingPseudoLegalMoves returns a bitboard containing all possible
// moves for the King considering current/opponent pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (bbs piecePos) kingPseudoLegalMoves(ply Player) (res uint64) {
	// TODO
	return res
}

// knightsPseudoLegalMoves returns a bitboard containing all possible
// moves for the King considering current/opponent pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (bbs piecePos) knightsPseudoLegalMoves(ply Player) (res uint64) {
	// TODO
	return res
}

// rooksPseudoLegalMoves returns a bitboard containing all possible
// moves for the King considering current/opponent pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (bbs piecePos) rooksPseudoLegalMoves(ply Player) (res uint64) {
	// TODO
	return res
}

// bishopsPseudoLegalMoves returns a bitboard containing all possible
// moves for the King considering current/opponent pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (bbs piecePos) bishopsPseudoLegalMoves(ply Player) (res uint64) {
	// TODO
	return res
}

// queenPseudoLegalMoves returns a bitboard containing all possible
// moves for the Queen considering current/opponent pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (bbs piecePos) queenPseudoLegalMoves(ply Player) (res uint64) {
	// TODO
	return res
}

// pawnsCaptureMoves returns a bitboard containing all possible attacking
// moves for the pawns (side attacks) considering current/opponent pieces positions.
// There are 2 types of attacking moves:
//   - side captures
//   - en-passant captures
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (bbs piecePos) pawnsCaptureMoves(ply Player) (res uint64) {
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
func (bbs piecePos) pawnsPseudoLegalMoves(ply Player) (res uint64) {
	// First get the capture moves
	res |= bbs.pawnsCaptureMoves(ply)
	// Next get push moves
	// TODO
	return res
}
