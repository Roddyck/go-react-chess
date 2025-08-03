package game

import "fmt"

type BishopPiece struct {
	Type  PieceType `json:"type"`
	Color Color     `json:"color"`
}

func (b *BishopPiece) GetType() PieceType {
	return b.Type
}

func (b *BishopPiece) GetColor() Color {
	return b.Color
}

func (b *BishopPiece) CheckLegalMove(g *Game, move *Move) error {
	from := move.From
	to := move.To
	board := g.Board
	fromPiece := board[from.Y][from.X]
	toPiece := board[to.Y][to.X]

	if toPiece != nil && toPiece.GetColor() == fromPiece.GetColor() {
		return fmt.Errorf("invalid move: trying to capture own piece")
	}

	if absInt(to.X-from.X) != absInt(to.Y-from.Y) {
		return fmt.Errorf("invalid move: not moving in a diagonal")
	}

	i := from.X
	j := from.Y
	boundX := to.X
	boundY := to.Y
	if i < boundX {
		boundX--
	} else {
		boundX++
	}

	if j < boundY {
		boundY--
	} else {
		boundY++
	}

	for i != boundX && j != boundY {
		if i < boundX {
			i++
		} else {
			i--
		}

		if j < boundY {
			j++
		} else {
			j--
		}

		if i < 8 && j < 8 {
			if board[j][i] != nil {
				return fmt.Errorf("invalid move: piece in the way")
			}
		} else {
			return fmt.Errorf("invalid move: out of bounds")
		}
	}

	return nil
}
