package board

/************************************************************/

type PiecePositions map[Piece]Bitboard

type PlayerPiecePositions map[Player]PiecePositions

// All combines all bitboards of all pieces into a single bitboard
func (pp PiecePositions) All() (res Bitboard) {
	for _, pos := range pp {
		res |= pos
	}
	return res
}

// Empty return a bitboard representing empty squares in the board
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
	playerKing := ppp[ply][KING]
	playerPieces := ppp[ply].All() ^ playerKing // excluding king
	/*kingMoveRules := map[Square]Bitboard{
		A1: Bitboard(770),
		A2: Bitboard(0),
		A3: Bitboard(0),
		A4: Bitboard(0),
		A5: Bitboard(0),
		A6: Bitboard(0),
		A7: Bitboard(0),
		A8: Bitboard(144959613005987840),
		B1: Bitboard(0),
		B2: Bitboard(0),
		B3: Bitboard(0),
		B4: Bitboard(0),
		B5: Bitboard(0),
		B6: Bitboard(0),
		B7: Bitboard(0),
		B8: Bitboard(0),
		C1: Bitboard(0),
		C2: Bitboard(0),
		C3: Bitboard(0),
		C4: Bitboard(0),
		C5: Bitboard(0),
		C6: Bitboard(0),
		C7: Bitboard(0),
		C8: Bitboard(0),
		D1: Bitboard(0),
		D2: Bitboard(0),
		D3: Bitboard(0),
		D4: Bitboard(120596463616),
		D5: Bitboard(0),
		D6: Bitboard(0),
		D7: Bitboard(0),
		D8: Bitboard(0),
		E1: Bitboard(0),
		E2: Bitboard(0),
		E3: Bitboard(0),
		E4: Bitboard(0),
		E5: Bitboard(0),
		E6: Bitboard(0),
		E7: Bitboard(0),
		E8: Bitboard(0),
		F1: Bitboard(0),
		F2: Bitboard(0),
		F3: Bitboard(0),
		F4: Bitboard(0),
		F5: Bitboard(0),
		F6: Bitboard(0),
		F7: Bitboard(0),
		F8: Bitboard(0),
		G1: Bitboard(0),
		G2: Bitboard(0),
		G3: Bitboard(0),
		G4: Bitboard(0),
		G5: Bitboard(0),
		G6: Bitboard(0),
		G7: Bitboard(0),
		G8: Bitboard(0),
		H1: Bitboard(49216),
		H2: Bitboard(0),
		H3: Bitboard(0),
		H4: Bitboard(0),
		H5: Bitboard(0),
		H6: Bitboard(0),
		H7: Bitboard(0),
		H8: Bitboard(4665729213955833856),
	}*/
	squares := make([]Square, 0, 8)
	for sq := A1; sq <= H8; sq++ {
		if playerKing.IsSet(sq) {
			//res = kingMoveRules[sq]
			if left, err := sq.Left(); err == nil {
				squares = append(squares, left)
			}
			if right, err := sq.Right(); err == nil {
				squares = append(squares, right)
			}
			if up, err := sq.Up(); err == nil {
				squares = append(squares, up)
			}
			if down, err := sq.Down(); err == nil {
				squares = append(squares, down)
			}
			if downLeft, err := sq.DownLeft(); err == nil {
				squares = append(squares, downLeft)
			}
			if downRight, err := sq.DownRight(); err == nil {
				squares = append(squares, downRight)
			}
			if upLeft, err := sq.UpperLeft(); err == nil {
				squares = append(squares, upLeft)
			}
			if upRight, err := sq.UpperRight(); err == nil {
				squares = append(squares, upRight)
			}
			break
		}
	}
	res.FromSquares(squares)
	res ^= playerPieces
	return res
}

// knightsPseudoLegalMoves returns a bitboard containing all possible
// moves for the Knights of a player considering current/opponent player pieces positions.
// Pseudo-Legal moves don't consider pins and attacks on the King.
func (ppp PlayerPiecePositions) knightsPseudoLegalMoves(ply Player) (res Bitboard) {
	// WIP
	playerKnights := ppp[ply][KNIGHT]
	playerPieces := ppp[ply].All() ^ playerKnights

	squares := make([]Square, 0, 8)
	for sq := A1; sq <= H8; sq++ {
		if playerKnights.IsSet(sq) {
			// find squares the knight can jump to
			knightJumps := sq.KnightJumps()
			squares = append(squares, knightJumps...)
		}
	}
	res.FromSquares(squares)
	res ^= playerPieces
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
