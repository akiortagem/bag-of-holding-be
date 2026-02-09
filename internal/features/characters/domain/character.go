package domain

type CharacterCreateParams struct {
	Name      string  `json:"name"`
	UserID    int64   `json:"user_id"`
	Archetype *string `json:"archetype,omitempty"`
	Campaign  *string `json:"campaign,omitempty"`
	Game      *string `json:"game,omitempty"`
}

type CharacterResponse struct {
	Name      string  `json:"name"`
	UserID    int64   `json:"user_id"`
	Archetype *string `json:"archetype,omitempty"`
	Campaign  *string `json:"campaign,omitempty"`
	Game      *string `json:"game,omitempty"`
}

type ListCharacterParams struct {
	UserID   int64   `json:"user_id"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
	Keyword  *string `json:"keyword,omitempty"`
}
