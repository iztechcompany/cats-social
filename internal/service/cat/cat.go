package cat

import (
	"cats-social/internal/repository/cat"
	"cats-social/model"
	"context"
)

type CatService interface {
	GetAll(ctx context.Context, params *model.CatParams) ([]*model.CatResponse, error)
	Create(ctx context.Context, request *model.CatRequest) (*model.CatCreateResponse, error)
}

type catService struct {
	catRepository cat.CatRepository
}

func NewCatService(catRepository cat.CatRepository) CatService {
	return &catService{
		catRepository: catRepository,
	}
}
