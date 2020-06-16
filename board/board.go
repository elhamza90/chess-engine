package board

import (
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

	b.Pieces.Black = map[Piece]uint64{}
	b.Pieces.White = map[Piece]uint64{}
	res := uint64(0)
	//log.Printf("Pieces: %s", parts[0])
	var (
		ix       int = 0
		nbrZeros int = 0
		inc      int = 0
	)
	for i := len(parts[0]) - 1; i >= 0; i-- {
		c := rune(parts[0][i])
		if c != '/' {
			//log.Print(string(c), ix)
			if unicode.IsDigit(c) {
				// set "c" ones in the current index going to the left
				nbrZeros, _ = strconv.Atoi(string(c))
				for ii := 0; ii < nbrZeros; ii++ {
					res |= (uint64(1) << (ix + ii))
				}
				inc = nbrZeros
				//log.Printf("There are %d zeros in the %dth index", nbrZeros, ix)
				//log.Printf("%b", res)
			} else {
				inc = 1
				// Check type and color of piece
				p := Piece(unicode.ToUpper(c))
				pos := uint64(1) << ix
				if unicode.IsUpper(c) {
					b.Pieces.White[p] |= pos
				} else {
					b.Pieces.Black[p] |= pos
				}

			}
			// increment only when there is no separator
			// and by number of empty squares or by one piece
			ix += inc
		}
	}
	b.Pieces.Empty = res
}

// InitStandard sets the board
// to the initial position of standard chess
func (b Board) InitStandard() {
	//
}
