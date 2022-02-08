package task

import (
	"context"
	"time"

	"github.com/burntcarrot/pm/errors"
	"github.com/go-playground/validator"
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

func (u *Usecase) GetTasks(ctx context.Context, username, projectName string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	if username == "" {
		return []Domain{}, errors.ErrValidationFailed
	}

	return u.Repo.GetTasks(ctx, username, projectName)
}

func (u *Usecase) GetTaskByName(ctx context.Context, username, projectName, taskName string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	if username == "" && taskName == "" {
		return []Domain{}, errors.ErrValidationFailed
	}

	return u.Repo.GetTaskByName(ctx, username, projectName, taskName)
}

func (u *Usecase) CreateTask(ctx context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	validate := validator.New()
	err := validate.Struct(domain)

	if err != nil {
		return Domain{}, errors.ErrValidationFailed
	}

	return u.Repo.CreateTask(ctx, domain)
}
