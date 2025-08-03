package game

import "fmt"

type QueenPiece struct {
	Type  PieceType `json:"type"`
	Color Color     `json:"color"`
}

func (q *QueenPiece) GetType() PieceType {
	return q.Type
}

func (q *QueenPiece) GetColor() Color {
	return q.Color
}

func (q *QueenPiece) CheckLegalMove(g *Game, move *Move) error {
	from := move.From
	to := move.To
	board := g.Board
	fromPiece := board[from.Y][from.X]
	toPiece := board[to.Y][to.X]

	if toPiece != nil && toPiece.GetColor() == fromPiece.GetColor() {
		return fmt.Errorf("invalid move: trying to capture own piece")
	}

	if absInt(to.X-from.X)*absInt(to.Y-from.Y) != 0 &&
		absInt(to.X-from.X) != absInt(to.Y-from.Y) {
		return fmt.Errorf("invalid move: invalid queen move")
	}

	i := from.X
	j := from.Y
	boundX := to.X
	boundY := to.Y

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
