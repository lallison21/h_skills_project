package computer_repository

import (
	"database/sql"
	"github.com/lallison/h_skills_project/internal/entities"
)

type ComputerRepository struct {
	db *sql.DB
}

func New(
	db *sql.DB,
) *ComputerRepository {
	repo := &ComputerRepository{db: db}
	return repo
}

func (r *ComputerRepository) Computer() (*entities.Computer, error) {
	return nil, nil
}
