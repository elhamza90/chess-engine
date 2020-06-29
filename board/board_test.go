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
		Pieces: piecePos{
			Positions: map[Player]map[Piece]Bitboard{
				WHITE: map[Piece]Bitboard{
					'P': Bitboard(65280),
					'R': Bitboard(129),
					'N': Bitboard(66),
					'B': Bitboard(36),
					'K': Bitboard(16),
					'Q': Bitboard(8),
				},
				BLACK: map[Piece]Bitboard{
					'P': Bitboard(71776119061217280),
					'R': Bitboard(9295429630892703744),
					'N': Bitboard(4755801206503243776),
					'B': Bitboard(2594073385365405696),
					'K': Bitboard(1152921504606846976),
					'Q': Bitboard(576460752303423488),
				},
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

func TestFen_BoardFromFen(t *testing.T) {
	tests := map[string]Board{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": {
			State: boardState{
				currPlayer: WHITE,
				inCheck:    false,
				playersCastleRights: map[Player]castleRights{
					WHITE: CASTLE_KING + CASTLE_QUEEN,
					BLACK: CASTLE_KING + CASTLE_QUEEN,
				},
				epSquare: 0,
			},
			Pieces: piecePos{
				Positions: map[Player]map[Piece]Bitboard{
					WHITE: map[Piece]Bitboard{
						'P': Bitboard(65280),
						'R': Bitboard(129),
						'N': Bitboard(66),
						'B': Bitboard(36),
						'K': Bitboard(16),
						'Q': Bitboard(8),
					},
					BLACK: map[Piece]Bitboard{
						'P': Bitboard(71776119061217280),
						'R': Bitboard(9295429630892703744),
						'N': Bitboard(4755801206503243776),
						'B': Bitboard(2594073385365405696),
						'K': Bitboard(1152921504606846976),
						'Q': Bitboard(576460752303423488),
					},
				},
				Empty: 281474976645120,
			},
		},
		"r2q1bnr/pp1bkppp/2n5/3pp3/2Pp3P/1P3N2/PB1NPPP1/2RQKB1R b K c3 0 8": {
			State: boardState{
				currPlayer: BLACK,
				inCheck:    false,
				playersCastleRights: map[Player]castleRights{
					WHITE: CASTLE_KING,
					BLACK: CASTLE_NONE,
				},
				epSquare: Bitboard(262144),
			},
			Pieces: piecePos{
				Positions: map[Player]map[Piece]Bitboard{
					WHITE: map[Piece]Bitboard{
						'P': Bitboard(2214752512),
						'R': Bitboard(132),
						'N': Bitboard(2099200),
						'B': Bitboard(544),
						'K': Bitboard(16),
						'Q': Bitboard(8),
					},
					BLACK: map[Piece]Bitboard{
						'P': Bitboard(63894922926751744),
						'R': Bitboard(9295429630892703744),
						'N': Bitboard(4611690416473899008),
						'B': Bitboard(2308094809027379200),
						'K': Bitboard(4503599627370496),
						'Q': Bitboard(576460752303423488),
					},
				},
				Empty: Bitboard(1586669940241171523),
			},
		},
		"rnbqk1nr/pppp1ppp/8/4P3/1b6/8/PPP1PPPP/RNBQKBNR w KQkq - 1 3": {
			State: boardState{
				currPlayer: WHITE,
				inCheck:    true,
				playersCastleRights: map[Player]castleRights{
					WHITE: CASTLE_KING + CASTLE_QUEEN,
					BLACK: CASTLE_KING + CASTLE_QUEEN,
				},
				epSquare: 0,
			},
			Pieces: piecePos{
				Positions: map[Player]map[Piece]Bitboard{
					WHITE: map[Piece]Bitboard{
						'P': Bitboard(68719539968),
						'R': Bitboard(129),
						'N': Bitboard(66),
						'B': Bitboard(36),
						'K': Bitboard(16),
						'Q': Bitboard(8),
					},
					BLACK: map[Piece]Bitboard{
						'P': Bitboard(67272519433846784),
						'R': Bitboard(9295429630892703744),
						'N': Bitboard(4755801206503243776),
						'B': Bitboard(288230376185266176),
						'K': Bitboard(1152921504606846976),
						'Q': Bitboard(576460752303423488),
					},
				},
				Empty: Bitboard(2310628015064680448),
			},
		},
	}
	for fen, expected := range tests {
		t.Run(fen, func(t *testing.T) {
			b := Board{}
			b.FromFen(fen)
			if b.State.CurrPlayer() != expected.State.CurrPlayer() {
				t.Errorf("Error setting current player from fen( expected %s  but got %s)", string(expected.State.currPlayer), string(b.State.currPlayer))
			}

			castleRights := b.State.PlayersCastleRights()
			if castleRights == nil {
				t.Errorf("Error setting players castling rights from fen.")
			} else if castleRights[WHITE] != expected.State.playersCastleRights[WHITE] {
				t.Errorf("Error setting WHITE player castling rights from fen")
			} else if castleRights[BLACK] != expected.State.playersCastleRights[BLACK] {
				t.Errorf("Error setting BLACK player castling rights from fen")
			}

			if b.State.EpSquare() != expected.State.epSquare {
				t.Errorf("Error setting en passant square. Expecting %d, got %d", expected.State.epSquare, b.State.EpSquare())
			}

			if b.Pieces.Empty != expected.Pieces.Empty {
				t.Errorf("Error in Empty squares. Expecting %b but got %b", expected.Pieces.Empty, b.Pieces.Empty)
			}

			for p, expectedPos := range expected.Pieces.Positions[WHITE] {
				resultPos := b.Pieces.Positions[WHITE][p]
				if expectedPos != resultPos {
					t.Errorf("Error setting White Piece %c position. Expected %b but got %b", p, expectedPos, resultPos)
				}
			}

			for p, expectedPos := range expected.Pieces.Positions[BLACK] {
				resultPos := b.Pieces.Positions[BLACK][p]
				if expectedPos != resultPos {
					t.Errorf("Error setting Black Piece %c position. Expected %b but got %b", p, expectedPos, resultPos)
				}
			}

		})
	}
}
