package game

type KnightPiece struct {
	Type PieceType `json:"type"`
	Color Color    `json:"color"`
}

func (k *KnightPiece) GetType() PieceType {
	return k.Type
}

func (k *KnightPiece) GetColor() Color {
	return k.Color
}
