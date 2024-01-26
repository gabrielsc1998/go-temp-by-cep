package cep_presenter

import (
	"reflect"
	"testing"

	address "github.com/gabrielsc1998/go-temp-by-cep/internal/cep/domain"
)

func TestPresent(t *testing.T) {
	presenter := NewViaCepPresenter()

	data := map[string]interface{}{
		"cep":        "12345-678",
		"uf":         "SP",
		"localidade": "Sao Paulo",
		"bairro":     "Centro",
		"logradouro": "Rua XYZ",
	}

	expected := &address.Address{
		Code:     "12345-678",
		State:    "SP",
		City:     "Sao Paulo",
		District: "Centro",
		Address:  "Rua XYZ",
	}

	result, err := presenter.Present(data)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}
