import { useState } from "react";
import type { Game, Color, Move, Position } from "./chess";
import Square from "./Square";

interface ChessBoardProps {
  game: Game;
  playerColor: Color;
  onMove: (move: Move) => void;
}

function ChessBoard({ game, playerColor, onMove }: ChessBoardProps) {
  const [selectedPos, setSelectedPos] = useState<Position | null>(null);
  const renderBoardForWhite = () => {
    return (
      <div className="flex justify-center items-center min-h-screen bg-gray-900 p-4">
        <div className="shadow-2xl rounded-lg overflow-hidden">
          {game.board.map((row, y) => (
            <div key={y} style={{ display: "flex" }}>
              {row.map((piece, x) => {
                const isSelected = selectedPos?.x === x && selectedPos?.y === y;
                const isLight = (x + y) % 2 === 0;
                return (
                  <Square
                    key={`${x}-${y}`}
                    position={{ x, y }}
                    piece={piece}
                    isLight={isLight}
                    isSelected={isSelected}
                    onClick={handleSquareClick}
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
                const isSelected =
                  selectedPos?.x === x && selectedPos?.y === 7 - y;
                const isLight = (x + (7 - y)) % 2 === 0;
                return (
                  <Square
                    key={`${7 - x}-${7 - y}`}
                    position={{ x: x, y: 7 - y }}
                    piece={piece}
                    isLight={isLight}
                    isSelected={isSelected}
                    onClick={handleSquareClick}
                  />
                );
              })}
            </div>
          ))}
        </div>
      </div>
    );
  };

  const handleSquareClick = (pos: Position) => {
    if (selectedPos) {
      console.log(selectedPos, pos);
      const move = { from: selectedPos, to: pos };
      onMove(move);
      setSelectedPos(null);
    } else {
      const piece = game.board[pos.y][pos.x];
      if (piece && piece.color === playerColor && piece.color === game.turn) {
        setSelectedPos(pos);
      }
    }
  };

  return (
    <div className="flex flex-col items-center justify-center">
      {playerColor === "white" ? renderBoardForWhite() : renderBoardForBlack()}
    </div>
  );
}

export { ChessBoard };
