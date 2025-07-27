package game

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
