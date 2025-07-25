import { useState, useEffect } from "react";
import { ChessBoard } from "./ChessBoard";
import type { Game } from "./chess";

function Home() {
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

  if (!game) return <div>Loading...</div>;

  return (
    <div>
      <ChessBoard game={game} />
    </div>
  );
}

export { Home };
