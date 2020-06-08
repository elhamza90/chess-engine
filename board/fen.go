package board

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// FromFen sets the state of the board
// to the state given by the FEN representation
func (b Board) FromFen(fen string) {
	if ok := fenIsValid(fen); !ok {
		return
	}
	// TODO
}

func fenIsValid(fen string) bool {
	p := "^((?i)[pnbrqk0-9]*/){7}((?i)[pnbrqk0-9])* [wb] ((?i)[kq-])* [a-h0-8-]* [0-9]* [0-9]*$"
	r, err := regexp.Compile(p)
	if err != nil {
		log.Print("Regex could not compile")
		return false
	}
	if match := r.MatchString(fen); !match {
		log.Print("Fen did not match Regex pattern structure")
		return false
	}
	parts := strings.Split(fen, " ")
	ranks := strings.Split(parts[0], "/")
	if len(ranks) != 8 {
		log.Print("Number of ranks is not 8")
		return false
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
			log.Printf("Rank %d (%s)contains more than 8 squares! (%d pieces and %d empty)", i+1, rank, nbrPieces, nbrEmpty)
			return false
		}

	}
	return true
}
