package user

import (
	"context"
	"database/sql"
	"time"
)

type Repository interface {
	Insert(ctx context.Context, user *Model) error
	GetByID(ctx context.Context, id int) (*Model, error)
}

type repository struct {
	db *sql.DB
}

func (r *repository) Insert(ctx context.Context, user *Model) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, "INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).
		Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByID(ctx context.Context, id int) (*Model, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var user Model
	err := r.db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
