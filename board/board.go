package board

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Board struct {
	Pieces bbRepr

	State boardState
}

func fenCastleRights(s string) (castleRights, castleRights) {
	//log.Printf("Castling Rights: %s", s)
	var w_castling byte = 0
	var b_castling byte = 0
	for _, b := range s {
		if b == 'K' {
			w_castling += byte(CASTLE_KING)
		} else if b == 'Q' {
			w_castling += byte(CASTLE_QUEEN)
		} else if b == 'k' {
			b_castling += byte(CASTLE_KING)
		} else if b == 'q' {
			b_castling += byte(CASTLE_QUEEN)
		}
	}
	return castleRights(w_castling), castleRights(b_castling)
}

func fenToBitboardPieces(fen string) (map[Piece]uint64, map[Piece]uint64) {
	blackMap := map[Piece]uint64{}
	whiteMap := map[Piece]uint64{}
	//log.Printf("Pieces: %s", parts[0])
	var (
		ix    int    = 0
		inc   int    = 0
		pos   uint64 = 0
		piece Piece
	)
	rows := strings.Split(fen, "/")
	for r := len(rows) - 1; r >= 0; r-- {
		row := rows[r]
		//log.Print(row)
		for c := 0; c < len(row); c++ {
			char := rune(row[c])
			if unicode.IsDigit(char) {
				// Get number of empty squares
				inc, _ = strconv.Atoi(string(char))
				//log.Printf("There are %d zeros in the %dth index", inc, ix)
			} else {
				inc = 1
				// Check type and color of piece
				piece = Piece(unicode.ToUpper(char))
				pos = uint64(1) << ix
				if unicode.IsUpper(char) {
					whiteMap[piece] |= pos
				} else {
					blackMap[piece] |= pos
				}

			}
			ix += inc
		}
	}
	return whiteMap, blackMap
}

// FromFen sets the state of the board
// to the state given by the FEN representation
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
		epIndex, err := squareToIndex(parts[3])
		if err != nil {
			return
		}
		b.State.epSquare = uint64(1) << epIndex
	}

	// Set Pieces locations in Bitboards
	b.Pieces.White, b.Pieces.Black = fenToBitboardPieces(parts[0])

	// Calculate Empty Squares from pieces locations
	var occupied uint64 = 0
	for _, pos := range b.Pieces.Black {
		occupied += pos
	}
	for _, pos := range b.Pieces.White {
		occupied += pos
	}
	b.Pieces.Empty = ^occupied

}

// InitStandard sets the board
// to the initial position of standard chess
func (b Board) InitStandard() {
	//
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
