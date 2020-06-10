package board

import "testing"

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
		},
		"r2q1bnr/pp1bkppp/2n5/3pp3/2Pp3P/1P3N2/PB1NPPP1/2RQKB1R b K c3 0 8": {
			State: boardState{
				currPlayer: BLACK,
				inCheck:    false,
				playersCastleRights: map[Player]castleRights{
					WHITE: CASTLE_KING,
					BLACK: CASTLE_NONE,
				},
				epSquare: 262144,
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

			if b.State.InCheck() != expected.State.inCheck {
				t.Errorf("Error setting Incheck. Expecting %t, got %t", expected.State.inCheck, b.State.InCheck())
			}
		})
	}
}
