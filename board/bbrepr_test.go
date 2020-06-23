package board

import (
	"testing"
)

func TestUtil_squareToIndex(t *testing.T) {
	tests := map[string]byte{
		"a1": 0,
		"a8": 56,
		"b1": 1,
		"e3": 20,
		"d5": 35,
		"h8": 63,
	}

	for sq, ix := range tests {
		t.Run("", func(t *testing.T) {
			res, err := squareToIndex(sq)
			if res != ix || err != nil {
				t.Errorf("Expecting %d but Got %d for %s", ix, res, sq)
			}
		})
	}
}
