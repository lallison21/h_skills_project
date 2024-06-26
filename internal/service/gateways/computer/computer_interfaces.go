package computer_gateways

import (
	"github.com/lallison/h_skills_project/internal/entities"
	"log/slog"
)

//go:generate mockgen -destination=mocks/computer_usecase.go -package=mocks . ComputerUseCase

type ComputerUseCase interface {
	Computers(logger *slog.Logger) ([]*entities.Computer, error)
	ComputerByID(logger *slog.Logger, id string) (*entities.Computer, error)
	CreateComputer(logger *slog.Logger, computer *entities.Computer) error
}
