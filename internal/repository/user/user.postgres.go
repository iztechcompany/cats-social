package user

import (
	"cats-social/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type userPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) UserRepository {
	return &userPostgresRepository{
		db: db,
	}
}

func (r *userPostgresRepository) GetAll(limit, offset, query string) (*[]model.User, error) {
	return nil, nil
}

func (r *userPostgresRepository) GetById(id int) (*model.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE id = $1"

	row := r.db.QueryRow(query, id)

	var user model.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		// Periksa jika tidak ada baris yang ditemukan atau ada kesalahan lainnya
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with ID %f not found", id)
		}
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	return &user, nil
}

func (r *userPostgresRepository) GetByEmail(email string) (*model.User, error) {
	query := "SELECT id, name, email, password, created_at FROM users WHERE email = $1"

	row := r.db.QueryRow(query, email)

	var user model.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		// Periksa jika tidak ada baris yang ditemukan atau ada kesalahan lainnya
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	return &user, nil
}

func (r *userPostgresRepository) Create(ctx context.Context, request *model.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`

	// Execute the query
	_, err := r.db.Exec(query, request.Name, request.Email, request.Password)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (r *userPostgresRepository) Update(id string, request *model.User) error {
	return nil
}

func (r *userPostgresRepository) Delete(id string) error {
	return nil
}
