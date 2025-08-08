package game

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type GameStatus string

const (
	active         GameStatus = "active"
	blackCheckmate GameStatus = "black_checkmate"
	whiteCheckmate GameStatus = "white_checkmate"
)

type Game struct {
	ID             uuid.UUID           `json:"id"`
	Board          [8][8]Piece         `json:"board"`
	Turn           Color               `json:"turn"`
	History        []*Move             `json:"history"`
	Players        map[Color]uuid.UUID `json:"players"`
	KingsPositions [2]*Position        `json:"king_positions"`
	Status         GameStatus          `json:"status"`
}

func NewGame() *Game {
	game := &Game{
		Board:          [8][8]Piece{},
		Turn:           White,
		History:        []*Move{},
		Players:        make(map[Color]uuid.UUID),
		KingsPositions: [2]*Position{},
		Status:         active,
	}

	game.initBoard()

	game.KingsPositions[0] = &Position{X: 4, Y: 7}
	game.KingsPositions[1] = &Position{X: 4, Y: 0}

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

func (g *Game) updateKingPositions() {
	board := g.Board
	wkPos := g.KingsPositions[0]
	bkPos := g.KingsPositions[1]
	wk, isWk := board[wkPos.Y][wkPos.Y].(*KingPiece)
	bk, isBk := board[bkPos.Y][bkPos.Y].(*KingPiece)
	if isWk && isBk && wk.GetColor() == White && bk.GetColor() == Black {
		return
	}

	for y := range 8 {
		for x := range 8 {
			if k, ok := board[y][x].(*KingPiece); ok {
				if k.GetColor() == White {
					g.KingsPositions[0] = &Position{X: x, Y: y}
				} else {
					g.KingsPositions[1] = &Position{X: x, Y: y}
				}
			}
		}
	}
}

func (g *Game) kingInCheck() bool {
	g.updateKingPositions()
	var kingPos *Position

	if g.Turn == White {
		kingPos = g.KingsPositions[0]
	} else {
		kingPos = g.KingsPositions[1]
	}

	for y := range 8 {
		for x := range 8 {
			square := &Position{X: x, Y: y}

			piece := g.Board[y][x]

			if piece == nil {
				continue
			}
			if piece.GetColor() == g.Turn {
				continue
			}
			if err := piece.CheckLegalMove(g, &Move{
				From: square,
				To:   kingPos,
			}); err == nil {
				if k, ok := g.Board[kingPos.Y][kingPos.X].(*KingPiece); ok {
					k.InCheck = true
					return true
				}

			}
		}
	}

	return false
}

// Some refactoring would be nice
func (g *Game) kingInCheckmate() bool {
	threatPiecePositions := []*Position{}
	var kingPos *Position

	if g.Turn == White {
		kingPos = g.KingsPositions[0]
	} else {
		kingPos = g.KingsPositions[1]
	}

	for y := range 8 {
		for x := range 8 {
			square := &Position{X: x, Y: y}
			piece := g.Board[square.Y][square.X]

			if piece == nil {
				continue
			}
			if piece.GetColor() == g.Turn {
				continue
			}

			if err := piece.CheckLegalMove(g, &Move{
				From: square,
				To:   kingPos,
			}); err == nil {
				threatPiecePositions = append(threatPiecePositions, square)
			}
		}
	}

	king, ok := g.Board[kingPos.Y][kingPos.X].(*KingPiece)
	if !ok {
		log.Println("kingInCheckmate(): king is not in given position")
		return false
	}

	for i := kingPos.Y - 1; i <= kingPos.Y+1; i++ {
		for j := kingPos.X - 1; j <= kingPos.X+1; j++ {
			if i < 0 || i > 7 || j < 0 || j > 7 {
				continue
			}

			if err := king.CheckLegalMove(g, &Move{
				From: kingPos,
				To:   &Position{X: j, Y: i},
			}); err == nil {
				return true
			}
		}
	}

	// if king can't run and there are multiple pieces attacking
	// then this is checkmate
	if len(threatPiecePositions) > 1 {
		return true
	}

	threatPos := threatPiecePositions[0]

	// we could either capture attacking piece
	for y := range 8 {
		for x := range 8 {
			square := &Position{X: x, Y: y}
			piece := g.Board[y][x]

			if piece == nil {
				continue
			}
			if piece.GetColor() != g.Turn {
				continue
			}
			if err := piece.CheckLegalMove(g, &Move{
				From: square,
				To:   threatPos,
			}); err == nil {
				return false
			}
		}
	}

	// or try to block
	// if attacking piece knight or pawn, we can't block check
	if _, ok := g.Board[threatPos.Y][threatPos.X].(*KnightPiece); ok {
		return true
	}
	if _, ok := g.Board[threatPos.Y][threatPos.X].(*PawnPiece); ok {
		return true
	}

	inBetweens := make([]*Position, 0, 6)
	i := threatPos.X
	j := threatPos.Y
	boundX := kingPos.X
	boundY := kingPos.Y
	if i < boundX {
		boundX--
	} else if i > boundX {
		boundX++
	}
	if j < boundY {
		boundY--
	} else if j > boundY {
		boundY++
	}

	for i != boundX || j != boundY {
		if i < boundX {
			i++
		} else if i > boundX {
			i--
		}
		if j < boundY {
			j++
		} else if j > boundY {
			j--
		}

		inBetweens = append(inBetweens, &Position{X: i, Y: j})
	}

	for _, pos := range inBetweens {
		for y := range 8 {
			for x := range 8 {
				square := &Position{X: x, Y: y}
				piece := g.Board[y][x]

				if piece == nil {
					continue
				}
				if piece.GetColor() != king.GetColor() {
					continue
				}
				if err := piece.CheckLegalMove(g, &Move{
					From: square,
					To:   pos,
				}); err == nil {
					return false
				}

			}
		}
	}

	return true
}

func (g *Game) checkAndNextTurn() {
	if g.Turn == White {
		g.Turn = Black
	} else {
		g.Turn = White
	}

	if g.kingInCheck() && g.kingInCheckmate() {
		if g.Turn == White {
			g.Status = blackCheckmate
		} else {
			g.Status = whiteCheckmate
		}
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
					g.KingsPositions[0] = &Position{X: 6, Y: 7}
				case 0:
					g.Board[7][2] = piece
					g.Board[7][3] = NewPiece(Rook, White)
					g.KingsPositions[0] = &Position{X: 2, Y: 7}
				}
				g.Board[move.To.Y][move.To.X] = nil
			} else {
				g.Board[move.To.Y][move.To.X] = piece
				g.KingsPositions[0] = move.To
			}
		} else {
			if err := p.canCastle(g, move); err == nil {
				switch move.To.X {
				case 7:
					g.Board[0][6] = piece
					g.Board[0][5] = NewPiece(Rook, Black)
					g.KingsPositions[1] = &Position{X: 6, Y: 0}
				case 0:
					g.Board[0][2] = piece
					g.Board[0][3] = NewPiece(Rook, Black)
					g.KingsPositions[1] = &Position{X: 2, Y: 0}
				}
				g.Board[move.To.Y][move.To.X] = nil
			} else {
				g.Board[move.To.Y][move.To.X] = piece
				g.KingsPositions[1] = move.To
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
	g.checkAndNextTurn()

	g.History = append(g.History, move)

	return nil
}
