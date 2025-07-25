package game

type Game struct {
	ID      string           `json:"id"`
	Board   [8][8]*Piece     `json:"board"`
	Turn    Color            `json:"turn"`
	History []Move           `json:"history"`
	Players map[Color]string `json:"players"`
}

func NewGame() *Game {
	game := &Game{
		Board:   [8][8]*Piece{},
		Turn:    White,
		History: []Move{},
		Players: make(map[Color]string),
	}

	game.initBoard()

	return game
}

func (g *Game) initBoard() {
	for x := range 8 {
		g.Board[1][x] = &Piece{Type: Pawn, Color: White}
		g.Board[6][x] = &Piece{Type: Pawn, Color: Black}
	}

	pieces := []PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}

	for x, piece := range pieces {
		g.Board[0][x] = &Piece{Type: piece, Color: White}
		g.Board[7][x] = &Piece{Type: piece, Color: Black}
	}
}
