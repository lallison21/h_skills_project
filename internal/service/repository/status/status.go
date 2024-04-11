package repository_status

import (
	"database/sql"
	"github.com/lallison/h_skills_project/internal/entities"
	"log/slog"
	"time"
)

type RepositoryStatus struct {
	db *sql.DB
}

func New(db *sql.DB) *RepositoryStatus {
	return &RepositoryStatus{
		db: db,
	}
}

func (r *RepositoryStatus) Status(logger *slog.Logger) (*entities.Response, error) {
	res := &entities.Response{
		ServiceName:    "h_skills_project",
		ServiceVersion: "0.0.1",
		Status:         "OK",
		Timestamp:      time.Now(),
	}

	return res, nil
}
