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

type Piece interface {
	GetType() PieceType
	GetColor() Color
	CheckLegalMove(g *Game, move *Move) error
}

func NewPiece(pieceType PieceType, color Color) Piece {
	switch pieceType {
	case Pawn:
		return &PawnPiece{Type: pieceType, Color: color}
	case Rook:
		return &RookPiece{Type: pieceType, Color: color}
	case Knight:
		return &KnightPiece{Type: pieceType, Color: color}
	case Bishop:
		return &BishopPiece{Type: pieceType, Color: color}
	case Queen:
		return &QueenPiece{Type: pieceType, Color: color}
	case King:
		return &KingPiece{Type: pieceType, Color: color}
	default:
		return nil
	}
}

// ok golang, you actually don't have it in the standard library
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
