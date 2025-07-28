import { useAuth } from "../context/AuthContext";
import { Navigate, Outlet } from "react-router";

function ProtectedRoute() {
  const { isAuthenticated } = useAuth();

  return isAuthenticated ? <Outlet /> : <Navigate to="/login" replace />;
}

export { ProtectedRoute };
