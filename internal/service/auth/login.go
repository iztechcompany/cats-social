package auth

import (
	"cats-social/helper"
	"cats-social/model"
	"context"
	"errors"
)

func (s *authService) Login(ctx context.Context, request *model.LoginRequest) (*model.AuthResponse, error) {
	user, err := s.userRepository.GetByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if helper.CheckPassword(user.Password, request.Password) {
		return nil, errors.New("invalid credentials")
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
