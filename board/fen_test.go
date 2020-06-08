package board

import (
	"testing"
)

func TestFen_FromFen(t *testing.T) {}

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
		"Wrong Half Clock":                          "rnbqkrnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - X 1",
		"Wrong Full Clock":                          "rnbqkrnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 Z",
		"3 Rooks in 8th rank":                       "rnbqkrnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"Wrong player to move":                      "r1bq1rk1/ppppn1pp/5n2/5p2/1bPP4/2NQ3N/PP1BPPBP/R3K2R x KQ - 9 9",
		"Wrong number of empty squares in 5th rank": "r1bq1rk1/pppp2pp/5nn1/5p3/2PP4/2PQ3N/P2BPPBP/R3K1R1 b Q - 2 11",
		"Wrong Castling rights":                     "rnbqkr2/pppp2pp/5n2/5p2/1bPP4/2NQ4/PP2PP1P/R1B1KBNR w KRq - 4 7",
		"Wrong En Passant":                          "rnbq1rk1/pppp2pp/5n2/8/1bPPPp2/2NQ4/PP3PBP/R1B1K1NR b KQ e9 0 8",
	}

	for _, fen := range validFens {
		t.Run("Valid Fens", func(t *testing.T) {
			if !fenIsValid(fen) {
				t.Errorf("Valid FEN was rejected: %s", fen)
			}
		})
	}

	for name, fen := range invalidFens {
		t.Run(name, func(t *testing.T) {
			if fenIsValid(fen) {
				t.Errorf("%s", fen)
			}
		})
	}
}
