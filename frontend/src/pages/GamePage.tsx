import { useEffect, useState } from "react";
import { ChessBoard } from "../components/ChessBoard";
import type { Game } from "../components/chess";
import { useParams } from "react-router";
import { useWebSocket } from "../api/websocket";
import { useAuth } from "../context/AuthContext";

function GamePage() {
  const { sessionID } = useParams<{ sessionID: string }>();
  const [gameID, setGameID] = useState<string | null>(null);
  const [game, setGame] = useState<Game | null>(null);
  const { token } = useAuth();
  const { user } = useAuth();

  const { sendMessage } = useWebSocket(
    `ws://localhost:8080/ws/sessions/${sessionID}?userID=${user?.id}&username=${user?.name}`,
    (msg) => {
      if (!msg) {
        console.error("No message");
      }
      console.log(msg);
      setGameID(msg.data.game_id);
      console.log(msg.data.game_id);
    }
  );

  const sendHello = () => {
    console.log("Sending hello");

    if (sessionID) {
      sendMessage({ action: "hello", session_id: sessionID, data: {msg: "Hello, World!"} });
    } else {
      console.error("No session ID");
    }
  };

  useEffect(() => {
    const fetchGame = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/games", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
          body: JSON.stringify({ game_id: gameID }),
        });
        const game = await response.json();
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
      <div className="flex justify-center items-center">
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg"
          onClick={sendHello}
        >
          Send Hello
        </button>
      </div>
    </div>
  );
}

export { GamePage };
