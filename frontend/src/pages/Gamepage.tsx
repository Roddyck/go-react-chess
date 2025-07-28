import { useEffect, useState } from "react";
import { ChessBoard } from "../components/ChessBoard";
import type { Game } from "../components/chess";

function GamePage() {
  const [game, setGame] = useState<Game | null>(null);

  useEffect(() => {
    const fetchGame = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/games");
        const game = await response.json();
        console.log(game);
        setGame(game);
      } catch (error) {
        console.error(error);
      }
    };

    fetchGame();
  }, []);

  if (!game) return <div>WTF...</div>;

  return (
    <div>
      <ChessBoard game={game} />
    </div>
  );
}

export { GamePage };
