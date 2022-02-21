package user

import (
	"context"
	"time"

	"github.com/burntcarrot/quii/errors"
	"github.com/go-playground/validator/v10"
)

type Usecase struct {
	Repo       DomainRepo
	ctxTimeout time.Duration
}

func NewUsecase(repo DomainRepo, timeout time.Duration) *Usecase {
	return &Usecase{
		Repo:       repo,
		ctxTimeout: timeout,
	}
}

func (u *Usecase) Login(ctx context.Context, username, password string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	validate := validator.New()
	err := validate.Struct(LoginDomain{Username: username, Password: password})

	if err != nil {
		return Domain{}, errors.ErrValidationFailed
	}

	return u.Repo.Login(ctx, username, password)
}

func (u *Usecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	validate := validator.New()
	err := validate.Struct(domain)

	if err != nil {
		return Domain{}, errors.ErrValidationFailed
	}

	return u.Repo.Register(ctx, domain)
}

func (u *Usecase) GetByName(ctx context.Context, id string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.Repo.GetByName(ctx, id)
}
