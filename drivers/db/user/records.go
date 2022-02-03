package user

import "github.com/burntcarrot/pm/entity/user"

type User struct {
	ID       uint   `gorm:"primary_key"`
	Email    string `gorm:"unique_index"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
}

func (u *User) ToDomain() user.Domain {
	return user.Domain{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}
