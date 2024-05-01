package user

import (
	"cats-social/model"
	"context"
)

type UserRepository interface {
	GetAll(limit, offset, query string) (*[]model.User, error)
	GetById(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(ctx context.Context, request *model.User) error
	Update(id string, request *model.User) error
	Delete(id string) error
}
