export type Color = "white" | "black";

export type PieceType =
  | "pawn"
  | "rook"
  | "knight"
  | "bishop"
  | "queen"
  | "king";

export interface Piece {
  type: PieceType;
  color: Color;
}

export interface Game {
  board: (Piece | null)[][];
  turn: Color;
}
