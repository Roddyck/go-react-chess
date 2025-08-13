import { useEffect, useMemo, useState } from "react";
import { useAuth } from "../context/AuthContext";
import { useNavigate } from "react-router";
import { authFetch } from "../api/authFetch";
import { API_URL } from "../api/chessApi";
import { v4 as uuidv4 } from "uuid";
import type { Session } from "./types";

function Home() {
  const { user } = useAuth();
  const [sessions, setSessions] = useState<Session[]>([]);
  const [sessionName, setSessionName] = useState("");
  const [showNameInput, setShowNameInput] = useState(false);
  const visibleSessions = useMemo(() => {
    return sessions?.filter((session) => session.status !== "full");
  }, [sessions]);

  const navigate = useNavigate();

  useEffect(() => {
    const fetchSessions = async () => {
      try {
        const response = await authFetch(`${API_URL}/ws/sessions`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        });

        const data = await response.json();
        console.log(data);
        setSessions(data);
      } catch (error) {
        console.error(error);
      }
    };

    fetchSessions();
  }, []);

  const enterSession = (sessionID: string) => {
    if (!user) {
      console.error("User not logged in");
      return;
    }
    navigate(`/session/${sessionID}?userID=${user.id}&username=${user.name}`);
  };

  const handleSubmit = async (e: React.SyntheticEvent) => {
    e.preventDefault();

    try {
      const response = await authFetch(`${API_URL}/ws/sessions`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ id: uuidv4(), name: sessionName }),
      });

      const data = await response.json();

      if (response.ok) {
        console.log("Session created, id: ", data.id);
        navigate(`/session/${data.id}`);
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen bg-gray-900 text-white">
      <h1 className="text-3xl font-bold mb-4">Welcome to the Chess Game</h1>
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg"
        onClick={() => setShowNameInput(true)}
      >
        Create Game
      </button>

      {showNameInput && (
        <div className="fixed inset-0 z-50 flex flex-col items-center justify-center min-h-screen bg-gray-900 bg-opacity-50">
          <input
            className="mb-4 rounded-lg border-2 border-gray-500 px-4 py-2"
            type="text"
            placeholder="Session name"
            value={sessionName}
            onChange={(e) => setSessionName(e.target.value)}
          />
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg"
            type="submit"
            onClick={handleSubmit}
          >
            Create Game
          </button>
        </div>
      )}

      <div className="mt-4">
        {visibleSessions && (
          <div className="items-center justify-center flex flex-col">
            <h2 className="text-2xl font-bold">Active Sessions</h2>
            <ul className="list-none">
              {visibleSessions.map((session) => (
                <div key={session.session_id} className="flex flex-col mb-1">
                  <li
                    key={session.session_id}
                    className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg"
                    onClick={() => enterSession(session.session_id)}
                  >
                    <div className="flex justify-between">
                      <div className="flex items-center">
                        <h3 className="text-xl font-bold text-white text-center">
                          {session.name}
                        </h3>
                      </div>
                    </div>
                  </li>
                </div>
              ))}
            </ul>
          </div>
        )}
      </div>
    </div>
  );
}

export { Home };
