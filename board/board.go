package board

import (
	"log"
	"strings"
	"unicode"
)

type Board struct {
	Pieces bbRepr

	State boardState
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
	var w_castling byte = 0
	var b_castling byte = 0
	log.Printf("Castling Rights: %s", parts[2])
	for _, b := range parts[2] {
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
	b.State.playersCastleRights[WHITE] = castleRights(w_castling)
	b.State.playersCastleRights[BLACK] = castleRights(b_castling)

	// Set En Passant Square
	if parts[3] != "-" {
		epIndex, err := squareToIndex(parts[3])
		if err != nil {
			return
		}
		b.State.epSquare = uint64(1) << epIndex
	}
}

// InitStandard sets the board
// to the initial position of standard chess
func (b Board) InitStandard() {
	//
}
