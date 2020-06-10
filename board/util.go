package board

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Define Players
type Player byte

const (
	WHITE Player = 'W'
	BLACK Player = 'B'
)

// Define Pieces
type Piece byte

const (
	PAWN   Piece = 'P'
	KNIGHT Piece = 'N'
	BISHOP Piece = 'B'
	ROOK   Piece = 'R'
	QUEEN  Piece = 'Q'
	KING   Piece = 'K'
)

// squareToBitBoard takes a square (eg. b2)
// and returns the index of the square
// in a flattened board (eg. b2 -> 9)
func squareToIndex(sq string) (byte, error) {
	if match, err := regexp.MatchString("^[a-h][1-9]$", sq); (err != nil) || !match {
		return 0, err
	}
	colMap := map[byte]byte{
		'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7,
	}
	col := colMap[sq[0]]
	row, err := strconv.Atoi(string(sq[1]))
	if err != nil {
		return 0, err
	}
	row = row - 1
	//log.Print(sq, " ", row, byte(row), col)
	return byte(row)*8 + byte(col), nil
}

func fenIsValid(fen string) error {
	p := "^((?i)[pnbrqk0-9]*/){7}((?i)[pnbrqk0-9])* [wb] ((?i)[kq-])* [a-h36-]* [0-9]* [0-9]*$"
	r, err := regexp.Compile(p)
	if err != nil {
		return errors.New("Regex could not compile")
	}
	if match := r.MatchString(fen); !match {
		return errors.New("Fen did not match Regex pattern structure")
	}
	parts := strings.Split(fen, " ")
	ranks := strings.Split(parts[0], "/")
	if len(ranks) != 8 {
		return errors.New("Number of ranks is not 8")
	}

	for i, rank := range ranks {
		// Check number of overall pieces
		nbrEmpty := 0
		nbrPieces := 0
		for _, r := range rank {
			if unicode.IsDigit(r) {
				// It's a number of consecutive empty spaces
				n, _ := strconv.Atoi(string(r))
				nbrEmpty += n
			} else {
				// It's a piece
				nbrPieces += 1
			}
		}
		if nbrEmpty+nbrPieces > 8 {
			//log.Printf()
			return errors.New(fmt.Sprintf("Rank %d (%s)contains more than 8 squares! (%d pieces and %d empty)", i+1, rank, nbrPieces, nbrEmpty))
		}

	}
	return nil
}
