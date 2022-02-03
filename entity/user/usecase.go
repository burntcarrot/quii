package user

import (
	"context"
	"fmt"
	"time"

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

func (u *Usecase) Login(ctx context.Context, email, password string) (Domain, error) {
	validate := validator.New()
	err := validate.Struct(LoginDomain{Email: email, Password: password})

	if err != nil {
		return Domain{}, err
	}

	return u.Repo.Login(ctx, email, password)
}

func (u *Usecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	validate := validator.New()
	err := validate.Struct(domain)

	// TODO: remove log print
	fmt.Println(err)

	if err != nil {
		return Domain{}, err
	}

	return u.Repo.Create(ctx, domain)
}
