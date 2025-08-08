import type { Color } from "./chess";

type GameEndModalProps = {
  result: {
    type: string;
    winner?: "white" | "black";
  };
  playerColor: Color;
  onClose: () => void;
};

function GameEndModal({ result, playerColor, onClose }: GameEndModalProps) {
  const getModalContent = () => {
    switch (result.type) {
      case "checkmate":
        const isWinner = result.winner === playerColor;
        return {
          title: isWinner ? "You won!" : "You lost",
          message: isWinner
            ? "Checkmate! Congratulations on your victory"
            : "Checkmate! Better luck next time",
          bgColor: isWinner ? "bg-green-100" : "bg-red-100",
          borderColor: isWinner ? "bg-green-400" : "bg-red-400",
          textColor: isWinner ? "text-green-800" : "text-red-800",
        };
      default:
        return {
          title: "Game Over",
          message: "The game has ended.",
          bgColor: "bg-gray-100",
          borderColor: "border-gray-400",
          textColor: "text-gray-800",
        };
    }
  };

  const { title, message, bgColor, borderColor, textColor } = getModalContent();

  return (
    <div
      className={`fixed inset-0 z-50 flex items-center justify-center
      bg-black bg-opacity-50`}
    >
      <div
        className={`w-full max-w-md mx-4 p-6 rounded-lg border-2
        ${borderColor} ${bgColor} shadow-xl`}
      >
        <div className="text-center">
          <h2 className={`text-2xl font-bold mb-2 ${textColor}`}>{title}</h2>
          <p className={`mb-6 ${textColor}`}>{message}</p>

          <div className="flex justify-center space-x-4">
            <button
              onClick={onClose}
              className={`px-4 py-2 bg-gray-200 hover:bg-gray-300 rounded-md
              transition-colors`}
            >
              Close
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export { GameEndModal };
