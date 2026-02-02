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
