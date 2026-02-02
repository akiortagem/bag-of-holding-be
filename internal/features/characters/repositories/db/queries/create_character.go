package queries

import (
	"context"
	"database/sql"
	"fmt"
)

type DatabaseCharacter struct {
	ID        int64
	Name      string
	UserID    int64
	Archetype sql.NullString
	Campaign  sql.NullString
	Game      sql.NullString
}

type DatabaseCharacterCreateParams struct {
	Name      string
	UserID    int64
	Archetype *string
	Campaign  *string
	Game      *string
}

const createCharacterQuery = `
INSERT INTO characters (name, user_id, archetype, campaign, game)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, user_id, archetype, campaign, game;
`

func CreateCharacterQuery(ctx context.Context, db *sql.DB, params DatabaseCharacterCreateParams) (DatabaseCharacter, error) {
	var character DatabaseCharacter

	if err := db.QueryRowContext(ctx, createCharacterQuery, params.Name, params.UserID, params.Archetype, params.Campaign, params.Game).Scan(
		&character.ID,
		&character.Name,
		&character.UserID,
		&character.Archetype,
		&character.Campaign,
		&character.Game,
	); err != nil {
		return DatabaseCharacter{}, fmt.Errorf("create character: %w", err)
	}

	return character, nil
}
