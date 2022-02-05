package project

import "context"

type Domain struct {
	ID          string
	Username    string `validate:"required"`
	Name        string `validate:"required"`
	Description string
	Github      string `validate:"required"`
}

type DomainRepo interface {
	GetProjects(ctx context.Context, username string) ([]Domain, error)
	CreateProject(ctx context.Context, domain Domain) (Domain, error)
}
