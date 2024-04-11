package usecase_status

import (
	"github.com/lallison/h_skills_project/internal/entities"
	"log/slog"
)

type UseCaseStatus struct {
	repository StatusRepository
}

func New(repository StatusRepository) *UseCaseStatus {
	return &UseCaseStatus{repository: repository}
}

func (u *UseCaseStatus) Status(logger *slog.Logger) (*entities.Response, error) {
	return u.repository.Status(logger)
}
