package game

type PieceType string

type Color string

const (
	White Color = "white"
	Black Color = "black"
)

const (
	Pawn   PieceType = "pawn"
	Rook   PieceType = "rook"
	Knight PieceType = "knight"
	Bishop PieceType = "bishop"
	Queen  PieceType = "queen"
	King   PieceType = "king"
)

type Piece struct {
	Type  PieceType `json:"type"`
	Color Color     `json:"color"`
}
