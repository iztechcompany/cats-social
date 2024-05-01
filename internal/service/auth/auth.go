package auth

import (
	"cats-social/internal/repository/user"
	"cats-social/model"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, request *model.LoginRequest) (*model.AuthResponse, error)
	Register(ctx context.Context, request *model.RegisterRequest) (*model.AuthResponse, error)
}

type authService struct {
	userRepository user.UserRepository
}

func NewAuthService(userRepository user.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}
