package domain

import "context"

type CharacterService interface {
	CreateCharacter(context context.Context, params CharacterCreateParams) (CharacterResponse, error)
	ListCharacter(context context.Context, params ListCharacterParams) ([]CharacterResponse, error)
}
