package auth

type RegisterRequest struct {
	ID       uint
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
