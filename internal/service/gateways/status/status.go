package gateway_status

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type GatewayStatus struct {
	UseCase StatusUseCase
}

func New(useCase StatusUseCase) *GatewayStatus {
	return &GatewayStatus{UseCase: useCase}
}

func (g *GatewayStatus) Status(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := g.UseCase.Status(logger)
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
}
