package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed), err
}

func CompareHashAndPassord(hashed, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return errors.New("Invalid Credential")
	}
	return nil
}

func HashAPIKey(rawKey string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(rawKey))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateRandomString(length int) string {
	byteLen := (length * 3) / 4

	b := make([]byte, byteLen)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.EncodeToString(b)[:length]
}

func GenerateAPIKey() string {
	return "ak_" + GenerateRandomString(32)
}
