package api

import (
	"encoding/json"
	"net/http"

	"github.com/Roddyck/go-react-chess/backend/internal/database"
	"github.com/Roddyck/go-react-chess/backend/internal/game"
)

func (cfg *apiConfig) HandlerCreateGame(w http.ResponseWriter, r *http.Request) {
	g := game.NewGame()

	board, err := json.Marshal(g.Board)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error marshalling board", err)
		return
	}

	history, _ := json.Marshal(g.History)
	players, _ := json.Marshal(g.Players)
	dbGame, err := cfg.db.CreateGame(r.Context(), database.CreateGameParams{
		Board:   board,
		Turn:    string(g.Turn),
		History: history,
		Players: players,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating game", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, game.Game{
		ID:      dbGame.ID,
		Board:   g.Board,
		Turn:    g.Turn,
		History: g.History,
		Players: g.Players,
	})
}
