package board

type castlingRights byte

const (
	KING_SIDE  castlingRights = 'K'
	QUEEN_SIDE castlingRights = 'Q'
	BOTH       castlingRights = '*'
	NONE       castlingRights = '-'
)

type boardState struct {
	currPlayer Player

	InCheck bool

	PlayersCastleRights map[Player]castlingRights

	EnPassantSquare int
}
