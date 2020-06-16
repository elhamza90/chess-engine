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
	for i := len(fen) - 1; i >= 0; i-- {
		c := rune(fen[i])
		if c != '/' {
			//log.Print(string(c), ix)
			if unicode.IsDigit(c) {
				// Get number of empty squares
				inc, _ = strconv.Atoi(string(c))
				//log.Printf("There are %d zeros in the %dth index", nbrZeros, ix)
			} else {
				inc = 1
				// Check type and color of piece
				piece = Piece(unicode.ToUpper(c))
				pos = uint64(1) << ix
				if unicode.IsUpper(c) {
					whiteMap[piece] |= pos
				} else {
					blackMap[piece] |= pos
				}

			}
			// increment only when there is no separator
			// and by number of empty squares or by one piece
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
