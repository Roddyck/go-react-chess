import type { Game, Piece as ChessPiece } from "./chess";
import Square from "./Square";

interface ChessBoardProps {
  game: Game;
}

function ChessBoard({ game }: ChessBoardProps) {
  return (
    <div style={{ display: "inline-block" }}>
      {game.board.map((row, y) => (
        <div key={y} style={{ display: "flex" }}>
          {row.map((piece, x) => {
            const isLight = (x + y) % 2 === 0;
            return <Square key={`${x}-${y}`} piece={piece} isLight={isLight} />;
          })}
        </div>
      ))}
    </div>
  );
}

export { ChessBoard };
