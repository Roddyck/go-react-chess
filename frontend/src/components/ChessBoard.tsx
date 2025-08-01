import type { Game, Color } from "./chess";
import Square from "./Square";

interface ChessBoardProps {
  game: Game;
  playerColor: Color;
}

function ChessBoard({ game, playerColor }: ChessBoardProps) {
  const renderBoardForWhite = () => {
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
  };

  const renderBoardForBlack = () => {
    const revesedBoard = game.board.map((row) => row.reverse()).reverse();
    return (
      <div className="flex justify-center items-center min-h-screen bg-gray-900 p-4">
        <div className="shadow-2xl rounded-lg overflow-hidden">
          {revesedBoard.map((row, y) => (
            <div key={y} style={{ display: "flex" }}>
              {row.map((piece, x) => {
                const isLight = ((7 - x) + (7 - y)) % 2 === 0;
                return (
                  <Square
                    key={`${7 - x}-${7 - y}`}
                    position={{ x: 7 - x, y: 7 - y }}
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
  };

  return (
    <div className="flex flex-col items-center justify-center">
      {playerColor === "white" ? renderBoardForWhite() : renderBoardForBlack()}
    </div>
  );
}

export { ChessBoard };
