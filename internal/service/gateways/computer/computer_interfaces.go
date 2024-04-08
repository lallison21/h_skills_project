package computer_gateways

import "github.com/lallison/h_skills_project/internal/entities"

type ComputerUseCase interface {
	Computers() ([]*entities.Computer, error)
}
