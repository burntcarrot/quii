package redis

import (
	"context"
	"math/rand"

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

func (u *UserRepo) Login(ctx context.Context, email, password string) (user.Domain, error) {
	pass, err := u.Conn.Get(ctx, email).Result()
	if err != nil {
		return user.Domain{}, err
	}

	us := dbUser.User{
		Password: pass,
	}

	return us.ToDomain(), nil
}

func (u *UserRepo) Create(ctx context.Context, us user.Domain) (user.Domain, error) {
	// hash password
	hashedPassword, err := helpers.HashPassword(us.Password)
	if err != nil {
		return user.Domain{}, err
	}

	createdUser := dbUser.User{
		// TODO: Use UUID or something like that for other DBs
		// Since Redis is a test DB, I'm using Random Integer
		ID:       uint(rand.Int()),
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
