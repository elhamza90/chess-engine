package board

import (
	"regexp"
)

/******************** Definition Piece ****************************/

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

func (p Piece) String(ply Player) rune {
	switch p {
	case PAWN:
		if ply == WHITE {
			return 'P'
		} else {
			return 'p'
		}
	case KNIGHT:
		if ply == WHITE {
			return 'N'
		} else {
			return 'n'
		}
	case BISHOP:
		if ply == WHITE {
			return 'B'
		} else {
			return 'b'
		}
	case ROOK:
		if ply == WHITE {
			return 'R'
		} else {
			return 'r'
		}
	case QUEEN:
		if ply == WHITE {
			return 'Q'
		} else {
			return 'q'
		}
	case KING:
		if ply == WHITE {
			return 'K'
		} else {
			return 'k'
		}
	default:
		return 'X'
	}
}

/******************** Definition Square ****************************/
// Define Square Indices in a bitboard A1 = 0, B1 = 1, ... H8 = 63
type Square byte

const (
	A1 Square = iota
	B1 Square = iota
	C1 Square = iota
	D1 Square = iota
	E1 Square = iota
	F1 Square = iota
	G1 Square = iota
	H1 Square = iota
	A2 Square = iota
	B2 Square = iota
	C2 Square = iota
	D2 Square = iota
	E2 Square = iota
	F2 Square = iota
	G2 Square = iota
	H2 Square = iota
	A3 Square = iota
	B3 Square = iota
	C3 Square = iota
	D3 Square = iota
	E3 Square = iota
	F3 Square = iota
	G3 Square = iota
	H3 Square = iota
	A4 Square = iota
	B4 Square = iota
	C4 Square = iota
	D4 Square = iota
	E4 Square = iota
	F4 Square = iota
	G4 Square = iota
	H4 Square = iota
	A5 Square = iota
	B5 Square = iota
	C5 Square = iota
	D5 Square = iota
	E5 Square = iota
	F5 Square = iota
	G5 Square = iota
	H5 Square = iota
	A6 Square = iota
	B6 Square = iota
	C6 Square = iota
	D6 Square = iota
	E6 Square = iota
	F6 Square = iota
	G6 Square = iota
	H6 Square = iota
	A7 Square = iota
	B7 Square = iota
	C7 Square = iota
	D7 Square = iota
	E7 Square = iota
	F7 Square = iota
	G7 Square = iota
	H7 Square = iota
	A8 Square = iota
	B8 Square = iota
	C8 Square = iota
	D8 Square = iota
	E8 Square = iota
	F8 Square = iota
	G8 Square = iota
	H8 Square = iota
)

// squareToBitBoard takes a square (eg. b2)
// and returns the index of the square
// in a flattened board (eg. b2 -> 9)
func (sq *Square) fromString(sqStr string) error {
	if match, err := regexp.MatchString("^[a-h]{1}[1-9]{1}$", sqStr); (err != nil) || !match {
		return err
	}
	m := map[string]Square{
		"a1": A1, "a2": A2, "a3": A3, "a4": A4, "a5": A5, "a6": A6, "a7": A7, "a8": A8,
		"b1": B1, "b2": B2, "b3": B3, "b4": B4, "b5": B5, "b6": B6, "b7": B7, "b8": B8,
		"c1": C1, "c2": C2, "c3": C3, "c4": C4, "c5": C5, "c6": C6, "c7": C7, "c8": C8,
		"d1": D1, "d2": D2, "d3": D3, "d4": D4, "d5": D5, "d6": D6, "d7": D7, "d8": D8,
		"e1": E1, "e2": E2, "e3": E3, "e4": E4, "e5": E5, "e6": E6, "e7": E7, "e8": E8,
		"f1": F1, "f2": F2, "f3": F3, "f4": F4, "f5": F5, "f6": F6, "f7": F7, "f8": F8,
		"g1": G1, "g2": G2, "g3": G3, "g4": G4, "g5": G5, "g6": G6, "g7": G7, "g8": G8,
		"h1": H1, "h2": H2, "h3": H3, "h4": H4, "h5": H5, "h6": H6, "h7": H7, "h8": H8,
	}
	*sq = m[sqStr]
	return nil
}

func (sq *Square) String() string {
	m := map[Square]string{
		A1: "a1", A2: "A2", A3: "A3", A4: "A4", A5: "A5", A6: "A6", A7: "A7", A8: "A8",
		B1: "b1", B2: "B2", B3: "B3", B4: "B4", B5: "B5", B6: "B6", B7: "B7", B8: "B8",
		C1: "c1", C2: "C2", C3: "C3", C4: "C4", C5: "C5", C6: "C6", C7: "C7", C8: "C8",
		D1: "d1", D2: "D2", D3: "D3", D4: "D4", D5: "D5", D6: "D6", D7: "D7", D8: "D8",
		E1: "e1", E2: "E2", E3: "E3", E4: "E4", E5: "E5", E6: "E6", E7: "E7", E8: "E8",
		F1: "f1", F2: "f2", F3: "F3", F4: "F4", F5: "F5", F6: "F6", F7: "F7", F8: "F8",
		G1: "g1", G2: "g2", G3: "G3", G4: "G4", G5: "G5", G6: "G6", G7: "G7", G8: "G8",
		H1: "h1", H2: "h2", H3: "H3", H4: "H4", H5: "H5", H6: "H6", H7: "H7", H8: "H8",
	}
	return m[*sq]
}

/******************** Definition Bitboard ****************************/

type Bitboard uint64

// isSet returns true if the given square (index) is set to one
func (bb Bitboard) IsSet(sq Square) bool {
	return bb&(Bitboard(1)<<sq) != 0
}

// Set sets the given square to one
func (bb *Bitboard) Set(sq Square) {
	*bb |= Bitboard(1) << sq
}

// FromString constructs a Bitboard from a String of a binary number
func (bb *Bitboard) FromSquares(squares []Square) {
	*bb = Bitboard(0)
	for _, sq := range squares {
		bb.Set(sq)
	}
}
