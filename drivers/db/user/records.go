package user

import "github.com/burntcarrot/quii/entity/user"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// type User struct {
// 	ID       uint   `gorm:"primary_key"`
// 	Email    string `gorm:"unique_index"`
// 	Password string `gorm:"not null"`
// 	Role     string `gorm:"not null"`
// }

func (u *User) ToDomain() user.Domain {
	return user.Domain{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}
