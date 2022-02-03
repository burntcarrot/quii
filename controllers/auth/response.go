package auth

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
