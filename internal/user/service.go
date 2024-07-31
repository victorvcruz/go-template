package user

import (
	"context"
)

type Service interface {
	Insert(ctx context.Context, user *Model) error
	GetByID(ctx context.Context, id int) (*Model, error)
}

type service struct {
	repository Repository
}

func (r *service) Insert(ctx context.Context, user *Model) error {
	return r.repository.Insert(ctx, user)
}

func (r *service) GetByID(ctx context.Context, id int) (*Model, error) {
	return r.repository.GetByID(ctx, id)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
