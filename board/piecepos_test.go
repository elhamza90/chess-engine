package board

import "testing"

func Test_PiecePos_All(t *testing.T) {
	pp := PiecePositions{
		PAWN:   Bitboard(65280),
		KNIGHT: Bitboard(66),
		BISHOP: Bitboard(36),
		ROOK:   Bitboard(129),
		QUEEN:  Bitboard(8),
		KING:   Bitboard(16),
	}
	expected := Bitboard(65535)
	res := pp.All()
	if res != expected {
		t.Errorf("Error combining all pieces positions in one bitboard.\n Expected %b but got %b", expected, res)
	}
}

func Test_PlayerPiecePos_Empty(t *testing.T) {
	ppp := PlayerPiecePositions{
		WHITE: PiecePositions{
			PAWN:   Bitboard(65280),
			ROOK:   Bitboard(129),
			KNIGHT: Bitboard(66),
			BISHOP: Bitboard(36),
			KING:   Bitboard(16),
			QUEEN:  Bitboard(8),
		},
		BLACK: PiecePositions{
			PAWN:   Bitboard(71776119061217280),
			ROOK:   Bitboard(9295429630892703744),
			KNIGHT: Bitboard(4755801206503243776),
			BISHOP: Bitboard(2594073385365405696),
			KING:   Bitboard(1152921504606846976),
			QUEEN:  Bitboard(576460752303423488),
		},
	}
	res := ppp.Empty()
	expected := Bitboard(281474976645120)
	if res != expected {
		t.Errorf("Error generating Empty Cases Bitboard.\n  Expected %b but got %b", expected, res)
	}
}

