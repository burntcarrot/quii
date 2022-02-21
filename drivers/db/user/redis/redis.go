package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	dbUser "github.com/burntcarrot/quii/drivers/db/user"
	"github.com/burntcarrot/quii/entity/user"
	"github.com/burntcarrot/quii/errors"
	"github.com/burntcarrot/quii/helpers"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserRepo struct {
	Conn   *redis.Client
	Logger *zap.SugaredLogger
}

func NewUserRepo(conn *redis.Client, logger *zap.SugaredLogger) user.DomainRepo {
	return &UserRepo{Conn: conn, Logger: logger}
}

func (u *UserRepo) Login(ctx context.Context, username, password string) (user.Domain, error) {
	raw, err := u.Conn.Get(ctx, strings.ToLower(username)).Result()
	if err != nil {
		u.Logger.Errorf("[login] user not found: %s", username)
		return user.Domain{}, errors.ErrInternalServerError
	}

	uu := new(dbUser.User)

	if err := json.Unmarshal([]byte(raw), uu); err != nil {
		u.Logger.Errorf("[login] failed to unmarshal user data: %v", err)
		return user.Domain{}, errors.ErrInternalServerError
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
		u.Logger.Error("[register] failed to hash password")
		return user.Domain{}, errors.ErrInternalServerError
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
		u.Logger.Error("[register] failed to marshal user data")
		return user.Domain{}, errors.ErrInternalServerError
	}

	insertErr := u.Conn.Set(ctx, createdUser.Username, raw, 0).Err()
	if insertErr != nil {
		u.Logger.Errorf("[register] failed to set username key in redis: %v", err)
		return user.Domain{}, errors.ErrInternalServerError
	}

	// set counter for project while creating user itself
	counter := fmt.Sprintf("%s:projects:counter", createdUser.Username)
	counterErr := u.Conn.Set(ctx, counter, 1, 0).Err()
	if counterErr != nil {
		u.Logger.Error("[register] failed to set project counter in redis")
		return user.Domain{}, errors.ErrInternalServerError
	}

	return createdUser.ToDomain(), nil
}

func (u *UserRepo) GetByName(ctx context.Context, id string) (user.Domain, error) {
	raw, err := u.Conn.Get(ctx, id).Result()
	if err != nil {
		u.Logger.Error("[getbyname] failed to get user key in redis")
		return user.Domain{}, errors.ErrInternalServerError
	}

	uu := new(dbUser.User)

	if err := json.Unmarshal([]byte(raw), uu); err != nil {
		u.Logger.Errorf("[getbyname] failed to unmarshal user data: %v", err)
		return user.Domain{}, errors.ErrInternalServerError
	}

	us := dbUser.User{
		Username: uu.Username,
		Email:    uu.Email,
		Role:     uu.Role,
	}

	return us.ToDomain(), nil
}
