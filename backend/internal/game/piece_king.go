package game

type KingPiece struct {
	Type PieceType `json:"type"`
	Color Color    `json:"color"`
}

func (k *KingPiece) GetType() PieceType {
	return k.Type
}

func (k *KingPiece) GetColor() Color {
	return k.Color
}
