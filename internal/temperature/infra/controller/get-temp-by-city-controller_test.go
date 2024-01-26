package temp_controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	cep_gateway "github.com/gabrielsc1998/go-temp-by-cep/internal/cep/infra/gateway"
	temp_gateway "github.com/gabrielsc1998/go-temp-by-cep/internal/temperature/infra/gateway"
	"github.com/go-chi/chi/v5"
)

var controller *GetTempByCityController = nil

func init() {
	// ----- Read .env file ----- //
	data, err := os.ReadFile("../../../../.env")
	if err != nil {
		panic(err)
	}
	dataSt := string(data)
	key := strings.Split(strings.Split(dataSt, "\n")[1], "=")[1]
	os.Setenv("WEATHER_API_KEY", key)

	// ----- Setup ----- //

	tempGateway := temp_gateway.NewWeatherAPIGateway(os.Getenv("WEATHER_API_KEY"))
	cepGateway := cep_gateway.NewViaCepGateway()
	controller = NewGetTempByCityController(tempGateway, cepGateway)
}

func TestShouldReturn422WhenTheCepIsInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/temp/city/123", nil)
	if err != nil {
		t.Fatal(err)
	}
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("cep", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	recorder := httptest.NewRecorder()
	controller.Handle(recorder, req)
	if recorder.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected %d, got %d", http.StatusUnprocessableEntity, recorder.Code)
	}
}

func TestShouldReturn404WhenTheCepIsNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/temp/city/88160000", nil)
	if err != nil {
		t.Fatal(err)
	}
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("cep", "88160000")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	recorder := httptest.NewRecorder()
	controller.Handle(recorder, req)
	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected %d, got %d", http.StatusNotFound, recorder.Code)
	}
}

func TestShouldReturn200WhenTheCepIsValid(t *testing.T) {
	req, err := http.NewRequest("GET", "/temp/city/01031970", nil)
	if err != nil {
		t.Fatal(err)
	}
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("cep", "01031970")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	recorder := httptest.NewRecorder()

	controller.Handle(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, recorder.Code)
	}

	if !strings.Contains(recorder.Body.String(), "temp_C") {
		t.Errorf("Expected %s, got %s", "temp_C", recorder.Body.String())
	}
	if !strings.Contains(recorder.Body.String(), "temp_F") {
		t.Errorf("Expected %s, got %s", "temp_F", recorder.Body.String())
	}
	if !strings.Contains(recorder.Body.String(), "temp_K") {
		t.Errorf("Expected %s, got %s", "temp_K", recorder.Body.String())
	}
}
