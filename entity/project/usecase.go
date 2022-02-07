package project

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

func (u *Usecase) GetProjects(ctx context.Context, username string) ([]Domain, error) {
	if username == "" {
		return []Domain{}, errors.New("empty username")
	}

	return u.Repo.GetProjects(ctx, username)
}

func (u *Usecase) GetProjectByName(ctx context.Context, username, projectID string) ([]Domain, error) {
	if username == "" && projectID == "" {
		return []Domain{}, errors.New("empty username and project")
	}

	return u.Repo.GetProjectByName(ctx, username, projectID)
}

func (u *Usecase) CreateProject(ctx context.Context, domain Domain) (Domain, error) {
	validate := validator.New()
	err := validate.Struct(domain)

	if err != nil {
		return Domain{}, err
	}

	return u.Repo.CreateProject(ctx, domain)
}
