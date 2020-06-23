package board

import (
	"testing"
)

func TestBoard_String(t *testing.T) {
	// Test Position
	b := Board{
		State: boardState{
			currPlayer: WHITE,
			inCheck:    false,
			playersCastleRights: map[Player]castleRights{
				WHITE: CASTLE_KING + CASTLE_QUEEN,
				BLACK: CASTLE_KING + CASTLE_QUEEN,
			},
			epSquare: 0,
		},
		Pieces: bbRepr{
			White: map[Piece]uint64{
				'P': uint64(65280),
				'R': uint64(129),
				'N': uint64(66),
				'B': uint64(36),
				'K': uint64(16),
				'Q': uint64(8),
			},
			Black: map[Piece]uint64{
				'P': uint64(71776119061217280),
				'R': uint64(9295429630892703744),
				'N': uint64(4755801206503243776),
				'B': uint64(2594073385365405696),
				'K': uint64(1152921504606846976),
				'Q': uint64(576460752303423488),
			},
			Empty: 281474976645120,
		},
	}
	res := b.String(WHITE)
	expected := `  +---+---+---+---+---+---+---+---+
8 | r | n | b | q | k | b | n | r |
  +---+---+---+---+---+---+---+---+
7 | p | p | p | p | p | p | p | p |
  +---+---+---+---+---+---+---+---+
6 |   |   |   |   |   |   |   |   |
  +---+---+---+---+---+---+---+---+
5 |   |   |   |   |   |   |   |   |
  +---+---+---+---+---+---+---+---+
4 |   |   |   |   |   |   |   |   |
  +---+---+---+---+---+---+---+---+
3 |   |   |   |   |   |   |   |   |
  +---+---+---+---+---+---+---+---+
2 | P | P | P | P | P | P | P | P |
  +---+---+---+---+---+---+---+---+
1 | R | N | B | Q | K | B | N | R |
  +---+---+---+---+---+---+---+---+
    A   B   C   D   E   F   G   H 
`
	if res != expected {
		t.Errorf("Error Converting board to string (WHITE). Got:\n%s\nExpected:\n%s", res, expected)
	}

	res = b.String(BLACK)
	expected = `  +---+---+---+---+---+---+---+---+
1 | R | N | B | K | Q | B | N | R |
  +---+---+---+---+---+---+---+---+
2 | P | P | P | P | P | P | P | P |
  +---+---+---+---+---+---+---+---+
3 |   |   |   |   |   |   |   |   |
  +---+---+---+---+---+---+---+---+
4 |   |   |   |   |   |   |   |   |
  +---+---+---+---+---+---+---+---+
5 |   |   |   |   |   |   |   |   |
  +---+---+---+---+---+---+---+---+
6 |   |   |   |   |   |   |   |   |
  +---+---+---+---+---+---+---+---+
7 | p | p | p | p | p | p | p | p |
  +---+---+---+---+---+---+---+---+
8 | r | n | b | k | q | b | n | r |
  +---+---+---+---+---+---+---+---+
    H   G   F   E   D   C   B   A 
`
	if res != expected {
		t.Errorf("Error Converting board to string (BLACK). Got:\n%s\nExpected:\n%s", res, expected)
	}
}
