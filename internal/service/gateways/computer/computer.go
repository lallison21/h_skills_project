package computer_gateways

import (
	"encoding/json"
	"net/http"
)

type ComputerGateway struct {
	computerUseCase ComputerUseCase
}

func New(computerUseCase ComputerUseCase) *ComputerGateway {
	u := &ComputerGateway{computerUseCase: computerUseCase}
	return u
}

func (g *ComputerGateway) CreateComputer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (g *ComputerGateway) Computers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		computer, err := g.computerUseCase.Computers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(computer)
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

func (g *ComputerGateway) ComputerByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.FormValue("id")
	}
}
