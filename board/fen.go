package board

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// FromFen sets the state of the board and piece positions
// to the state and positions given by the FEN representation
func (b *Board) FromFen(fen string) {
	if err := fenIsValid(fen); err != nil {
		return
	}
	parts := strings.Split(fen, " ")

	// Set Current Player
	if len(parts[1]) == 1 {
		upper := unicode.ToUpper(rune(parts[1][0]))
		//log.Print(Player(upper))
		b.State.currPlayer = Player(upper)
	}

	// Set Castling rights
	b.State.playersCastleRights = map[Player]castleRights{}
	b.State.playersCastleRights[WHITE], b.State.playersCastleRights[BLACK] = fenCastleRights(parts[2])

	// Set En Passant Square
	if parts[3] != "-" {
		epIndex, err := squareToIndex(parts[3])
		if err != nil {
			return
		}
		b.State.epSquare = uint64(1) << epIndex
	}

	// Set Pieces locations in Bitboards
	b.Pieces.White, b.Pieces.Black = fenToBitboardPieces(parts[0])

	// Calculate Empty Squares from pieces locations
	var occupied uint64 = 0
	for _, pos := range b.Pieces.Black {
		occupied += pos
	}
	for _, pos := range b.Pieces.White {
		occupied += pos
	}
	b.Pieces.Empty = ^occupied

}

// fenIsValid returns an error if the given Fen string is invalid
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

// fenCastleRights returns castle rights for White & Black
// respectively given the fen castle rights representation (KQkq)
func fenCastleRights(s string) (castleRights, castleRights) {
	//log.Printf("Castling Rights: %s", s)
	var w_castling byte = 0
	var b_castling byte = 0
	for _, b := range s {
		if b == 'K' {
			w_castling += byte(CASTLE_KING)
		} else if b == 'Q' {
			w_castling += byte(CASTLE_QUEEN)
		} else if b == 'k' {
			b_castling += byte(CASTLE_KING)
		} else if b == 'q' {
			b_castling += byte(CASTLE_QUEEN)
		}
	}
	return castleRights(w_castling), castleRights(b_castling)
}

// fenToBitboardPieces returns bitboards for White and Black
// respectively given the Fen Pieces representation
func fenToBitboardPieces(fen string) (map[Piece]uint64, map[Piece]uint64) {
	blackMap := map[Piece]uint64{}
	whiteMap := map[Piece]uint64{}
	//log.Printf("Pieces: %s", parts[0])
	var (
		ix    int    = 0
		inc   int    = 0
		pos   uint64 = 0
		piece Piece
	)
	rows := strings.Split(fen, "/")
	for r := len(rows) - 1; r >= 0; r-- {
		row := rows[r]
		//log.Print(row)
		for c := 0; c < len(row); c++ {
			char := rune(row[c])
			if unicode.IsDigit(char) {
				// Get number of empty squares
				inc, _ = strconv.Atoi(string(char))
				//log.Printf("There are %d zeros in the %dth index", inc, ix)
			} else {
				inc = 1
				// Check type and color of piece
				piece = Piece(unicode.ToUpper(char))
				pos = uint64(1) << ix
				if unicode.IsUpper(char) {
					whiteMap[piece] |= pos
				} else {
					blackMap[piece] |= pos
				}

			}
			ix += inc
		}
	}
	return whiteMap, blackMap
}
