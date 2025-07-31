import { createContext, useState, useContext, useEffect } from "react";
import type { User, AuthContextType } from "./types";
import { useNavigate } from "react-router";
import { API_URL } from "../api/chessApi";
import { authFetch, clearAuthTokens, setAuthTokens } from "../api/authFetch";

const AuthContext = createContext<AuthContextType | undefined>(undefined);

function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [token, setToken] = useState<string | null>(
    localStorage.getItem("accessToken")
  );
  const navigate = useNavigate();

  useEffect(() => {
    const loadUser = async () => {
      if (token) {
        try {
          const response = await authFetch(`${API_URL}/api/users`);

          const data = await response.json();
          if (response.status !== 200) {
            clearAuthTokens();
            setUser(null);
            navigate("/login");
          }

          setUser(data);
        } catch (error) {
          clearAuthTokens();
          setUser(null);
          navigate("/login");
        }
      }
      setLoading(false);
    };

    loadUser();
  }, [token]);

  const login = async (email: string, password: string) => {
    const response = await fetch("http://localhost:8080/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    const data = await response.json();

    if (response.status !== 200) {
      console.error(data);
      return;
    }

    setUser(data);
    setAuthTokens(data.access_token, data.refresh_token);
    setToken(data.access_token);
    navigate("/");
  };

  const register = async (name: string, email: string, password: string) => {
    const response = await fetch("http://localhost:8080/api/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name, email, password }),
    });

    const data = await response.json();
    if (response.status !== 200) {
      console.error(data);
      return;
    }

    setUser(data);
    setToken(data.access_token);
    localStorage.setItem("accessToken", data.access_token);

    navigate("/");
  };

  const logout = () => {
    clearAuthTokens();
    setToken(null);
    setUser(null);
    navigate("/login");
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        login,
        register,
        logout,
        isAuthenticated: !!user,
        loading,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within a AuthProvider");
  }
  return context;
}

export { AuthProvider, useAuth };
