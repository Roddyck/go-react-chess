package game

import "fmt"

type KingPiece struct {
	Type     PieceType `json:"type"`
	Color    Color     `json:"color"`
	HasMoved bool      `json:"has_moved"`
	InCheck  bool      `json:"in_check"`
}

func (k *KingPiece) GetType() PieceType {
	return k.Type
}

func (k *KingPiece) GetColor() Color {
	return k.Color
}

func (k *KingPiece) CheckLegalMove(g *Game, move *Move) error {
	from := move.From
	to := move.To
	board := g.Board
	fromPiece := board[from.Y][from.X]
	toPiece := board[to.Y][to.X]

	if toPiece != nil &&
		fromPiece.GetColor() == toPiece.GetColor() &&
		!isCastlingMove(move) {
		return fmt.Errorf("invalid move: invalid king move")
	}

	if absInt(from.X-to.X) <= 1 && absInt(from.Y-to.Y) <= 1 {
		if _, ok := toPiece.(*KingPiece); ok {
			return nil
		}

		tmpTaken := toPiece
		toPiece = fromPiece
		fromPiece = nil

		defer func() {
			fromPiece = toPiece
			toPiece = tmpTaken
		}()

		for i := range 8 {
			for j := range 8 {
				piece := board[i][j]

				if piece == nil {
					continue
				}

				if piece.GetColor() == toPiece.GetColor() {
					continue
				}

				if err := piece.CheckLegalMove(g, &Move{
					From: &Position{X: j, Y: i},
					To:   to,
				}); err == nil {
					return fmt.Errorf("invalid move: king will be under attack after move by %s %s at x:%d y:%d, to x:%d y:%d",
						piece.GetColor(), piece.GetType(), j, i, to.X, to.Y)
				}
			}
		}
		return nil
	}

	return k.canCastle(g, move)
}

func isCastlingMove(move *Move) bool {
	from := move.From
	to := move.To

	return from.X == 4 && to.Y == from.Y && (to.X == 0 || to.X == 7) && (from.Y == 0 || from.Y == 7)
}

func (k *KingPiece) canCastle(g *Game, move *Move) error {
	from := move.From
	to := move.To
	board := g.Board
	fromPiece := board[to.Y][to.X]
	toPiece := board[to.Y][to.X]

	if k.HasMoved {
		return fmt.Errorf("invalid move: can't castle after moving your king")
	}

	if r, ok := toPiece.(*RookPiece); !ok || r.HasMoved {
		return fmt.Errorf("invalid move: can't castle without a rook or after moving a rook")
	}

	if to.Y != from.Y {
		return fmt.Errorf("invalid move: can't castle to a different row")
	}

	if k.InCheck {
		return fmt.Errorf("invalid move: can't castle while in check")
	}

	if !isCastlingMove(move) {
		return fmt.Errorf("invalid move: invalid castling move")
	}

	inBetweens := make([]*Position, 0, 6)
	inBetweens = append(inBetweens, from)

	i := from.X
	j := from.Y
	boundX := to.X

	if i < boundX {
		boundX--
	} else if i > boundX {
		boundX++
	}

	for i != boundX {
		if i < boundX {
			i++
		} else {
			i--
		}

		if i > -1 && i < 8 {
			if board[j][i] != nil {
				return fmt.Errorf("invalid move: piece in the way")
			} else {
				inBetweens = append(inBetweens, &Position{i, j})
			}
		} else {
			return fmt.Errorf("invalid move: out of bounds")
		}
	}

	for _, pos := range inBetweens {
		for y := range 8 {
			for x := range 8 {
				square := &Position{X: x, Y: y}

				if board[y][x] == nil {
					continue
				}

				if board[y][x].GetColor() == fromPiece.GetColor() {
					continue
				}

				if err := board[y][x].CheckLegalMove(g, &Move{
					From: square,
					To:   pos,
				}); err == nil {
					return fmt.Errorf("invalid move: can't castle through check")
				}
			}
		}
	}

	return nil
}
