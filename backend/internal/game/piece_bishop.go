package game

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
