package usecase

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

func HashPassword(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

func CheckPasswordHash(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}

func MakeJWT(userID int64, expiresIn time.Duration, issuer string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		Issuer:    issuer,
		Subject:   strconv.FormatInt(userID, 10),
	})

	return token.SignedString([]byte(secret))
}

func MakeRefreshToken() (string, error) {
	rando := make([]byte, 32)
	_, err := rand.Read(rando)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(rando), nil
}
