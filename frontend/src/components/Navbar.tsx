import { useAuth } from "../context/AuthContext";

function Navbar() {
  const { user, isAuthenticated, logout } = useAuth();

  return (
    <div className="flex justify-between items-center bg-gray-900 text-white p-4">
      <h1
        className="text-3xl font-bold cursor-pointer"
        onClick={() => (window.location.href = "/")}
      >
        GRChess
      </h1>
      {isAuthenticated ? (
        <div className="flex items-center">
          <h3 className="text-xl font-bold text-white mr-4">{user?.name}</h3>
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg"
            onClick={logout}
          >
            Logout
          </button>
        </div>
      ) : (
        <div className="flex items-center">
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg"
            onClick={() => (window.location.href = "/login")}
          >
            Login
          </button>
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg ml-4"
            onClick={() => (window.location.href = "/register")}
          >
            Register
          </button>
        </div>
      )}
    </div>
  );
}

export { Navbar };
