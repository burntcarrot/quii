package user

import "context"

type Domain struct {
	ID       uint
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=50"`
	Role     string `validate:"required"`
}

type LoginDomain struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=50"`
}

type DomainRepo interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email, password string) (Domain, error)
}
