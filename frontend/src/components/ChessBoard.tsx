import type { Game } from "./chess";
import Square from "./Square";

interface ChessBoardProps {
  game: Game;
}

function ChessBoard({ game }: ChessBoardProps) {
  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-900 p-4">
      <div className="shadow-2xl rounded-lg overflow-hidden">
      {game.board.map((row, y) => (
        <div key={y} style={{ display: "flex" }}>
          {row.map((piece, x) => {
            const isLight = (x + y) % 2 === 0;
            return (
              <Square
                key={`${x}-${y}`}
                position={{ x, y }}
                piece={piece}
                isLight={isLight}
              />
            );
          })}
        </div>
      ))}
      </div>
    </div>
  );
}

export { ChessBoard };
