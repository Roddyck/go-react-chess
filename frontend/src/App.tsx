import { BrowserRouter as Router, Routes, Route } from "react-router";
import { Home } from "./pages/Home";
import { Login } from "./pages/Login";
import { Register } from "./pages/Register";
import { ProtectedRoute } from "./components/ProtectedRoute";
import { AuthProvider } from "./context/AuthContext";
import { GamePage } from "./pages/GamePage";
import { Navbar } from "./components/Navbar";

function App() {
  return (
    <Router>
      <AuthProvider>
        <Navbar />
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route element={<ProtectedRoute />}>
            <Route path="/" element={<Home />} />
            <Route path="/session/:sessionID" element={<GamePage />} />
          </Route>
        </Routes>
      </AuthProvider>
    </Router>
  );
}

export default App;
