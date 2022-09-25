package http

type RegistrationForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PaginatedLoginHistoriesForm struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}
