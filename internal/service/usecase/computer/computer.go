package computer_usecase

import (
	"github.com/lallison/h_skills_project/internal/entities"
	"log/slog"
)

type ComputerUseCase struct {
	computerRepository ComputerRepository
}

func New(
	computerRepository ComputerRepository,
) *ComputerUseCase {
	uc := &ComputerUseCase{
		computerRepository: computerRepository,
	}

	return uc
}

func (u *ComputerUseCase) Computers(logger *slog.Logger) ([]*entities.Computer, error) {
	return u.computerRepository.Computers(logger)
}

func (u *ComputerUseCase) ComputerByID(logger *slog.Logger, id string) (*entities.Computer, error) {
	return u.computerRepository.ComputerByID(logger, id)
}

func (u *ComputerUseCase) CreateComputer(logger *slog.Logger, computer *entities.Computer) error {
	return u.computerRepository.CreateComputer(logger, computer)
}
