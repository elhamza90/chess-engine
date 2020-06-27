package board

/************************************************************/

type boardState struct {
	currPlayer Player

	inCheck bool

	playersCastleRights map[Player]castleRights

	epSquare Bitboard // en passant square in a 64bit bitboard
}

func (bs boardState) CurrPlayer() Player {
	return bs.currPlayer
}

// oppPlayer return the opponent of the current Player
func (bs boardState) oppPlayer() Player {
	if bs.currPlayer == WHITE {
		return BLACK
	} else {
		return WHITE
	}
}

func (bs boardState) InCheck() bool {
	return bs.inCheck
}

func (bs boardState) PlayersCastleRights() map[Player]castleRights {
	return bs.playersCastleRights
}

func (bs boardState) EpSquare() Bitboard {
	return bs.epSquare
}

/******************** Definitions ****************************/

// Define castle rights codes
type castleRights byte

const (
	CASTLE_KING  castleRights = 2
	CASTLE_QUEEN castleRights = 1
	CASTLE_NONE  castleRights = 0
)

// Define Players
type Player byte

const (
	WHITE Player = 'W'
	BLACK Player = 'B'
)
