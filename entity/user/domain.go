package user

import "context"

type Domain struct {
	ID       string
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=50"`
	Role     string `validate:"required"`
}

type LoginDomain struct {
	Username string `validate:"required"`
	Password string `validate:"required,min=8,max=50"`
}

type DomainRepo interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email, password string) (Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
}
