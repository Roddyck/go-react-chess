package api

import (
	"context"
	"net/http"

	"github.com/Roddyck/go-react-chess/backend/internal/auth"
	"github.com/Roddyck/go-react-chess/backend/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db          *database.Queries
	TokenSecret string
}

func New(dbQuires *database.Queries, tokenSecret string) *apiConfig {
	return &apiConfig{
		db:          dbQuires,
		TokenSecret: tokenSecret,
	}
}

func (cfg *apiConfig) AuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Access token is missing from request headers", err)
			return
		}

		userID, err := auth.ValidateJWT(token, cfg.TokenSecret)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid access token", err)
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		handler(w, r.WithContext(ctx))
	}
}
