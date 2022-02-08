package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
	raw, err := u.Conn.Get(ctx, strings.ToLower(username)).Result()
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

func (u *UserRepo) Register(ctx context.Context, us user.Domain) (user.Domain, error) {
	// hash password
	hashedPassword, err := helpers.HashPassword(us.Password)
	if err != nil {
		return user.Domain{}, err
	}

	createdUser := dbUser.User{
		ID:       uuid.New().String(),
		Username: strings.ToLower(us.Username),
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

	// set counter for project while creating user itself
	counter := fmt.Sprintf("%s:projects:counter", createdUser.Username)
	counterErr := u.Conn.Set(ctx, counter, 1, 0).Err()
	fmt.Println("counter err:", counterErr)
	if counterErr != nil {
		return user.Domain{}, counterErr
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
