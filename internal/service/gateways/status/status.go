package gateway_status

import (
	"encoding/json"
	"github.com/lallison/h_skills_project/internal/entities"
	usecase_status "github.com/lallison/h_skills_project/internal/service/usecase/status"
	"net/http"
)

type GatewayStatus struct {
	UseCase *usecase_status.UseCaseStatus
}

func New(useCase *usecase_status.UseCaseStatus) *GatewayStatus {
	return &GatewayStatus{UseCase: useCase}
}

func (g *GatewayStatus) StatusHandle(w http.ResponseWriter, r *http.Request) {
	status, err := g.GetStatus()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (g *GatewayStatus) GetStatus() (*entities.Response, error) {
	return g.UseCase.GetStatus()
}
