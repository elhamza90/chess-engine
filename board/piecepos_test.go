package board

import "testing"

func TestPiecePos_kingPseudoLegalMoves(t *testing.T) {
	tests := []struct {
		fen      string   //Fen representation of the position to be tested
		position piecePos // Test Position of the pieces on the board
		player   Player   // Which King is concerned (white / black)
		expected Bitboard //expected bitboard representing pseudo-legal moves for the king
	}{
		{
			fen: "",
			position: piecePos{
				Positions: map[Player]map[Piece]Bitboard{
					WHITE: {
						PAWN:   Bitboard(0),
						KNIGHT: Bitboard(0),
						BISHOP: Bitboard(0),
						ROOK:   Bitboard(0),
						QUEEN:  Bitboard(0),
						KING:   Bitboard(0),
					},
					BLACK: {
						PAWN:   Bitboard(0),
						KNIGHT: Bitboard(0),
						BISHOP: Bitboard(0),
						ROOK:   Bitboard(0),
						QUEEN:  Bitboard(0),
						KING:   Bitboard(0),
					},
				},
				Empty: Bitboard(0),
			},
			player:   WHITE,
			expected: Bitboard(1),
		},
	}

	for _, test := range tests {
		res := test.position.kingPseudoLegalMoves(test.player)
		if res != test.expected {
			t.Errorf("Error in generating pseudo-legal moves for King (%s). Expected %b but got %b \n(in position %s).", string(test.player), test.expected, res, test.fen)
		}
	}

}
