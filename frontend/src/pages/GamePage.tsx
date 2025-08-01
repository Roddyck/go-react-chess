import { useEffect, useState } from "react";
import { ChessBoard } from "../components/ChessBoard";
import type { Game, Move } from "../components/chess";
import { useParams } from "react-router";
import { useWebSocket } from "../api/websocket";
import { useAuth } from "../context/AuthContext";

type Session = {
  session_id: string;
  game_id: string;
};

function GamePage() {
  const { sessionID } = useParams<{ sessionID: string }>();
  const [gameID, setGameID] = useState<string | null>(null);
  const [game, setGame] = useState<Game | null>(null);
  const { user } = useAuth();

  const { sendMessage } = useWebSocket(
    `ws://localhost:8080/ws/sessions/${sessionID}?userID=${user?.id}&username=${user?.name}`,
    (msg) => {
      if (!msg) {
        console.error("No message");
      }
      console.log(msg);
      setGame(msg.data.game);
    }
  );

  useEffect(() => {
    const getSessionInfo = async () => {
      try {
        const response = await fetch("http://localhost:8080/ws/sessions");
        if (response.status !== 200) {
          console.error("Error getting session info");
          return;
        }
        const sessions: Session[] = await response.json();
        const session = sessions.find((s) => s.session_id === sessionID);
        if (session) {
          setGameID(session.game_id);
          console.log(gameID);
        } else {
          console.error("Session not found");
        }
      } catch (error) {
        console.error(error);
      }
    };

    getSessionInfo();
  }, [sessionID]);

  const getPlayerColor = () => {
    return user?.id === game?.players.black ? "black" : "white";
  };

  const sendHello = () => {
    console.log("Sending hello");

    if (sessionID) {
      sendMessage(
        JSON.stringify({
          action: "hello",
          session_id: sessionID,
          data: {msg : "Hello, World!"},
        })
      );
    } else {
      console.error("No session ID");
    }
  };

  const handleMove = (move: Move) => {
    sendMessage(
      JSON.stringify({
        action: "move",
        session_id: sessionID,
        data: { move: move },
      })
    );
  };

  if (!game) return <div>WTF...</div>;

  return (
    <div className="bg-gray-900 text-white p-4 flex flex-col items-center justify-center">
      <ChessBoard game={game} playerColor={getPlayerColor()} onMove={handleMove} />
      <div className="flex justify-center items-center mt-4">
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
