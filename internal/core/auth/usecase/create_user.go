package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
	"golang.org/x/crypto/argon2"
)

type CreateUserUsecase struct {
	Service domain.AuthService
}

func (u *CreateUserUsecase) CreateUser(ctx context.Context, payload []byte) (domain.UserCreateResponse, error) {
	var params domain.UserCreateParams
	if err := json.Unmarshal(payload, &params); err != nil {
		return domain.UserCreateResponse{}, fmt.Errorf("decode create user params: %w", err)
	}

	hashedPassword, err := hashPasswordArgon2id(params.Password)
	if err != nil {
		return domain.UserCreateResponse{}, fmt.Errorf("hash password: %w", err)
	}
	params.Password = hashedPassword

	return u.Service.CreateUser(ctx, params)
}

const (
	argonTime    uint32 = 3
	argonMemory  uint32 = 64 * 1024
	argonThreads uint8  = 2
	argonKeyLen  uint32 = 32
	argonSaltLen        = 16
)

func hashPasswordArgon2id(password string) (string, error) {
	salt := make([]byte, argonSaltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate password salt: %w", err)
	}

	hash := argon2.IDKey([]byte(password), salt, argonTime, argonMemory, argonThreads, argonKeyLen)
	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)
	hashEncoded := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		argonMemory,
		argonTime,
		argonThreads,
		saltEncoded,
		hashEncoded,
	)

	return encoded, nil
}
