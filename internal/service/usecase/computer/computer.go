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

func (u *ComputerUseCase) ComputerByID(id string) (*entities.Computer, error) {
	return u.computerRepository.ComputerByID(id)
}

func (u *ComputerUseCase) CreateComputer(computer *entities.Computer) (*entities.Computer, error) {
	return u.computerRepository.CreateComputer(computer)
}
