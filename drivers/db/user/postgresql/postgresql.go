package postgresql

import (
	"context"

	dbUser "github.com/burntcarrot/pm/drivers/db/user"
	"github.com/burntcarrot/pm/entity/user"
	"github.com/burntcarrot/pm/helpers"

	"gorm.io/gorm"
)

type UserRepo struct {
	Conn *gorm.DB
}

func NewUserRepo(conn *gorm.DB) user.DomainRepo {
	return &UserRepo{Conn: conn}
}

func (u *UserRepo) Register(ctx context.Context, us user.Domain) (user.Domain, error) {
	// TODO: hash password
	hashedPassword, err := helpers.HashPassword(us.Password)
	if err != nil {
		return user.Domain{}, err
	}

	createdUser := dbUser.User{
		Email:    us.Email,
		Password: hashedPassword,
		Role:     us.Role,
	}

	insertErr := u.Conn.Create(&createdUser).Error
	if insertErr != nil {
		return user.Domain{}, insertErr
	}

	return createdUser.ToDomain(), nil
}

func (u *UserRepo) Login(ctx context.Context, email, password string) (user.Domain, error) {
	var us dbUser.User
	err := u.Conn.Where("email = ?", email).First(&us).Error
	if err != nil {
		return user.Domain{}, err
	}

	return us.ToDomain(), nil
}

func (u *UserRepo) GetByID(ctx context.Context, id string) (user.Domain, error) {
	var us dbUser.User
	if err := u.Conn.Where("id = ?", id).First(&us).Error; err != nil {
		return user.Domain{}, nil
	}

	return us.ToDomain(), nil
}
