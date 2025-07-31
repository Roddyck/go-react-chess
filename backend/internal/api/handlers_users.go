package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Roddyck/go-react-chess/internal/auth"
	"github.com/Roddyck/go-react-chess/internal/database"
	"github.com/Roddyck/go-react-chess/util"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	creds := credentials{}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	hash, err := auth.HashPassword(creds.Password)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Error hashing password", err)
		return
	}

	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		Name:           creds.Name,
		Email:          creds.Email,
		HashedPassword: hash,
	})

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt.UTC(),
		UpdatedAt: user.UpdatedAt.UTC(),
		Name:      user.Name,
		Email:     user.Email,
	})
}

func (cfg *apiConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	user, err := cfg.db.GetUserByID(r.Context(), userID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Error getting user", err)
	}

	util.RespondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt.UTC(),
		UpdatedAt: user.UpdatedAt.UTC(),
		Name:      user.Name,
		Email:     user.Email,
	})
}

func (cfg *apiConfig) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type loginResponse struct {
		User
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	creds := credentials{}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), creds.Email)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(creds.Password, user.HashedPassword)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, cfg.TokenSecret, time.Hour)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Error creating access token", err)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Error creating refresh token", err)
		return
	}

	_, err = cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 60).UTC(),
	})
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Error creating refresh token", err)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, loginResponse{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt.UTC(),
			UpdatedAt: user.UpdatedAt.UTC(),
			Name:      user.Name,
			Email:     user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
