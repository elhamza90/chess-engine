package board

import (
	"fmt"
	"log"
	"unicode"
)

type Board struct {
	Pieces bbRepr

	State boardState
}

func (b Board) String(ply Player) (res string) {
	sep := "  +---+---+---+---+---+---+---+---+\n"
	var rows []rune
	var cols []rune
	if ply == WHITE {
		rows = []rune{'8', '7', '6', '5', '4', '3', '2', '1'}
		cols = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
	} else {
		rows = []rune{'1', '2', '3', '4', '5', '6', '7', '8'}
		cols = []rune{'H', 'G', 'F', 'E', 'D', 'C', 'B', 'A'}
	}
	var sq string
	var bbIndex byte
	var err error
	var pc Piece
	for _, r := range rows {
		res += sep
		res += fmt.Sprintf("%c ", r)
		for _, c := range cols {
			sq = string(unicode.ToLower(c)) + string(r)
			bbIndex, err = squareToIndex(sq)
			if err != nil {
				log.Print("Error finding Bitboard Index from Square string")
				return ""
			}
			//log.Printf("%c%c (%s) => %d", c, r, sq, bbIndex)
			if binaryIndexIsOne(b.Pieces.Empty, bbIndex) {
				res += fmt.Sprintf("|   ")
			} else {
				for _, pc = range []Piece{PAWN, KNIGHT, BISHOP, ROOK, QUEEN, KING} {
					if binaryIndexIsOne(b.Pieces.White[pc], bbIndex) {
						res += fmt.Sprintf("| %c ", pc.String(WHITE))
						break
					} else if binaryIndexIsOne(b.Pieces.Black[pc], bbIndex) {
						res += fmt.Sprintf("| %c ", pc.String(BLACK))
						break
					}
				}
				// Here should not be reached.
			}

		}
		res += "|\n"
	}
	res += sep
	res += "  "
	for _, c := range cols {
		res += fmt.Sprintf("  %c ", c)
	}
	res += "\n"
	return res
}
