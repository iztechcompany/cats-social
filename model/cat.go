package model

import "time"

type Cat struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	AgeInMonth  int64     `json:"ageInMonth"`
	Description string    `json:"description"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type CatParams struct {
	ID     int64 `json:"id"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type CatRequest struct {
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	AgeInMonth  int64    `json:"ageInMonth"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageUrls"`
}

type CatResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	AgeInMonth  int64     `json:"ageInMonth"`
	Description string    `json:"description"`
	HasMatched  bool      `json:"hasMatched"`
	ImageUrls   []string  `json:"imageUrls"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CatCreateResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
