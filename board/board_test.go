package board

import (
	"testing"
)

func TestBoard_FromFen(t *testing.T) {
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
					'K': uint64(8),
					'Q': uint64(16),
				},
				Black: map[Piece]uint64{
					'P': uint64(71776119061217280),
					'R': uint64(9295429630892703744),
					'N': uint64(4755801206503243776),
					'B': uint64(2594073385365405696),
					'K': uint64(576460752303423488),
					'Q': uint64(1152921504606846976),
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
				White: map[Piece]uint64{},
				Black: map[Piece]uint64{},
				Empty: uint64(7503243165898056130),
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
				White: map[Piece]uint64{},
				Black: map[Piece]uint64{},
				Empty: uint64(290763615508566016),
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
					t.Errorf("Error setting White Piece position. Expected %b but got %b", expectedPos, resultPos)
				}
			}

			for p, expectedPos := range expected.Pieces.Black {
				resultPos := b.Pieces.Black[p]
				if expectedPos != resultPos {
					t.Errorf("Error setting Black Piece position. Expected %b but got %b", expectedPos, resultPos)
				}
			}

			if b.State.InCheck() != expected.State.inCheck {
				t.Errorf("Error setting Incheck. Expecting %t, got %t", expected.State.inCheck, b.State.InCheck())
			}
		})
	}
}
