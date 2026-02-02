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
