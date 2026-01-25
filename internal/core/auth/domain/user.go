package domain

type UserCreateParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
