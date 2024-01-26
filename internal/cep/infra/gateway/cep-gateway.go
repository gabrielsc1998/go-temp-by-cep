package cep_gateway

import (
	address "github.com/gabrielsc1998/go-temp-by-cep/internal/cep/domain"
)

type CepGateway interface {
	GetAddressByCep(cep string) (*address.Address, error)
}
