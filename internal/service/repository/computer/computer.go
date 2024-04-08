package computer_repository

import (
	"database/sql"
	"fmt"
	"github.com/lallison/h_skills_project/internal/entities"
	"log"
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

func (r *ComputerRepository) Computers() ([]*entities.Computer, error) {
	var res []*entities.Computer

	stmt, err := r.db.Prepare("SELECT * FROM computers")
	if err != nil {
		return res, fmt.Errorf("failed to prepare query: %w", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return res, fmt.Errorf("failed to query rows: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	for rows.Next() {
		var c *entities.Computer

		if err := rows.Scan(c); err != nil {
			return res, fmt.Errorf("failed to scan row: %w", err)
		}
		res = append(res, c)
	}
	if err := rows.Err(); err != nil {
		return res, fmt.Errorf("failed to scan rows: %w", err)
	}

	return res, nil
}
