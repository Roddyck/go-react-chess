import type { Piece, Position } from "./chess";
import { PieceSVG } from "./pieces";

interface SquareProps {
  piece: Piece | null;
  position: Position;
  isLight: boolean;
  isSelected: boolean;
  onClick: (pos: Position) => void;
}

function Square({
  piece,
  position,
  isLight,
  isSelected,
  onClick,
}: SquareProps) {
  const getFile = (x: number) => String.fromCharCode(97 + x);
  const getRank = (y: number) => 8 - y;

  const getSquareColor = (isLight: boolean) => {
    if (isSelected) return "bg-green-200";
    return isLight ? "bg-amber-100" : "bg-amber-800";
  };

  const handleClick = () => {
    onClick(position);
  };

  return (
    <div
      className={`
        w-14 h-14 md:w-16 md:h-16
        flex justify-center items-center
        relative cursor-pointer
        ${getSquareColor(isLight)}
        hover:opacity-80
        select-none
    `}
      onClick={handleClick}
    >
      {(position.y === 7 || position.x === 0) && (
        <div
          className={`absolute top-1 left-1 text-xs ${isLight ? "text-amber-800" : "text-amber-100"}`}
        >
          {position.y === 7 && getFile(position.x)}
          {position.x === 0 && getRank(position.y)}
        </div>
      )}

      {piece && (
        <div className="w-10 h-10 md:w-12 md:h-12 flex items-center justify-center">
          <PieceSVG type={piece.type} color={piece.color} />
        </div>
      )}
    </div>
  );
}

export default Square;
