package game

import (
	"fmt"

	"github.com/google/uuid"
)

type Game struct {
	ID      uuid.UUID           `json:"id"`
	Board   [8][8]Piece         `json:"board"`
	Turn    Color               `json:"turn"`
	History []Move              `json:"history"`
	Players map[Color]uuid.UUID `json:"players"`
}

func NewGame() *Game {
	game := &Game{
		Board:   [8][8]Piece{},
		Turn:    White,
		History: []Move{},
		Players: make(map[Color]uuid.UUID),
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

func (g *Game) HandleMove(move Move) error {
	piece := g.Board[move.From.Y][move.From.X]
	if piece.GetColor() != g.Turn {
		return fmt.Errorf("invalid move: trying to move a %s piece on %s turn", piece.GetType(), g.Turn)
	}

	g.Board[move.To.Y][move.To.X] = piece
	g.Board[move.From.Y][move.From.X] = nil
	g.History = append(g.History, move)

	if g.Turn == White {
		g.Turn = Black
	} else {
		g.Turn = White
	}

	return nil
}
