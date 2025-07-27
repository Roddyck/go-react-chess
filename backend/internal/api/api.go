package api

import (
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
