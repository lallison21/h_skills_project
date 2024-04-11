package gateway_status

import (
	"github.com/lallison/h_skills_project/internal/entities"
	"log/slog"
)

//go:generate mockgen -destination=mocks/status_usecase.go -package=mocks . StatusUseCase

type StatusUseCase interface {
	Status(logger *slog.Logger) (*entities.Response, error)
}
