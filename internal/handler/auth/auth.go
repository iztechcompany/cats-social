package auth

import (
	"cats-social/internal/service/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authHandler struct {
	authService auth.AuthService
	validate    *validator.Validate
}

func NewAuthHandler(authService auth.AuthService, validate *validator.Validate) AuthHandler {
	return &authHandler{
		authService: authService,
		validate:    validate,
	}
}
