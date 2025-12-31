package utils

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

type JWTManager struct {
	auth *jwtauth.JWTAuth
}

type Claims struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
}

func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{
		auth: jwtauth.New("HS256", []byte(secret), nil),
	}
}

func (j *JWTManager) GenerateToken(userId, email string) (string, error) {
	claims := map[string]any{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	_, token, err := j.auth.Encode(claims)
	return token, err
}

func (j *JWTManager) Middleware(next http.Handler) http.Handler {
	verifier := jwtauth.Verifier(j.auth)
	authenticator := jwtauth.Authenticator(j.auth)
	return verifier(authenticator(next))
}

func FromContext(ctx context.Context) (*Claims, bool) {
	_, rawClaims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return nil, false
	}

	userID, ok := rawClaims["user_id"].(string)
	if !ok {
		return nil, false
	}

	email, _ := rawClaims["email"].(string)

	return &Claims{
		UserId: userID,
		Email:  email,
	}, true
}
