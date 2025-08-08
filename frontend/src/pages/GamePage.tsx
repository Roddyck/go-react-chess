import { useEffect, useState } from "react";
import { ChessBoard } from "../components/ChessBoard";
import type { Game, Move } from "../components/chess";
import { useParams } from "react-router";
import { useWebSocket } from "../api/websocket";
import { useAuth } from "../context/AuthContext";
import { GameEndModal } from "../components/GameEndModal";

type Session = {
  session_id: string;
  game_id: string;
};

function GamePage() {
  const { sessionID } = useParams<{ sessionID: string }>();
  const [gameID, setGameID] = useState<string | null>(null);
  const [game, setGame] = useState<Game | null>(null);
  const { user } = useAuth();
  const [showModal, setShowModal] = useState(false);

  const { sendMessage } = useWebSocket(
    `ws://localhost:8080/ws/sessions/${sessionID}?userID=${user?.id}&username=${user?.name}`,
    (msg) => {
      if (!msg) {
        console.error("No message");
      }
      console.log(msg);
      setGame(msg.data.game);

      if (
        msg.data.game.status === "white_checkmate" ||
        msg.data.game.status === "black_checkmate"
      ) {
        setShowModal(true);
      }
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
    <div className="relative flex flex-col items-center justify-center min-h-screen bg-gray-900 p-4">
      <div className="bg-gray-900 text-white p-4 flex flex-col items-center justify-center">
        <ChessBoard
          game={game}
          playerColor={getPlayerColor()}
          onMove={handleMove}
        />
      </div>
      { showModal && game.status && (
        <GameEndModal
          result={{type: "checkmate", winner: game.status === "white_checkmate" ? "white" : "black" }}
          playerColor={getPlayerColor()}
          onClose={() => setShowModal(false)}
        />
      )}
    </div>
  );
}

export { GamePage };
