package board

import (
	"math"
	"testing"
)

func TestUtil_Piece_String(t *testing.T) {
	tests := map[string]struct {
		piece    Piece
		player   Player
		expected rune
	}{
		"White Pawn":   {PAWN, WHITE, 'P'},
		"Black Pawn":   {PAWN, BLACK, 'p'},
		"White Knight": {KNIGHT, WHITE, 'N'},
		"Black Knight": {KNIGHT, BLACK, 'n'},
		"White Bishop": {BISHOP, WHITE, 'B'},
		"Black Bishop": {BISHOP, BLACK, 'b'},
		"White Rook":   {ROOK, WHITE, 'R'},
		"Black Rook":   {ROOK, BLACK, 'r'},
		"White Queen":  {QUEEN, WHITE, 'Q'},
		"Black Queen":  {QUEEN, BLACK, 'q'},
		"White King":   {KING, WHITE, 'K'},
		"Black King":   {KING, BLACK, 'k'},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var p Piece = test.piece
			res := p.String(test.player)
			if res != test.expected {
				t.Errorf("Error Converting Piece to String: %s", name)
			}
		})
	}
}

func TestUtil_Square_fromString(t *testing.T) {
	tests := map[string]byte{
		"a1": 0,
		"a8": 56,
		"b1": 1,
		"e3": 20,
		"d5": 35,
		"h8": 63,
	}
	var res Square
	for sqStr, ix := range tests {
		t.Run("", func(t *testing.T) {
			err := res.fromString(sqStr)
			if byte(res) != ix || err != nil {
				t.Errorf("Expecting %d but Got %d for %s", ix, res, sqStr)
			}
		})
	}
}

func TestUtil_Bitboard_IsSet(t *testing.T) {
	bin := "1010000010000101011001001010000000000000000001001000011100010000"
	// TODO: make the conversion programatically. Here I did it offline and got the following decimal
	bitboard := Bitboard(11566761856328828688)
	tests := map[Square]bool{
		H8: true,
		G8: false,
		F8: true,
		H7: true,
		C7: true,
		C6: true,
		D4: false,
		A7: true,
		G6: true,
		F6: true,
		A2: true,
		B2: true,
		C2: true,
		E1: true,
		F1: false,
	}
	for sq, expected := range tests {
		res := bitboard.IsSet(sq)
		if res != expected {
			t.Errorf("Error IsSet for Square %d. Expected %t but Got %t (in %s)", sq, expected, res, bin)
		}
	}
}

func TestUtil_Bitboard_Set(t *testing.T) {
	tests := map[Square]Bitboard{
		H8: Bitboard(uint64(math.Pow(2, 63)) + uint64(16)),
		D5: Bitboard(uint64(math.Pow(2, 35)) + uint64(16)),
		A8: Bitboard(uint64(math.Pow(2, 56)) + uint64(16)),
		H1: Bitboard(uint64(math.Pow(2, 7)) + uint64(16)),
	}
	for sq, expected := range tests {
		bitboard := Bitboard(16) // Initial Bitboard is 10000
		bitboard.Set(sq)
		if bitboard != expected {
			t.Errorf("Error Setting the %s Square(%d) in bitboard. Expected %d but Got %d", sq.String(), sq, expected, bitboard)
		}

	}
}
