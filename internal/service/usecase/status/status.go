package usecase_status

import (
	"github.com/lallison/h_skills_project/internal/entities"
	"github.com/lallison/h_skills_project/internal/service/repository/status"
)

type UseCaseStatus struct {
	repository *repository_status.RepositoryStatus
}

func New(repository *repository_status.RepositoryStatus) *UseCaseStatus {
	return &UseCaseStatus{repository: repository}
}

func (u *UseCaseStatus) GetStatus() (*entities.Response, error) {
	return u.repository.GetStatus()
}
