package api

import (
	"encoding/json"
	"net/http"

	"github.com/Roddyck/go-react-chess/util"
	"github.com/google/uuid"
)

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
		return
	}

	g, err := cfg.db.GetGameByID(r.Context(), gameUUID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error getting game", err)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, g)
}
