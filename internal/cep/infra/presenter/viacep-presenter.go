package cep_presenter

import (
	"encoding/json"

	address "github.com/gabrielsc1998/go-temp-by-cep/internal/cep/domain"
)

type ViaCepPresenter struct {
}

type ViaCepResponse struct {
	Code     string `json:"cep"`
	State    string `json:"uf"`
	City     string `json:"localidade"`
	District string `json:"bairro"`
	Address  string `json:"logradouro"`
}

func NewViaCepPresenter() *ViaCepPresenter {
	return &ViaCepPresenter{}
}

func (v *ViaCepPresenter) Present(data map[string]interface{}) (*address.Address, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	viaCepResponse := ViaCepResponse{}
	err = json.Unmarshal(jsonData, &viaCepResponse)
	if err != nil {
		return nil, err
	}
	address := address.New(
		viaCepResponse.Code,
		viaCepResponse.State,
		viaCepResponse.City,
		viaCepResponse.District,
		viaCepResponse.Address,
	)
	return address, nil
}
