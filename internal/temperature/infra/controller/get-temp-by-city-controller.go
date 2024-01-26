package temp_controller

import (
	"fmt"
	"net/http"

	cep_gateway "github.com/gabrielsc1998/go-temp-by-cep/internal/cep/infra/gateway"
	cep_validators "github.com/gabrielsc1998/go-temp-by-cep/internal/cep/validators"
	temp_gateway "github.com/gabrielsc1998/go-temp-by-cep/internal/temperature/infra/gateway"
	"github.com/go-chi/chi/v5"
)

type GetTempByCityController struct {
	tempGateway temp_gateway.TemperatureGateway
	cepGateway  cep_gateway.CepGateway
}

func NewGetTempByCityController(tempGtw temp_gateway.TemperatureGateway, cepGtw cep_gateway.CepGateway) *GetTempByCityController {
	return &GetTempByCityController{
		tempGateway: tempGtw,
		cepGateway:  cepGtw,
	}
}

func (g *GetTempByCityController) Handle(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if !cep_validators.IsValidCep(cep) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	address, err := g.cepGateway.GetAddressByCep(cep)
	if err != nil {
		if err.Error() == "cep not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("can not found zipcode"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	temp, err := g.tempGateway.GetTemperatureByCity(address.City)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(fmt.Sprintf(`{"temp_C": %.1f, "temp_F": %.1f, "temp_K": %.1f}`, temp.TempC, temp.TempF, temp.TempK)))
}
