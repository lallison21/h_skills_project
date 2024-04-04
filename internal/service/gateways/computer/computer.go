package computer_gateways

import "github.com/lallison/h_skills_project/internal/entities"

type ComputerGateway struct {
	computerUseCase ComputerUseCase
}

func New(computerUseCase ComputerUseCase) *ComputerGateway {
	u := &ComputerGateway{computerUseCase: computerUseCase}
	return u
}

func (g *ComputerGateway) Computer() (*entities.Computer, error) {
	return nil, nil
}
