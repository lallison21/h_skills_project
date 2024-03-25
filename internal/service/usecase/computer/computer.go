package computer

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

func (u *ComputerUseCase) GetStatus() (*entities.Computer, error) {
	return u.computerRepository.Computer()
}
