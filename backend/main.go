package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Roddyck/go-react-chess/backend/internal/api"
	"github.com/Roddyck/go-react-chess/backend/internal/database"
	"github.com/Roddyck/go-react-chess/backend/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("error opening database: %v", err)
	}

	dbQueries := database.New(db)

	cfg := api.New(dbQueries)

	mux := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.AllowCors,
	)

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("GET /api/games", api.HandleGame)
	mux.HandleFunc("POST /api/users", cfg.HandlerCreateUser)

	httpServer := http.Server{
		Addr:    ":" + port,
		Handler: stack(mux),
	}

	log.Println("Listening on port", port)
	log.Fatal(httpServer.ListenAndServe())
}

