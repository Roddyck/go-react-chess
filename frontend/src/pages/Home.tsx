import { useEffect, useState } from "react";
import { useAuth } from "../context/AuthContext";
import { v4 as uuidv4 } from "uuid";
import { API_URL } from "../api/chessApi";
import { useNavigate } from "react-router";

type Sessions = Array<string>;

function Home() {
  const { user, token } = useAuth();
  const [sessions, setSessions] = useState<Sessions>([]);

  const navigate = useNavigate();

  useEffect(() => {
    const fetchSessions = async () => {
      try {
        const response = await fetch(`${API_URL}/ws/sessions`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        });

        const data = await response.json();
        console.log(data);
        setSessions(data.ids);
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
    console.log("UserID", user.id);
    navigate(`/session/${sessionID}?userID=${user.id}&username=${user.name}`);
  };

  const handleSubmit = async (e: React.SyntheticEvent) => {
    e.preventDefault();

    try {
      const response = await fetch(`${API_URL}/ws/sessions`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ id: uuidv4() }),
      });

      const data = await response.json();

      if (response.ok) {
        console.log("Session created, id: ", data.id);
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-3xl font-bold">Welcome to the Chess Game</h1>
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg"
        onClick={handleSubmit}
      >
        Create Game
      </button>
      {sessions && (
        <div>
          <h2 className="text-2xl font-bold">Active Sessions</h2>
          {sessions.map((session) => (
            <div key={session}>
              <li 
                key={session}
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg"
                onClick={() => enterSession(session)}
              >
                <div className="flex justify-between">
                  <div className="flex items-center">
                    <h3 className="text-xl font-bold text-white">{session}</h3>
                  </div>
                </div>
              </li>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export { Home };
