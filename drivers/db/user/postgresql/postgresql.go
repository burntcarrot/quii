package postgresql

import (
	"context"

	dbUser "github.com/burntcarrot/pm/drivers/db/user"
	"github.com/burntcarrot/pm/entity/user"

	"gorm.io/gorm"
)

type UserRepo struct {
	Conn *gorm.DB
}

func NewUserRepo(conn *gorm.DB) user.DomainRepo {
	return &UserRepo{Conn: conn}
}

func (u *UserRepo) Create(ctx context.Context, us user.Domain) (user.Domain, error) {
	// TODO: hash password

	createdUser := dbUser.User{
		Email:    us.Email,
		Password: us.Password,
		Role:     us.Role,
	}

	insertErr := u.Conn.Create(&createdUser).Error
	if insertErr != nil {
		return user.Domain{}, insertErr
	}

	return createdUser.ToDomain(), nil
}
