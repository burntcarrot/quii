package user

import "context"

type Domain struct {
	ID       uint
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=50"`
	Role     string `validate:"required"`
}

type DomainRepo interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
}
