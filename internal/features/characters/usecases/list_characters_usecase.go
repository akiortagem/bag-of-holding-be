package usecases

import (
	"context"
	"encoding/json"
	"fmt"

	authDom "github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
	"github.com/akiortagem/bag-of-holding-be/internal/features/characters/domain"
)

type ListCharactersUsecase struct {
	Service domain.CharacterService
}

func (u *ListCharactersUsecase) ListCharacters(ctx context.Context, payload []byte, contextUserID int64) ([]domain.CharacterResponse, error) {
	var params domain.ListCharacterParams
	if err := json.Unmarshal(payload, &params); err != nil {
		return nil, fmt.Errorf("decode list characters params: %w", err)
	}

	if contextUserID != params.UserID {
		return nil, authDom.ErrUserNotAuthorized
	}

	return u.Service.ListCharacter(ctx, params)
}
