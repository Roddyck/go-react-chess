import { useAuth } from "../context/AuthContext";
import { v4 as uuidv4 } from "uuid";
import { API_URL } from "../api/chessApi";
import { useNavigate } from "react-router";

function Home() {
  const { user, token } = useAuth();

  const navigate = useNavigate();

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
        navigate(`/session/${data.id}?userID=${user?.id}&username=${user?.name}`);
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
    </div>
  );
}

export { Home };
