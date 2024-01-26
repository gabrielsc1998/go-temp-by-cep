package weatherapi_presenter

import (
	"encoding/json"

	temp "github.com/gabrielsc1998/go-temp-by-cep/internal/temperature/domain"
)

type WeatherAPIPresenter struct{}

type WeatherAPICurrent struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
}

type WeatherAPIResponse struct {
	Current WeatherAPICurrent `json:"current"`
}

func NewWeatherAPIPresenter() *WeatherAPIPresenter {
	return &WeatherAPIPresenter{}
}

func (v *WeatherAPIPresenter) Present(data map[string]interface{}) (*temp.Temperature, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	weatherAPIResponse := WeatherAPIResponse{}
	err = json.Unmarshal(jsonData, &weatherAPIResponse)
	if err != nil {
		return nil, err
	}
	return temp.New(weatherAPIResponse.Current.TempC, weatherAPIResponse.Current.TempF), nil
}
