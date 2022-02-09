package task

import "context"

type Domain struct {
	ID          string
	Username    string `validate:"required"`
	ProjectName string `validate:"required"`
	Type        string `validate:"required"`
	Name        string `validate:"required"`
	Deadline    string
	Status      string `validate:"required"`
}

type DomainRepo interface {
	GetTasks(ctx context.Context, username, projectName string) ([]Domain, error)
	GetTaskByName(ctx context.Context, username, projectName, taskName string) ([]Domain, error)
	GetTaskByID(ctx context.Context, username, projectName, taskID string) ([]Domain, error)
	CreateTask(ctx context.Context, domain Domain) (Domain, error)
}

type DomainService interface {
	GetTasks(ctx context.Context, username, projectName string) ([]Domain, error)
	GetTaskByName(ctx context.Context, username, projectName, taskName string) ([]Domain, error)
	GetTaskByID(ctx context.Context, username, projectName, taskID string) ([]Domain, error)
	CreateTask(ctx context.Context, domain Domain) (Domain, error)
}
