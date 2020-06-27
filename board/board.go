package board

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

type Board struct {
	Pieces piecePos

	State boardState
}

// String returns a string representation of the board
// that can be printed in the terminal
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
	var bbIndex Square
	var err error
	var pc Piece
	for _, r := range rows {
		res += sep
		res += fmt.Sprintf("%c ", r)
		for _, c := range cols {
			sq = string(unicode.ToLower(c)) + string(r)
			err = bbIndex.fromString(sq)
			if err != nil {
				log.Print("Error finding Bitboard Index from Square string")
				return ""
			}
			//log.Printf("%c%c (%s) => %d", c, r, sq, bbIndex)
			if b.Pieces.Empty.isSet(bbIndex) {
				res += fmt.Sprintf("|   ")
			} else {
				for _, pc = range []Piece{PAWN, KNIGHT, BISHOP, ROOK, QUEEN, KING} {
					if b.Pieces.Positions[WHITE][pc].isSet(bbIndex) {
						res += fmt.Sprintf("| %c ", pc.String(WHITE))
						break
					} else if b.Pieces.Positions[BLACK][pc].isSet(bbIndex) {
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

// FromFen sets the state of the board and piece positions
// to the state and positions given by the FEN representation
func (b *Board) FromFen(fen string) {
	if err := fenIsValid(fen); err != nil {
		return
	}
	parts := strings.Split(fen, " ")

	// Set Current Player
	if len(parts[1]) == 1 {
		upper := unicode.ToUpper(rune(parts[1][0]))
		//log.Print(Player(upper))
		b.State.currPlayer = Player(upper)
	}

	// Set Castling rights
	b.State.playersCastleRights = map[Player]castleRights{}
	b.State.playersCastleRights[WHITE], b.State.playersCastleRights[BLACK] = fenCastleRights(parts[2])

	// Set En Passant Square
	if parts[3] != "-" {
		var epIndex Square
		err := epIndex.fromString(parts[3])
		if err != nil {
			return
		}
		b.State.epSquare = Bitboard(1) << epIndex
	}

	// Set Pieces locations in Bitboards
	b.Pieces.Positions = make(map[Player]map[Piece]Bitboard)
	b.Pieces.Positions[WHITE], b.Pieces.Positions[BLACK] = fenToBitboardPieces(parts[0])

	// Calculate Empty Squares from pieces locations
	var occupied Bitboard = 0
	for _, pos := range b.Pieces.Positions[BLACK] {
		occupied += pos
	}
	for _, pos := range b.Pieces.Positions[WHITE] {
		occupied += pos
	}
	b.Pieces.Empty = ^occupied

}
