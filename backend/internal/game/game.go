package game

import (
	"fmt"

	"github.com/google/uuid"
)

type Game struct {
	ID      uuid.UUID           `json:"id"`
	Board   [8][8]Piece         `json:"board"`
	Turn    Color               `json:"turn"`
	History []*Move             `json:"history"`
	Players map[Color]uuid.UUID `json:"players"`
}

func NewGame() *Game {
	game := &Game{
		Board:   [8][8]Piece{},
		Turn:    White,
		History: []*Move{},
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

func (g *Game) updateBoard(move *Move) {
	piece := g.Board[move.From.Y][move.From.X]
	g.Board[move.From.Y][move.From.X] = nil

	switch p := piece.(type) {
	case *PawnPiece:
		g.Board[move.To.Y][move.To.X] = piece
		if !p.HasMoved {
			p.HasMoved = true
		}
		if len(g.History) > 0 {
			if err := p.canEnpassant(g, move, g.History[len(g.History)-1]); err == nil {
				g.Board[move.From.Y][move.To.X] = nil
			}
		}
	case *KingPiece:
		if p.GetColor() == White {
			if err := p.canCastle(g, move); err == nil {
				switch move.To.X {
				case 7:
					g.Board[7][6] = piece
					g.Board[7][5] = NewPiece(Rook, White)
				case 0:
					g.Board[7][2] = piece
					g.Board[7][3] = NewPiece(Rook, White)
				}
				g.Board[move.To.Y][move.To.X] = nil
			} else {
				g.Board[move.To.Y][move.To.X] = piece
			}
		} else {
			if err := p.canCastle(g, move); err == nil {
				switch move.To.X {
				case 7:
					g.Board[0][6] = piece
					g.Board[0][5] = NewPiece(Rook, Black)
				case 0:
					g.Board[0][2] = piece
					g.Board[0][3] = NewPiece(Rook, Black)
				}
				g.Board[move.To.Y][move.To.X] = nil
			} else {
				g.Board[move.To.Y][move.To.X] = piece
			}
		}
		if !p.HasMoved {
			p.HasMoved = true
		}
	case *RookPiece:
		if !p.HasMoved {
			p.HasMoved = true
		}
		g.Board[move.To.Y][move.To.X] = piece
	default:
		g.Board[move.To.Y][move.To.X] = piece
	}
}

func (g *Game) HandleMove(move *Move) error {
	piece := g.Board[move.From.Y][move.From.X]
	if piece.GetColor() != g.Turn {
		return fmt.Errorf("invalid move: trying to move a %s piece on %s turn", piece.GetType(), g.Turn)
	}

	if err := piece.CheckLegalMove(g, move); err != nil {
		return err
	}

	g.updateBoard(move)

	if g.Turn == White {
		g.Turn = Black
	} else {
		g.Turn = White
	}

	g.History = append(g.History, move)

	return nil
}
