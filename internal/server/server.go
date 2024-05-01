package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"

	"cats-social/internal/database"
	authH "cats-social/internal/handler/auth"
	userR "cats-social/internal/repository/user"
	authS "cats-social/internal/service/auth"
)

type Server struct {
	port int

	db          database.Service
	authhandler authH.AuthHandler
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.New()
	validate := validator.New()

	userRepository := userR.NewUserPostgresRepository(db.GetDB())
	authService := authS.NewAuthService(userRepository)
	authHandler := authH.NewAuthHandler(authService, validate)

	NewServer := &Server{
		port: port,

		authhandler: authHandler,
		db:          db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
