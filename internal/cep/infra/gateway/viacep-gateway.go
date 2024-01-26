package cep_gateway

import (
	"encoding/json"
	"errors"
	"net/http"

	address "github.com/gabrielsc1998/go-temp-by-cep/internal/cep/domain"
	viacep_presenter "github.com/gabrielsc1998/go-temp-by-cep/internal/cep/infra/presenter"
)

type ViaCepGateway struct{}

func NewViaCepGateway() *ViaCepGateway {
	return &ViaCepGateway{}
}

func (v *ViaCepGateway) GetAddressByCep(cep string) (*address.Address, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data["erro"] != nil && data["erro"].(bool) {
		return nil, errors.New("cep not found")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error")
	}

	viaCepPresenter := viacep_presenter.NewViaCepPresenter()
	return viaCepPresenter.Present(data)
}
