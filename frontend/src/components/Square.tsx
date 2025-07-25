import type { Piece } from "./chess";

interface SquareProps {
  piece: Piece | null;
  isLight: boolean;
}

function Square({ piece, isLight }: SquareProps) {
  const getPieceSymbol = () => {
    if (!piece) return null;

    const symbols = {
      white: {
        king: "♔",
        queen: "♕",
        rook: "♖",
        bishop: "♗",
        knight: "♘",
        pawn: "♙",
      },
      black: {
        king: "♚",
        queen: "♛",
        rook: "♜",
        bishop: "♝",
        knight: "♞",
        pawn: "♟",
      },
    };

    return symbols[piece.color][piece.type];
  };

  const getSquareColor = () => {
    return isLight ? "bg-amber-100" : "bg-amber-800";
  };

  return (
    <div
      className={`
        w-14 h-14 md:w-16 md:h-16
        flex justify-center items-center
        relative cursor-pointer
        ${getSquareColor()}
        hover:opacity-80
        select-none
    `}
    >
      {piece && (
        <div
          className={`
            text-3xl md:text-4xl
            ${piece.color === "white" ? "text-white" : "text-black"}
            ${piece.color === "white" ? "drop-shadow-md" : "drop-shadow-[0_1px_1px_rgba(255,255,255,0.5)]"}
          `}
        >
          {getPieceSymbol()}
        </div>
      )}
    </div>
  );
}

export default Square;
