package domain

import "context"

type CharacterService interface {
	CreateCharacter(context context.Context, params CharacterCreateParams) (CharacterResponse, error)
}
