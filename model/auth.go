package model

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,max=20"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=20"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type AuthResponse struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
