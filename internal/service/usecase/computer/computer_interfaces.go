package computer_usecase

import (
	"github.com/lallison/h_skills_project/internal/entities"
	"log/slog"
)

//go:generate mockgen -destination=mocks/computer_repository.go -package=mocks . ComputerRepository
//go:generate mockgen -destination=mocks/computer_queue_producer.go -package=mocks . ComputerQueueProducer
//go:generate mockgen -destination=mocks/cache.go -package=mocks . Cache

type ComputerRepository interface {
	Computers(logger *slog.Logger) ([]*entities.Computer, error)
	ComputerByID(logger *slog.Logger, id string) (*entities.Computer, error)
	CreateComputer(logger *slog.Logger, computer *entities.Computer) error
}

type ComputerQueueProducer interface {
	Success() error
	Error() error
}

type Cache interface {
	Computers() ([]*entities.Computer, error)
}
