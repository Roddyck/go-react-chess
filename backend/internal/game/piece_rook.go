package game

type RookPiece struct {
	Type PieceType `json:"type"`
	Color Color    `json:"color"`
}

func (r *RookPiece) GetType() PieceType {
	return r.Type
}

func (r *RookPiece) GetColor() Color {
	return r.Color
}
