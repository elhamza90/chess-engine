package board

type castleRights byte

const (
	CASTLE_KING  castleRights = 2
	CASTLE_QUEEN castleRights = 1
	CASTLE_NONE  castleRights = 0
)

type boardState struct {
	currPlayer Player

	inCheck bool

	playersCastleRights map[Player]castleRights

	epSquare uint64 // en passant square in a 64bit bitboard
}

func (bs boardState) CurrPlayer() Player {
	return bs.currPlayer
}

func (bs boardState) InCheck() bool {
	return bs.inCheck
}

func (bs boardState) PlayersCastleRights() map[Player]castleRights {
	return bs.playersCastleRights
}

func (bs boardState) EpSquare() uint64 {
	return bs.epSquare
}
