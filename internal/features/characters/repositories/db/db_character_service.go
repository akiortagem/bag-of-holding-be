package db

import (
	"context"
	"database/sql"

	"github.com/akiortagem/bag-of-holding-be/internal/features/characters/domain"
	"github.com/akiortagem/bag-of-holding-be/internal/features/characters/repositories/db/queries"
)

type DBCharacterService struct {
	Db *sql.DB
}

func nullStringToPtr(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}

	str := value.String
	return &str
}

func (s *DBCharacterService) CreateCharacter(ctx context.Context, params domain.CharacterCreateParams) (domain.CharacterResponse, error) {
	dbCharacter, err := queries.CreateCharacterQuery(ctx, s.Db, queries.DatabaseCharacterCreateParams{
		Name:      params.Name,
		UserID:    params.UserID,
		Archetype: params.Archetype,
		Campaign:  params.Campaign,
		Game:      params.Game,
	})
	if err != nil {
		return domain.CharacterResponse{}, err
	}

	return domain.CharacterResponse{
		Name:      dbCharacter.Name,
		UserID:    dbCharacter.UserID,
		Archetype: nullStringToPtr(dbCharacter.Archetype),
		Campaign:  nullStringToPtr(dbCharacter.Campaign),
		Game:      nullStringToPtr(dbCharacter.Game),
	}, nil
}

func (s *DBCharacterService) ListCharacter(ctx context.Context, params domain.ListCharacterParams) ([]domain.CharacterResponse, error) {
	page := params.Page
	if page < 1 {
		page = 1
	}

	pageSize := params.PageSize
	if pageSize < 1 {
		pageSize = 20
	}

	keyword := ""
	if params.Keyword != nil {
		keyword = *params.Keyword
	}

	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	characters, err := queries.ListCharactersQuery(ctx, s.Db, queries.DatabaseCharacterListParams{
		UserID:  params.UserID,
		Keyword: keyword,
		Limit:   pageSize,
		Offset:  offset,
	})
	if err != nil {
		return nil, err
	}

	responses := make([]domain.CharacterResponse, 0, len(characters))
	for _, character := range characters {
		responses = append(responses, domain.CharacterResponse{
			Name:      character.Name,
			UserID:    character.UserID,
			Archetype: nullStringToPtr(character.Archetype),
			Campaign:  nullStringToPtr(character.Campaign),
			Game:      nullStringToPtr(character.Game),
		})
	}

	return responses, nil
}
