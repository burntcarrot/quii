package redis

import (
	"context"
	"encoding/json"

	dbUser "github.com/burntcarrot/pm/drivers/db/user"
	"github.com/burntcarrot/pm/entity/user"
	"github.com/burntcarrot/pm/helpers"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type UserRepo struct {
	Conn *redis.Client
}

func NewUserRepo(conn *redis.Client) user.DomainRepo {
	return &UserRepo{Conn: conn}
}

func (u *UserRepo) Login(ctx context.Context, username, password string) (user.Domain, error) {
	raw, err := u.Conn.Get(ctx, username).Result()
	if err != nil {
		return user.Domain{}, err
	}

	uu := new(dbUser.User)

	if err := json.Unmarshal([]byte(raw), uu); err != nil {
		return user.Domain{}, err
	}

	us := dbUser.User{
		ID:       uu.ID,
		Username: uu.Username,
		Email:    uu.Email,
		Password: uu.Password,
		Role:     uu.Role,
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
		ID:       uuid.New().String(),
		Username: us.Username,
		Email:    us.Email,
		Password: hashedPassword,
		Role:     us.Role,
	}

	raw, err := json.Marshal(createdUser)
	if err != nil {
		return user.Domain{}, err
	}

	insertErr := u.Conn.Set(ctx, createdUser.Username, raw, 0).Err()
	if insertErr != nil {
		return user.Domain{}, insertErr
	}

	return createdUser.ToDomain(), nil
}

func (u *UserRepo) GetByID(ctx context.Context, id string) (user.Domain, error) {
	raw, err := u.Conn.Get(ctx, id).Result()
	if err != nil {
		return user.Domain{}, err
	}

	uu := new(dbUser.User)

	if err := json.Unmarshal([]byte(raw), uu); err != nil {
		return user.Domain{}, err
	}

	us := dbUser.User{
		Username: uu.Username,
		Email:    uu.Email,
		Role:     uu.Role,
	}

	return us.ToDomain(), nil
}
