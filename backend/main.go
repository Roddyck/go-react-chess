package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Roddyck/go-react-chess/backend/internal/game"
	"github.com/Roddyck/go-react-chess/backend/middleware"
)

func main() {
	mux := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.AllowCors,
	)

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("/api/games", HandleGame)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: stack(mux),
	}

	log.Println("Listening on port 8080")
	log.Fatal(httpServer.ListenAndServe())
}

func HandleGame(w http.ResponseWriter, r *http.Request) {
	game := game.NewGame()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}
