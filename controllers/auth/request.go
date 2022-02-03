package auth

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	ID       uint
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
