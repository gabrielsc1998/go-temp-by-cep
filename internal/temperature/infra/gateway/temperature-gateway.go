package temp_gateway

import temp "github.com/gabrielsc1998/go-temp-by-cep/internal/temperature/domain"

type TemperatureGateway interface {
	GetTemperatureByCity(city string) (*temp.Temperature, error)
}
