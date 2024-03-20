package repository_status

import (
	"github.com/lallison/h_skills_project/internal/entities"
	"github.com/lallison/h_skills_project/version"
	"time"
)

type RepositoryStatus struct {
	response *entities.Response
}

func New() *RepositoryStatus {
	return &RepositoryStatus{
		response: &entities.Response{
			ServiceName:    "h_skills_project",
			ServiceVersion: version.SERVICE_VERSION,
			Timestamp:      time.Now(),
			Status:         "OK",
		},
	}
}

func (repository *RepositoryStatus) GetStatus() (*entities.Response, error) {
	return repository.response, nil
}
