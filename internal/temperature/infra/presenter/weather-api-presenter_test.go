package weatherapi_presenter

import (
	"encoding/json"
	"reflect"
	"testing"

	temp "github.com/gabrielsc1998/go-temp-by-cep/internal/temperature/domain"
)

func TestPresent(t *testing.T) {
	presenter := NewWeatherAPIPresenter()

	data := map[string]interface{}{
		"current": map[string]interface{}{
			"temp_c": 20.0,
			"temp_f": 68.0,
		},
	}

	expected := temp.New(20.0, 68.0)

	jsonData, _ := json.Marshal(data)
	weatherAPIResponse := WeatherAPIResponse{}
	_ = json.Unmarshal(jsonData, &weatherAPIResponse)

	result, err := presenter.Present(data)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}
