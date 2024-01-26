package main

import (
	"github.com/gabrielsc1998/go-temp-by-cep/configs"
	viacep_gateway "github.com/gabrielsc1998/go-temp-by-cep/internal/cep/infra/gateway"
	"github.com/gabrielsc1998/go-temp-by-cep/internal/common/webserver"
	temp_controller "github.com/gabrielsc1998/go-temp-by-cep/internal/temperature/infra/controller"
	weatherapi_gateway "github.com/gabrielsc1998/go-temp-by-cep/internal/temperature/infra/gateway"
)

func main() {
	config, err := configs.LoadConfig("xsxnamxa")
	if err != nil {
		panic(err)
	}

	viacepGateway := viacep_gateway.NewViaCepGateway()
	weatherApiGateway := weatherapi_gateway.NewWeatherAPIGateway(config.WeatherApiKey)

	getTempByCityController := temp_controller.NewGetTempByCityController(weatherApiGateway, viacepGateway)

	webserver := webserver.NewWebServer(config.WebServerPort)
	webserver.Get("/temp/city/{cep}", getTempByCityController.Handle)

	webserver.Start()
}