func Test_PiecePos_kingPseudoLegalMoves(t *testing.T) {
	tests := []struct {
		name     string               // Name of the test
		fen      string               // Fen representation of the position to be tested
		position PlayerPiecePositions // Test Position of the pieces on the board
		player   Player               // Which King is concerned (white / black)
		expected Bitboard             // Expected bitboard representing pseudo-legal moves for the king
	}{
		{
			name: "W-King alone in center",
			fen:  "8/8/8/8/3K4/8/8/8 w - - 0 1",
			position: PlayerPiecePositions{
				WHITE: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(0),
					BISHOP: Bitboard(0),
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(134217728), // 2**27 (D4)
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
			player:   WHITE,
			expected: Bitboard(120596463616), // D5,D3,C4,C3,C5,E4,E3,E5: 2^35 + 2^34 + 2^36 + 2^26 + 2^28 + 2^19 + 2^18 + 2^20
		},
		{
			name: "W-King Blocked by friendly bishop on the upper rank",
			fen:  "8/8/8/3B4/3K4/8/8/8 w - - 0 1",
			position: PlayerPiecePositions{
				WHITE: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(0),
					BISHOP: Bitboard(34359738368), // 2^35 (D5)
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(134217728), // 2^27 (D4)
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
			player:   WHITE,
			expected: Bitboard(86236725248), // D3,C4,C3,C5,E4,E3,E5: 2^34 + 2^36 + 2^26 + 2^28 + 2^19 + 2^18 + 2^20
		},
		{
			name: "W-King Blocked by friendly bishop on the lower rank",
			fen:  "8/8/8/8/3K4/3B4/8/8 w - - 0 1",
			position: PlayerPiecePositions{
				WHITE: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(0),
					BISHOP: Bitboard(524288), // 2^19 (D3)
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(134217728), // 2^27 (D4)
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
			player:   WHITE,
			expected: Bitboard(120595939328), // D5,D3,C4,C3,C5,E4,E3,E5: 2^35 + 2^34 + 2^36 + 2^26 + 2^28 + 2^18 + 2^20
		},
		{
			name: "W-King Blocked by friendly bishop on the left square",
			fen:  "8/8/8/8/2BK4/8/8/8 w - - 0 1",
			position: PlayerPiecePositions{
				WHITE: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(0),
					BISHOP: Bitboard(67108864), // 2^26 (C4)
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(134217728), // 2^27 (D4)
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
			player:   WHITE,
			expected: Bitboard(120529354752), // D5,D3,C3,C5,E4,E3,E5: 2^35 + 2^34 + 2^36 + 2^28 + 2^19 + 2^18 + 2^20
		},
		{
			name: "W-King Blocked by friendly bishop on the right square",
			fen:  "8/8/8/8/3KB3/8/8/8 w - - 0 1",
			position: PlayerPiecePositions{
				WHITE: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(0),
					BISHOP: Bitboard(268435456), // 2^28 (E4)
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(134217728), // 2^27 (D4)
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
			player:   WHITE,
			expected: Bitboard(120328028160), // D5,D3,C4,C3,C5,E3,E5: 2^35 + 2^34 + 2^36 + 2^26 + 2^19 + 2^18 + 2^20
		},
		{
			name: "W-King alone in lower-left corner (A1)",
			fen:  "8/8/8/8/8/8/8/K7 w - - 0 1",
			position: PlayerPiecePositions{
				WHITE: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(0),
					BISHOP: Bitboard(0),
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(1), // 2**0 (A1)
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
			player:   WHITE,
			expected: Bitboard(770), // A2,B1,B2: 2^1 + 2^8 + 2^9
		},
		{
			name: "W-King alone in upper-left corner (A8)",
			fen:  "K7/8/8/8/8/8/8/8 w - - 0 1",
			position: PlayerPiecePositions{
				WHITE: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(0),
					BISHOP: Bitboard(0),
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(72057594037927936), // 2^56 (A8)
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
			player:   WHITE,
			expected: Bitboard(144959613005987840), // B7,B8,A7: 2^57 + 2^49 + 2^48
		},
		{
			name: "W-King alone in upper-right corner (H8)",
			fen:  "7K/8/8/8/8/8/8/8 w - - 0 1",
			position: PlayerPiecePositions{
				WHITE: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(0),
					BISHOP: Bitboard(0),
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(9223372036854775808), // 2^63 (H8)
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
			player:   WHITE,
			expected: Bitboard(4665729213955833856), // H7,G7,G8: 2^62 + 2^54 + 2^55
		},
		{
			name: "W-King alone in lower-right corner (H1)",
			fen:  "8/8/8/8/8/8/8/7K w - - 0 1",
			position: PlayerPiecePositions{
				WHITE: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(0),
					BISHOP: Bitboard(0),
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(128), // 2^7 (H1)
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
			player:   WHITE,
			expected: Bitboard(49216), // G1, H2, G2: 2^14 + 2^15 + 2^6
		},
		{
			name: "W-King in center with enemy knight in upper right",
			fen:  "8/8/8/4b3/3K4/8/8/8 w - - 0 1",
			position: PlayerPiecePositions{
				WHITE: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(0),
					BISHOP: Bitboard(0),
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(134217728), // 2^27 (D4)
				},
				BLACK: {
					PAWN:   Bitboard(0),
					KNIGHT: Bitboard(68719476736), // 2^36 (E5)
					BISHOP: Bitboard(0),
					ROOK:   Bitboard(0),
					QUEEN:  Bitboard(0),
					KING:   Bitboard(0),
				},
			},
			player:   WHITE,
			expected: Bitboard(120596463616), // D5,D3,C4,C3,C5,E4,E3,E5: 2^35 + 2^34 + 2^36 + 2^26 + 2^28 + 2^19 + 2^18 + 2^20
		},
	}

	for _, test := range tests {
		res := test.position.kingPseudoLegalMoves(test.player)
		if res != test.expected {
			t.Errorf("Error in generating pseudo-legal moves for King (%s).\n  Expected %b but got %b \n(in position %s).", string(test.name), test.expected, res, test.fen)
		}
	}

}
