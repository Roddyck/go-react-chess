package game

type Game struct {
	ID      string           `json:"id"`
	Board   [8][8]Piece      `json:"board"`
	Turn    Color            `json:"turn"`
	History []Move           `json:"history"`
	Players map[Color]string `json:"players"`
}

func NewGame() *Game {
	game := &Game{
		Board:   [8][8]Piece{},
		Turn:    White,
		History: []Move{},
		Players: make(map[Color]string),
	}

	game.initBoard()

	return game
}

func (g *Game) initBoard() {
	// black pieces at the top and white pieces at the bottom
	for x := range 8 {
		g.Board[1][x] = NewPiece(Pawn, Black)
		g.Board[6][x] = NewPiece(Pawn, White)
	}

	pieces := []PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}

	for x, piece := range pieces {
		g.Board[0][x] = NewPiece(piece, Black)
		g.Board[7][x] = NewPiece(piece, White)
	}
}
