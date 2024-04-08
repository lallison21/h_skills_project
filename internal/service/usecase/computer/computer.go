package computer_usecase

import "github.com/lallison/h_skills_project/internal/entities"

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

func (u *ComputerUseCase) Computers() ([]*entities.Computer, error) {
	return u.computerRepository.Computers()
}
