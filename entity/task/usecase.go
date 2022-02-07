package task

import (
	"context"
	"errors"
	"time"

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
	if username == "" {
		return []Domain{}, errors.New("empty username")
	}

	return u.Repo.GetTasks(ctx, username, projectName)
}

func (u *Usecase) GetTaskByName(ctx context.Context, username, projectName, taskName string) ([]Domain, error) {
	if username == "" && taskName == "" {
		return []Domain{}, errors.New("empty username and task")
	}

	return u.Repo.GetTaskByName(ctx, username, projectName, taskName)
}

func (u *Usecase) CreateTask(ctx context.Context, domain Domain) (Domain, error) {
	validate := validator.New()
	err := validate.Struct(domain)

	if err != nil {
		return Domain{}, err
	}

	return u.Repo.CreateTask(ctx, domain)
}
