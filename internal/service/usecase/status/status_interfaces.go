package usecase_status

import (
	"github.com/lallison/h_skills_project/internal/entities"
	"log/slog"
)

//go:generate mockgen -destination=mocks/status_repository.go -package=mocks . StatusRepository

type StatusRepository interface {
	Status(logger *slog.Logger) (*entities.Response, error)
}
