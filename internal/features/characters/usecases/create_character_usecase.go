package usecases

import (
	"context"
	"encoding/json"
	"fmt"

	authDom "github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
	"github.com/akiortagem/bag-of-holding-be/internal/features/characters/domain"
)

type CreateCharacterUsecase struct {
	Service domain.CharacterService
}

func (u *CreateCharacterUsecase) CreateCharacter(ctx context.Context, payload []byte, contextUserID int64) (domain.CharacterResponse, error) {
	var params domain.CharacterCreateParams
	if err := json.Unmarshal(payload, &params); err != nil {
		return domain.CharacterResponse{}, fmt.Errorf("decode create character params: %w", err)
	}

	if contextUserID != params.UserID {
		return domain.CharacterResponse{}, authDom.ErrUserNotAuthorized
	}

	return u.Service.CreateCharacter(ctx, params)
}
