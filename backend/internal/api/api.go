package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Roddyck/go-react-chess/backend/internal/game"
	"github.com/Roddyck/go-react-chess/backend/middleware"
)

func StartServer(port string) error {
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
		Addr:    ":" + port,
		Handler: stack(mux),
	}

	log.Println("Listening on port 8080")
	return httpServer.ListenAndServe()
}

func HandleGame(w http.ResponseWriter, r *http.Request) {
	game := game.NewGame()

	respondWithJSON(w, http.StatusOK, game)
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
	    return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing payload: %v", err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string, err error) {
	type errorPayload struct {
		Error string `json:"error"`
	}

	resp := errorPayload{
		Error: fmt.Sprintf("%s: %s", message, err.Error()),
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling error payload: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing error payload: %v", err)
	}
}
