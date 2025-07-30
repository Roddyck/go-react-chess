package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Roddyck/go-react-chess/internal/api"
	"github.com/Roddyck/go-react-chess/internal/database"
	"github.com/Roddyck/go-react-chess/internal/ws"
	"github.com/Roddyck/go-react-chess/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("error opening database: %v", err)
	}

	port := os.Getenv("PORT")
	tokenSecret := os.Getenv("TOKEN_SECRET")

	dbQueries := database.New(db)

	cfg := api.New(dbQueries, tokenSecret)

	hub := ws.NewHub(dbQueries)
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	mux := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.AllowCors,
	)

	httpServer := http.Server{
		Addr:    ":" + port,
		Handler: stack(mux),
	}

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("POST /api/users", cfg.HandlerCreateUser)
	mux.HandleFunc("GET /api/users", cfg.AuthMiddleware(cfg.GetUser))
	mux.HandleFunc("POST /api/games", cfg.AuthMiddleware(cfg.HandlerGetGame))
	mux.HandleFunc("POST /api/login", cfg.HandlerLogin)

	mux.HandleFunc("POST /ws/sessions", cfg.AuthMiddleware(wsHandler.CreateSession))
	mux.HandleFunc("GET /ws/sessions", wsHandler.GetSessions)
	mux.HandleFunc("/ws/sessions/{roomID}", wsHandler.JoinSession)

	log.Println("Listening on port", port)
	log.Fatal(httpServer.ListenAndServe())
}
