import { useEffect, useState } from "react";
import { ChessBoard } from "../components/ChessBoard";
import type { Game, Move } from "../components/chess";
import { useNavigate, useParams } from "react-router";
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
  const [isDrawOffer, setIsDrawOffer] = useState(false);
  const { user } = useAuth();
  const [showModal, setShowModal] = useState(false);
  const navigate = useNavigate();

  const { sendMessage } = useWebSocket(
    `ws://localhost:8080/ws/sessions/${sessionID}?userID=${user?.id}&username=${user?.name}`,
    (msg) => {
      if (!msg) {
        console.error("No message");
      }
      console.log(msg);
      setGame(msg.data.game);

      if (msg.data.game.status != "active") {
        setShowModal(true);
      }
      if (msg.action === "draw_offer" && msg.data.user_id !== user?.id) {
        setIsDrawOffer(true);
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

  const handleModalClose = () => {
    setShowModal(false);
    navigate("/");
  };

  const handleDrawOffer = () => {
    sendMessage(
      JSON.stringify({
        action: "draw_offer",
        session_id: sessionID,
        data: { user_id: user?.id },
      })
    );
  };

  const handleAcceptDraw = () => {
    sendMessage(
      JSON.stringify({
        action: "draw_accept",
        session_id: sessionID,
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
      <div className="bg-gray-900 text-white p-4 flex flex-col items-center justify-center">
        {isDrawOffer ? (
          <button
            className="bg-blue-900 text-white p-2 rounded-md border-2 border-blue-500 hover:bg-blue-700"
            onClick={handleAcceptDraw}
          >
            Accept Draw
          </button>
        ) : (
          <button
            className="bg-blue-900 text-white p-2 rounded-md border-2 border-blue-500 hover:bg-blue-700"
            onClick={handleDrawOffer}
          >
            Offer Draw
          </button>
        )}
      </div>
      {showModal && game.status && (
        <GameEndModal
          status={game.status}
          turn={game.turn}
          playerColor={getPlayerColor()}
          onClose={handleModalClose}
        />
      )}
    </div>
  );
}

export { GamePage };
