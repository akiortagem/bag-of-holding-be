package queries

import (
	"context"
	"database/sql"
	"fmt"
)

type DatabaseCharacterListParams struct {
	UserID  int64
	Keyword string
	Limit   int
	Offset  int
}

const listCharactersQuery = `
SELECT id, name, user_id, archetype, campaign, game
FROM characters
WHERE user_id = $1
  AND (
    $2 = '' OR
    name ILIKE '%' || $2 || '%' OR
    COALESCE(archetype, '') ILIKE '%' || $2 || '%' OR
    COALESCE(campaign, '') ILIKE '%' || $2 || '%' OR
    COALESCE(game, '') ILIKE '%' || $2 || '%'
  )
ORDER BY name ASC, id ASC
LIMIT $3
OFFSET $4;
`

func ListCharactersQuery(ctx context.Context, db *sql.DB, params DatabaseCharacterListParams) ([]DatabaseCharacter, error) {
	rows, err := db.QueryContext(ctx, listCharactersQuery, params.UserID, params.Keyword, params.Limit, params.Offset)
	if err != nil {
		return nil, fmt.Errorf("list characters: %w", err)
	}
	defer rows.Close()

	characters := make([]DatabaseCharacter, 0)
	for rows.Next() {
		var character DatabaseCharacter
		if err := rows.Scan(
			&character.ID,
			&character.Name,
			&character.UserID,
			&character.Archetype,
			&character.Campaign,
			&character.Game,
		); err != nil {
			return nil, fmt.Errorf("list characters scan: %w", err)
		}
		characters = append(characters, character)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("list characters rows: %w", err)
	}

	return characters, nil
}
