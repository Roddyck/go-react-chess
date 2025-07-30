package api

import (
	"encoding/json"
	"net/http"

	"github.com/Roddyck/go-react-chess/internal/database"
	"github.com/Roddyck/go-react-chess/internal/game"
	"github.com/Roddyck/go-react-chess/util"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerCreateGame(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		util.RespondWithError(w, http.StatusInternalServerError, "user id not found in context", nil)
		return
	}

	g := game.NewGame()

	board, err := json.Marshal(g.Board)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error marshalling board", err)
		return
	}

	g.Players[game.White] = userID

	history, _ := json.Marshal(g.History)
	players, _ := json.Marshal(g.Players)
	dbGame, err := cfg.db.CreateGame(r.Context(), database.CreateGameParams{
		Board:   board,
		Turn:    string(g.Turn),
		History: history,
		Players: players,
	})
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error creating game", err)
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, game.Game{
		ID:      dbGame.ID,
		Board:   g.Board,
		Turn:    g.Turn,
		History: g.History,
		Players: g.Players,
	})
}

func (cfg *apiConfig) HandlerGetGame(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID string `json:"id"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "error decoding game id", err)
	    return
	}

	gameUUID, err := uuid.Parse(params.ID)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "invalid game id", err)
	}

	game, err := cfg.db.GetGameByID(r.Context(), gameUUID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error getting game", err)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, game)
}
