package board

import "testing"

func TestFen_fenIsValid(t *testing.T) {
	validFens := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"r1bq1rk1/ppppn1pp/5n2/5p2/1bPP4/2NQ3N/PP1BPPBP/R3K2R b KQ - 9 9",
		"r1bq1rk1/pppp2pp/5nn1/5p2/2PP4/2PQ3N/P2BPPBP/R3K1R1 b Q - 2 11",
		"rnbqkr2/pppp2pp/5n2/5p2/1bPP4/2NQ4/PP2PP1P/R1B1KBNR w KQq - 4 7",
		"5r1k/pp4p1/3p3p/1R5n/3PP3/2P2n2/P4N1P/2B3K1 w - - 8 31",
		"rnbq1rk1/pppp2pp/5n2/8/1bPPPp2/2NQ4/PP3PBP/R1B1K1NR b KQ e3 0 8",
	}
	invalidFens := map[string]string{
		"Wrong Format: missing player":                 "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR KQkq - 0 1",
		"Wrong Format: missing en passant square":      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq 0 1",
		"Wrong format: missing one rank":               "rnbqkbnr/pppppppp/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"Wrong Half Clock":                             "rnbqkrnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - X 1",
		"Wrong Full Clock":                             "rnbqkrnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 Z",
		"Wrong player to move":                         "r1bq1rk1/ppppn1pp/5n2/5p2/1bPP4/2NQ3N/PP1BPPBP/R3K2R x KQ - 9 9",
		"Wrong number of empty squares in 5th rank":    "r1bq1rk1/pppp2pp/5nn1/5p3/2PP4/2PQ3N/P2BPPBP/R3K1R1 b Q - 2 11",
		"Wrong Castling rights":                        "rnbqkr2/pppp2pp/5n2/5p2/1bPP4/2NQ4/PP2PP1P/R1B1KBNR w KRq - 4 7",
		"Wrong En Passant(square doesnt exist)":        "rnbq1rk1/pppp2pp/5n2/8/1bPPPp2/2NQ4/PP3PBP/R1B1K1NR b KQ e9 0 8",
		"Wrong En Passant(square cant be en passant) ": "rnbq1rk1/pppp2pp/5n2/8/1bPPPp2/2NQ4/PP3PBP/R1B1K1NR b KQ e4 0 8",
	}

	for _, fen := range validFens {
		t.Run("Valid Fens", func(t *testing.T) {
			if err := fenIsValid(fen); err != nil {
				t.Errorf("Valid FEN was rejected: %s (%s)", fen, err)
			}
		})
	}

	for name, fen := range invalidFens {
		t.Run(name, func(t *testing.T) {
			if err := fenIsValid(fen); err == nil {
				t.Errorf("Invalid Fen was accepted: %s", fen)
			}
		})
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
		},
		"r2q1bnr/pp1bkppp/2n5/3pp3/2Pp3P/1P3N2/PB1NPPP1/2RQKB1R b K c3 0 8": {
			State: boardState{
				currPlayer: BLACK,
				inCheck:    false,
				playersCastleRights: map[Player]castleRights{
					WHITE: CASTLE_KING,
					BLACK: CASTLE_NONE,
				},
				epSquare: uint64(262144),
			},
			Pieces: bbRepr{
				White: map[Piece]uint64{
					'P': uint64(2214752512),
					'R': uint64(132),
					'N': uint64(2099200),
					'B': uint64(544),
					'K': uint64(16),
					'Q': uint64(8),
				},
				Black: map[Piece]uint64{
					'P': uint64(63894922926751744),
					'R': uint64(9295429630892703744),
					'N': uint64(4611690416473899008),
					'B': uint64(2308094809027379200),
					'K': uint64(4503599627370496),
					'Q': uint64(576460752303423488),
				},
				Empty: uint64(1586669940241171523),
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
			Pieces: bbRepr{
				White: map[Piece]uint64{
					'P': uint64(68719539968),
					'R': uint64(129),
					'N': uint64(66),
					'B': uint64(36),
					'K': uint64(16),
					'Q': uint64(8),
				},
				Black: map[Piece]uint64{
					'P': uint64(67272519433846784),
					'R': uint64(9295429630892703744),
					'N': uint64(4755801206503243776),
					'B': uint64(288230376185266176),
					'K': uint64(1152921504606846976),
					'Q': uint64(576460752303423488),
				},
				Empty: uint64(2310628015064680448),
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

			for p, expectedPos := range expected.Pieces.White {
				resultPos := b.Pieces.White[p]
				if expectedPos != resultPos {
					t.Errorf("Error setting White Piece %c position. Expected %b but got %b", p, expectedPos, resultPos)
				}
			}

			for p, expectedPos := range expected.Pieces.Black {
				resultPos := b.Pieces.Black[p]
				if expectedPos != resultPos {
					t.Errorf("Error setting Black Piece %c position. Expected %b but got %b", p, expectedPos, resultPos)
				}
			}

			if b.State.InCheck() != expected.State.inCheck {
				t.Errorf("Error setting Incheck. Expecting %t, got %t", expected.State.inCheck, b.State.InCheck())
			}
		})
	}
}
