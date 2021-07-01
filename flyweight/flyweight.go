package flyweight

var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "che",
		Color: "red",
	},
	2: {
		ID:    2,
		Name:  "pao",
		Color: "red",
	},
}

type ChessPieceUnit struct {
	ID    int
	Name  string
	Color string
}

func NewChessPieceUnit(id int) *ChessPieceUnit {
	return units[id]
}

type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
	for id := range units {
		board.chessPieces[id] = &ChessPiece{
			Unit: NewChessPieceUnit(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

func (c *ChessBoard) Move(id, x, y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y
}
