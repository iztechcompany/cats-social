package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	authGroup := r.Group("/auth")
	authGroup.POST("/login", s.authhandler.Login)
	authGroup.POST("/register", s.authhandler.Register)

	catGroup := r.Group("/v1/cat")
	catGroup.POST("/", s.catHandler.Create)

	return r
}
