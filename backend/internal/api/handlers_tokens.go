package api

import (
	"net/http"
	"time"

	"github.com/Roddyck/go-react-chess/internal/auth"
	"github.com/Roddyck/go-react-chess/util"
)

func (cfg *apiConfig) HandlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		AccessToken string `json:"access_token"`
	}

	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "refresh token was not provided in headers", err)
		return
	}

	dbToken, err := cfg.db.GetRefreshToken(r.Context(), refreshToken)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "given refresh token does not exist", err)
		return
	}

	if dbToken.ExpiresAt.Before(time.Now()) {
		util.RespondWithError(w, http.StatusUnauthorized, "given refresh token has expired", err)
		return
	}

	if dbToken.RevokedAt.Valid {
		util.RespondWithError(w, http.StatusUnauthorized, "given refresh token has been revoked", err)
		return
	}

	accessToken, err := auth.MakeJWT(dbToken.UserID, cfg.TokenSecret, time.Minute*15)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error creating access token", err)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, response{AccessToken: accessToken})
}

func (cfg *apiConfig) HandlerRevoke(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "refresh token was not provided in headers", err)
		return
	}

	err = cfg.db.RevokeRefreshToken(r.Context(), refreshToken)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error revoking refresh token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
