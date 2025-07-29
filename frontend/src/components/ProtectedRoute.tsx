import { useAuth } from "../context/AuthContext";
import { Navigate, Outlet } from "react-router";

function ProtectedRoute() {
  const { isAuthenticated, loading } = useAuth();

  if (loading) return <div>Loading...</div>;

  return isAuthenticated ? <Outlet /> : <Navigate to="/login" replace />;
}

export { ProtectedRoute };
