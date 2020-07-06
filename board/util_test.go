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

func TestUtil_Square_Up(t *testing.T) {
	testsNoErrors := map[Square]Square{
		A1: A2,
		D1: D2,
		H1: H2,
		F5: F6,
		B7: B8,
	}
	testsError := []Square{A8, B8, C8, D8, E8, F8, G8, H8}
	var res Square
	var err error
	for sq, expected := range testsNoErrors {
		if res, err = sq.Up(); err != nil {
			t.Errorf("Error getting Upper square of %s. Expected %s but got Error: %s", sq.String(), expected.String(), err)
		} else if res != expected {
			t.Errorf("Error getting Upper square of %s. Expected %s but got %s", sq.String(), expected.String(), res.String())
		}
	}
	for _, sq := range testsError {
		if res, err = sq.Up(); err == nil {
			t.Errorf("Error getting Upper square of %s. Expected Error but got %s", sq.String(), res.String())
		}
	}
}

func TestUtil_Square_Down(t *testing.T) {
	testsNoErrors := map[Square]Square{
		A8: A7,
		D4: D3,
		H7: H6,
		F2: F1,
		B7: B6,
	}
	testsError := []Square{A1, B1, C1, D1, E1, F1, G1, H1}
	var res Square
	var err error
	for sq, expected := range testsNoErrors {
		if res, err = sq.Down(); err != nil {
			t.Errorf("Error getting Lower square of %s. Expected %s but got Error: %s", sq.String(), expected.String(), err)
		} else if res != expected {
			t.Errorf("Error getting Lower square of %s. Expected %s but got %s", sq.String(), expected.String(), res.String())
		}
	}
	for _, sq := range testsError {
		if res, err = sq.Down(); err == nil {
			t.Errorf("Error getting Lower square of %s. Expected Error but got %s", sq.String(), res.String())
		}
	}
}

func TestUtil_Square_Left(t *testing.T) {
	testsNoErrors := map[Square]Square{
		D4: C4,
		H7: G7,
		F2: E2,
		B3: A3,
	}
	testsError := []Square{A1, A2, A3, A4, A5, A6, A7, A8}
	var res Square
	var err error
	for sq, expected := range testsNoErrors {
		if res, err = sq.Left(); err != nil {
			t.Errorf("Error getting Left square of %s. Expected %s but got Error: %s", sq.String(), expected.String(), err)
		} else if res != expected {
			t.Errorf("Error getting Left square of %s. Expected %s but got %s", sq.String(), expected.String(), res.String())
		}
	}
	for _, sq := range testsError {
		if res, err = sq.Left(); err == nil {
			t.Errorf("Error getting Left square of %s. Expected Error but got %s", sq.String(), res.String())
		}
	}
}

func TestUtil_Square_Right(t *testing.T) {
	testsNoErrors := map[Square]Square{
		D4: E4,
		G7: H7,
		F2: G2,
		B3: C3,
	}
	testsError := []Square{H1, H2, H3, H4, H5, H6, H7, H8}
	var res Square
	var err error
	for sq, expected := range testsNoErrors {
		if res, err = sq.Right(); err != nil {
			t.Errorf("Error getting Right square of %s. Expected %s but got Error: %s", sq.String(), expected.String(), err)
		} else if res != expected {
			t.Errorf("Error getting Right square of %s. Expected %s but got %s", sq.String(), expected.String(), res.String())
		}
	}
	for _, sq := range testsError {
		if res, err = sq.Right(); err == nil {
			t.Errorf("Error getting Right square of %s. Expected Error but got %s", sq.String(), res.String())
		}
	}
}

func TestUtil_Square_UpperRight(t *testing.T) {
	testsNoErrors := map[Square]Square{
		D4: E5,
		G7: H8,
		F2: G3,
		B3: C4,
	}
	testsError := []Square{H1, H2, H3, H4, H5, H6, H7, H8, G8, F8, E8, D8, C8, B8, A8}
	var res Square
	var err error
	for sq, expected := range testsNoErrors {
		if res, err = sq.UpperRight(); err != nil {
			t.Errorf("Error getting Upper-Right square of %s. Expected %s but got Error: %s", sq.String(), expected.String(), err)
		} else if res != expected {
			t.Errorf("Error getting Upper-Right square of %s. Expected %s but got %s", sq.String(), expected.String(), res.String())
		}
	}
	for _, sq := range testsError {
		if res, err = sq.UpperRight(); err == nil {
			t.Errorf("Error getting Upper-Right square of %s. Expected Error but got %s", sq.String(), res.String())
		}
	}
}

