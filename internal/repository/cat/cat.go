package cat

import (
	"cats-social/model"
	"context"
)

type CatRepository interface {
	GetAll(ctx context.Context, limit, offset, search string) ([]*model.Cat, error)
	Create(ctx context.Context, request *model.Cat) error
	Update(ctx context.Context, id string, request *model.Cat) error
	Delete(ctx context.Context, id string) error
}
