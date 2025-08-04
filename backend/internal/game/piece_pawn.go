package game

import "fmt"

type PawnPiece struct {
	Type     PieceType `json:"type"`
	Color    Color     `json:"color"`
	HasMoved bool      `json:"has_moved"`
}

func (p *PawnPiece) GetType() PieceType {
	return p.Type
}

func (p *PawnPiece) GetColor() Color {
	return p.Color
}

func (p *PawnPiece) CheckLegalMove(g *Game, move *Move) error {
	from := move.From
	to := move.To
	board := g.Board
	fromPiece := board[from.Y][from.X]
	toPiece := board[to.Y][to.X]

	if toPiece != nil && toPiece.GetColor() == fromPiece.GetColor() {
		return fmt.Errorf("invalid move: trying to capture own piece")
	}

	if (fromPiece.GetColor() == White && to.Y > from.Y) ||
		(fromPiece.GetColor() == Black && to.Y < from.Y) {
		return fmt.Errorf("invalid move: pawns can't move backwards")
	}

	if absInt(to.X-from.X) > 1 ||
		absInt(to.Y-from.Y) > 2 ||
		absInt(to.Y-from.Y) == 0 {
		return fmt.Errorf("invalid move: not a valid number of squares or direction")
	}

	if to.X == from.X {
		if toPiece != nil {
			return fmt.Errorf("invalid move: invalid capture")
		}

		// moving two squares
		if dir := to.Y - from.Y; dir == 2 || dir == -2 {
			dir /= 2

			if from.Y+dir < 8 {
				if board[from.Y+dir][from.X] != nil {
					return fmt.Errorf("invalid move: piece in the way")
				}
			} else {
				return fmt.Errorf("invalid move: out of bounds")
			}
			if p.HasMoved {
				return fmt.Errorf("invalid move: can't move two squares past the first move")
			}
			return nil
		}

		return nil
	} else {
		if dir := to.Y - from.Y; dir == 2 || dir == -2 {
			return fmt.Errorf("invalid move: diagonal move by two ranks")
		}
		if toPiece != nil {
			return fmt.Errorf("invalid move: invalid enpassant")
		}
		return nil
	}
}

func (p *PawnPiece) canEnpassant(g *Game, move *Move, lastMove *Move) error {
	from := move.From
	to := move.To

	if lastMove == nil {
		return fmt.Errorf("invalid move: can't enpassant without a previous move")
	}

	_, ok := g.Board[lastMove.To.Y][lastMove.To.X].(*PawnPiece)
	lastMoveDist := absInt(lastMove.To.Y - lastMove.From.Y)
	if ok && lastMoveDist == 2 && lastMove.To.Y == from.Y && lastMove.To.X == to.X {
		return nil
	}

	return fmt.Errorf("invalid move: can't enpassant")
}
