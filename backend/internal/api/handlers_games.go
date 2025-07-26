package api

import (
	"net/http"

	"github.com/Roddyck/go-react-chess/backend/internal/game"
)

func HandleGame(w http.ResponseWriter, r *http.Request) {
	game := game.NewGame()

	respondWithJSON(w, http.StatusOK, game)
}
