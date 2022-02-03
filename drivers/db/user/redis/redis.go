package redis

import (
	"context"

	dbUser "github.com/burntcarrot/pm/drivers/db/user"
	"github.com/burntcarrot/pm/entity/user"
	"github.com/burntcarrot/pm/helpers"
	"github.com/go-redis/redis/v8"
)

type UserRepo struct {
	Conn *redis.Client
}

func NewUserRepo(conn *redis.Client) user.DomainRepo {
	return &UserRepo{Conn: conn}
}

func (u *UserRepo) Create(ctx context.Context, us user.Domain) (user.Domain, error) {
	// hash password
	hashedPassword, err := helpers.HashPassword(us.Password)
	if err != nil {
		return user.Domain{}, err
	}

	createdUser := dbUser.User{
		Email:    us.Email,
		Password: hashedPassword,
		Role:     us.Role,
	}

	insertErr := u.Conn.Set(ctx, us.Email, hashedPassword, 0).Err()
	if insertErr != nil {
		return user.Domain{}, insertErr
	}

	return createdUser.ToDomain(), nil
}