func TestUtil_Square_DownRight(t *testing.T) {
	testsNoErrors := map[Square]Square{
		D4: E3,
		G7: H6,
		F2: G1,
		B3: C2,
	}
	testsError := []Square{H1, H2, H3, H4, H5, H6, H7, H8, A1, B1, C1, D1, E1, F1, G1}
	var res Square
	var err error
	for sq, expected := range testsNoErrors {
		if res, err = sq.DownRight(); err != nil {
			t.Errorf("Error getting Down-Right square of %s. Expected %s but got Error: %s", sq.String(), expected.String(), err)
		} else if res != expected {
			t.Errorf("Error getting Down-Right square of %s. Expected %s but got %s", sq.String(), expected.String(), res.String())
		}
	}
	for _, sq := range testsError {
		if res, err = sq.DownRight(); err == nil {
			t.Errorf("Error getting Down-Right square of %s. Expected Error but got %s", sq.String(), res.String())
		}
	}
}

func TestUtil_Square_DownLeft(t *testing.T) {
	testsNoErrors := map[Square]Square{
		D4: C3,
		G7: F6,
		F2: E1,
		B3: A2,
	}
	testsError := []Square{A1, A2, A3, A4, A5, A6, A7, A8, B1, C1, D1, E1, F1, G1}
	var res Square
	var err error
	for sq, expected := range testsNoErrors {
		if res, err = sq.DownLeft(); err != nil {
			t.Errorf("Error getting Down-Left square of %s. Expected %s but got Error: %s", sq.String(), expected.String(), err)
		} else if res != expected {
			t.Errorf("Error getting Down-Left square of %s. Expected %s but got %s", sq.String(), expected.String(), res.String())
		}
	}
	for _, sq := range testsError {
		if res, err = sq.DownLeft(); err == nil {
			t.Errorf("Error getting Down-Left square of %s. Expected Error but got %s", sq.String(), res.String())
		}
	}
}

func TestUtil_Square_UpperLeft(t *testing.T) {
	testsNoErrors := map[Square]Square{
		D4: C5,
		G7: F8,
		F2: E3,
		B3: A4,
	}
	testsError := []Square{H8, G8, F8, E8, D8, C8, B8, A8, A1, A2, A3, A4, A5, A6, A7}
	var res Square
	var err error
	for sq, expected := range testsNoErrors {
		if res, err = sq.UpperLeft(); err != nil {
			t.Errorf("Error getting Upper-Left square of %s. Expected %s but got Error: %s", sq.String(), expected.String(), err)
		} else if res != expected {
			t.Errorf("Error getting Upper-Left square of %s. Expected %s but got %s", sq.String(), expected.String(), res.String())
		}
	}
	for _, sq := range testsError {
		if res, err = sq.UpperLeft(); err == nil {
			t.Errorf("Error getting Upper-Left square of %s. Expected Error but got %s", sq.String(), res.String())
		}
	}
}

func TestUtil_Square_KnightJumps(t *testing.T) {
	tests := map[Square][]Square{
		D4: []Square{C6, B5, C2, B3, E2, F3, F5, E6},
		G7: []Square{H5, E8, F5, E6},
		F2: []Square{H1, H3, G4, E4, D3, D1},
		B3: []Square{A1, C1, D2, C5, D4, A5},
		A1: []Square{B3, C2},
		H1: []Square{G3, F2},
	}
	var res []Square
	sumSquares := func(sqs []Square) int {
		s := 0
		for _, sq := range sqs {
			s += int(sq)
		}
		return s
	}
	for sq, expected := range tests {
		t.Run(sq.String(), func(t *testing.T) {
			res = sq.KnightJumps()
			resSum := sumSquares(res)
			expectedSum := sumSquares(expected)
			if resSum != expectedSum {
				t.Errorf("Error getting Knight Jump square of %s. Expected %v but got %v", sq.String(), expected, res)
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

func TestUtil_Bitboard_FromSquares(t *testing.T) {
	tests := map[Bitboard][]Square{
		Bitboard(1):                    []Square{A1},
		Bitboard(256):                  []Square{A2},
		Bitboard(257):                  []Square{A1, A2},
		Bitboard(9223372036854775808):  []Square{H8},
		Bitboard(18446462598732840960): []Square{A8, B8, C8, D8, E8, F8, G8, H8, A7, B7, C7, D7, E7, F7, G7, H7},
		Bitboard(18446462598732906495): []Square{A8, B8, C8, D8, E8, F8, G8, H8, A7, B7, C7, D7, E7, F7, G7, H7, A2, B2, C2, D2, E2, F2, G2, H2, A1, B1, C1, D1, E1, F1, G1, H1},
	}
	var res Bitboard
	for expected, squares := range tests {
		res.FromSquares(squares)
		if res != expected {
			t.Errorf("Error constructing Bitboard from Squares %v. Expected %d but got %d", squares, expected, res)
		}
	}
}
