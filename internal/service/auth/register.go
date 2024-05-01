package auth

import (
	"cats-social/helper"
	"cats-social/model"
	"context"
)

func (s *authService) Register(ctx context.Context, request *model.RegisterRequest) (*model.AuthResponse, error) {
	user := &model.User{
		Name:  request.Name,
		Email: request.Email,
	}

	if err := user.HashPassword(user.Password, 14); err != nil {
		return nil, err
	}

	err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	accessToken, err := helper.NewToken(uint64(user.ID)).Create()
	if err != nil {
		return nil, err
	}

	response := &model.AuthResponse{
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: accessToken,
	}

	return response, nil
}
