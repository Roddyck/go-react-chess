package game

import "fmt"

type KnightPiece struct {
	Type  PieceType `json:"type"`
	Color Color     `json:"color"`
}

func (k *KnightPiece) GetType() PieceType {
	return k.Type
}

func (k *KnightPiece) GetColor() Color {
	return k.Color
}

func (k *KnightPiece) CheckLegalMove(g *Game, move *Move) error {
	from := move.From
	to := move.To
	board := g.Board
	fromPiece := board[from.Y][from.X]
	toPiece := board[to.Y][to.X]

	if toPiece != nil && toPiece.GetColor() == fromPiece.GetColor() {
		return fmt.Errorf("invalid move: trying to capture own piece")
	}

	if absInt(to.X-from.X)*absInt(to.Y-from.Y) != 2 {
		return fmt.Errorf("invalid move: invalid knight move")
	}

	return nil
}
