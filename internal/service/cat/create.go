package cat

import (
	"cats-social/model"
	"context"
)

func (s *catService) Create(ctx context.Context, request *model.CatRequest) (*model.CatCreateResponse, error) {
	cat := &model.Cat{
		Name:        request.Name,
		Race:        request.Race,
		Sex:         request.Sex,
		AgeInMonth:  request.AgeInMonth,
		Description: request.Description,
	}

	err := s.catRepository.Create(ctx, cat)
	if err != nil {
		return nil, err
	}

	catResponse := &model.CatCreateResponse{
		ID:        cat.ID,
		CreatedAt: cat.CreatedAt,
	}

	return catResponse, nil
}
