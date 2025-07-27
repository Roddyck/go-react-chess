package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Roddyck/go-react-chess/backend/internal/api"
	"github.com/Roddyck/go-react-chess/backend/internal/database"
	"github.com/Roddyck/go-react-chess/backend/internal/websocket"
	"github.com/Roddyck/go-react-chess/backend/middleware"
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

	hub := websocket.NewHub()
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
	mux.HandleFunc("GET /api/games", cfg.HandlerCreateGame)
	mux.HandleFunc("POST /api/users", cfg.HandlerCreateUser)
	mux.HandleFunc("POST /api/login", cfg.HandlerLogin)
	mux.HandleFunc("/ws/games/{id}", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandlerWebSocket(hub, w, r)
	})

	log.Println("Listening on port", port)
	log.Fatal(httpServer.ListenAndServe())
}
