package temp_gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"unicode"

	temp "github.com/gabrielsc1998/go-temp-by-cep/internal/temperature/domain"
	weatherapi_presenter "github.com/gabrielsc1998/go-temp-by-cep/internal/temperature/infra/presenter"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type WeatherAPIGateway struct {
	key string
}

func NewWeatherAPIGateway(key string) *WeatherAPIGateway {
	return &WeatherAPIGateway{key: key}
}

func (w *WeatherAPIGateway) GetTemperatureByCity(city string) (*temp.Temperature, error) {
	adjustedCityName, err := w.adjustCityNameForQuery(city)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(w.url(adjustedCityName))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error")
	}

	data := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	weatherApiPresenter := weatherapi_presenter.NewWeatherAPIPresenter()
	return weatherApiPresenter.Present(data)
}

func (w *WeatherAPIGateway) url(city string) string {
	return fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", w.key, city)
}

func (w *WeatherAPIGateway) adjustCityNameForQuery(city string) (string, error) {
	// Replaces accents and special characters
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	city, _, err := transform.String(t, city)
	if err != nil {
		return "", err
	}

	// Replaces spaces with dashes
	city = strings.ToLower(city)
	splittedCity := strings.Split(city, " ")
	return strings.Join(splittedCity, "-"), nil
}
