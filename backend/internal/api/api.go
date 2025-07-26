package api

import (
	"github.com/Roddyck/go-react-chess/backend/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func New(dbQuires *database.Queries) *apiConfig {
	return &apiConfig{
		db: dbQuires,
	}
}
