package user

import (
	"cats-social/internal/database"
	"cats-social/model"
	"context"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	service := database.New()
	repo := NewUserPostgresRepository(service.GetDB())

	user1 := &model.User{
		Name:  "Umar Bawazir",
		Email: "umarbawazirfaiz@gmail.com",
	}

	user1.HashPassword("umar123", 14)

	err := repo.Create(ctx, user1)
	if err != nil {
		t.Errorf("failed to create user: %v", err)
		return
	}
}
