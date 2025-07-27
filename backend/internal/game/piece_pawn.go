package game

type PawnPiece struct {
	Type  PieceType `json:"type"`
	Color Color     `json:"color"`
}

func (p *PawnPiece) GetType() PieceType {
	return p.Type
}

func (p *PawnPiece) GetColor() Color {
	return p.Color
}
