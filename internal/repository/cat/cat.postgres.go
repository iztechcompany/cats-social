package cat

import (
	"cats-social/model"
	"context"
	"database/sql"
	"log"
	"time"
)

type catPostgresRepository struct {
	db *sql.DB
}

func NewCatPostgresRepository(db *sql.DB) CatRepository {
	return &catPostgresRepository{
		db: db,
	}
}

func (r *catPostgresRepository) GetAll(ctx context.Context, limit, offset, search string) ([]*model.Cat, error) {
	start := time.Now()

	query := "SELECT id, name, race, sex, age_in_month, description, has_matched, created_at FROM cats"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []*model.Cat
	for rows.Next() {
		var cat model.Cat
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.HasMatched, &cat.CreatedAt); err != nil {
			return nil, err
		}
		cats = append(cats, &cat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	duration := time.Since(start)
	log.Printf("Create Execution time: %s", duration)

	return cats, nil
}

func (r *catPostgresRepository) Create(ctx context.Context, request *model.Cat) error {
	start := time.Now()

	query := "INSERT INTO cats (name, race, sex, age_in_month, description, has_matched) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at"

	err := r.db.QueryRowContext(ctx, query, request.Name, request.Race, request.Sex, request.AgeInMonth, request.Description, request.HasMatched).Scan(&request.ID, &request.CreatedAt)
	if err != nil {
		return err
	}

	duration := time.Since(start)
	log.Printf("Create Execution time: %s", duration)

	return nil
}

func (r *catPostgresRepository) Update(ctx context.Context, id string, request *model.Cat) error {
	return nil
}

func (r *catPostgresRepository) Delete(ctx context.Context, id string) error {
	return nil
}
