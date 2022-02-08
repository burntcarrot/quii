package project

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

func (u *Usecase) GetProjects(ctx context.Context, username string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	if username == "" {
		return []Domain{}, errors.ErrValidationFailed
	}

	return u.Repo.GetProjects(ctx, username)
}

func (u *Usecase) GetProjectByID(ctx context.Context, username, projectID string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	if username == "" && projectID == "" {
		return []Domain{}, errors.ErrValidationFailed
	}

	return u.Repo.GetProjectByID(ctx, username, projectID)
}

func (u *Usecase) CreateProject(ctx context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	validate := validator.New()
	err := validate.Struct(domain)

	if err != nil {
		return Domain{}, errors.ErrValidationFailed
	}

	return u.Repo.CreateProject(ctx, domain)
}
