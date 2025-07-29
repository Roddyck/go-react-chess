export type Color = "white" | "black";

export type PieceType =
  | "pawn"
  | "rook"
  | "knight"
  | "bishop"
  | "queen"
  | "king";

export type Piece = {
  type: PieceType;
  color: Color;
}

export type Position = {
  x: number;
  y: number;
}

export type Game = {
  ID: string;
  board: (Piece | null)[][];
  turn: Color;
}
