package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/burntcarrot/pm/entity/user"
	"github.com/burntcarrot/pm/entity/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepo mocks.DomainRepo
var userService user.DomainService
var userDomain user.Domain

func setup() {
	userService = user.NewUsecase(&userRepo, time.Minute*15)
	userDomain = user.Domain{
		ID:       "1",
		Username: "burntcarrot",
		Email:    "burntcarrot@github.com",
		Password: "strongpass",
		Role:     "user",
	}
}

func TestLogin(t *testing.T) {
	setup()
	userRepo.On("Login", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userDomain, nil).Once()
	t.Run("Valid Login", func(t *testing.T) {
		user, err := userService.Login(context.Background(), userDomain.Email, userDomain.Password)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if user.ID != "1" {
			t.Errorf("Expected: %s, got: %s", "1", user.ID)
		}

		assert.Nil(t, err)
		assert.Equal(t, userDomain.Username, user.Username)
	})

	t.Run("Invalid Login (Empty Email)", func(t *testing.T) {
		_, err := userService.Login(context.Background(), "", userDomain.Password)
		assert.NotNil(t, err)
	})

	t.Run("Invalid Login (Empty Password)", func(t *testing.T) {
		_, err := userService.Login(context.Background(), userDomain.Email, "")
		assert.NotNil(t, err)
	})
}

func TestRegister(t *testing.T) {
	setup()
	userRepo.On("Register", mock.Anything, mock.AnythingOfType("Domain")).Return(userDomain, nil).Once()

	t.Run("Valid Register", func(t *testing.T) {
		user, err := userService.Register(context.Background(), userDomain)
		assert.Nil(t, err)
		assert.Equal(t, userDomain.Email, user.Email)
	})

	t.Run("Invalid Register", func(t *testing.T) {
		_, err := userService.Register(context.Background(), user.Domain{
			Username: "",
			Email:    "",
			Password: "",
			Role:     "",
		})

		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	setup()
	userRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()

	t.Run("Get user by ID", func(t *testing.T) {
		user, err := userService.GetByID(context.Background(), "test")
		assert.Nil(t, err)
		assert.Equal(t, userDomain.Username, user.Username)
	})
}
