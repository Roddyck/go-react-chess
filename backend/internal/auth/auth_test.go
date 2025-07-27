package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	err = CheckPasswordHash(password, hash)
	if err != nil {
		t.Errorf("error checking password hash: %v", err)
	}
}

func TestMakeJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "secret"
	token, err := MakeJWT(userID, tokenSecret, time.Hour)
	if err != nil {
		t.Errorf("error making JWT: %v", err)
	}

	id, err := ValidateJWT(token, tokenSecret)
	if err != nil {
		t.Errorf("error validating JWT: %v", err)
	}

	if id != userID {
		t.Errorf("expected user id %s, got %s", userID, id)
	}
}
