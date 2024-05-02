package cat

import (
	"cats-social/model"
	"context"
)

type catMockRepository struct {
}

func NewCatMockRepository() CatRepository {
	return &catMockRepository{}
}

func (r *catMockRepository) GetAll(ctx context.Context, limit, offset, search string) ([]*model.Cat, error) {
	return nil, nil
}

func (r *catMockRepository) Create(ctx context.Context, request *model.Cat) error {
	return nil
}

func (r *catMockRepository) Update(ctx context.Context, id string, request *model.Cat) error {
	return nil
}

func (r *catMockRepository) Delete(ctx context.Context, id string) error {
	return nil
}
