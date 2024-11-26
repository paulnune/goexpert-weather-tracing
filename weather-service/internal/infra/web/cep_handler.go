package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/paulnune/goexpert-weather-tracing/configs"
	"github.com/paulnune/goexpert-weather-tracing/weather-service/internal/entity"
	"github.com/paulnune/goexpert-weather-tracing/weather-service/internal/infra/repo"
	"github.com/paulnune/goexpert-weather-tracing/weather-service/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type WebCEPHandler struct {
	CEPRepository     entity.CEPRepositoryInterface
	WeatherRepository entity.WeatherRepositoryInterface
	Configs           *configs.Conf
	Tracer            trace.Tracer
}

func NewWebCEPHandler(conf *configs.Conf, tracer trace.Tracer) *WebCEPHandler {
	return &WebCEPHandler{
		CEPRepository:     repo.NewCEPRepository(),
		WeatherRepository: repo.NewWeatherRepository(&http.Client{}),
		Configs:           conf,
		Tracer:            tracer,
	}
}

func NewWebCEPHandlerWithDeps(cepRepo entity.CEPRepositoryInterface, weatherRepo entity.WeatherRepositoryInterface, configs *configs.Conf) *WebCEPHandler {
	return &WebCEPHandler{
		CEPRepository:     cepRepo,
		WeatherRepository: weatherRepo,
		Configs:           configs,
	}
}

func (h *WebCEPHandler) Get(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	_, span := h.Tracer.Start(ctx, "get-city-name")

	cep_address := chi.URLParam(r, "cep")
	open_weathermap_api_key := h.Configs.OpenWeathermapApiKey

	// CEP FLOW
	validate_cep_dto := usecase.ValidateCEPInputDTO{
		CEP: cep_address,
	}

	validateCEP := usecase.NewValidateCEPUseCase(h.CEPRepository)
	is_valid := validateCEP.Execute(validate_cep_dto)
	if !is_valid {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	get_cep_dto := usecase.CEPInputDTO{
		CEP: cep_address,
	}

	getCEP := usecase.NewGetCEPUseCase(h.CEPRepository)
	cep_output, err := getCEP.Execute(get_cep_dto)

	if err != nil {
		http.Error(w, fmt.Sprintf("error getting cep: %v", err), http.StatusInternalServerError)
		return
	}

	if cep_output.Localidade == "" {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}
	span.End()

	// WEATHER FLOW
	_, span = h.Tracer.Start(ctx, "get-city-temp")
	weather_dto := usecase.WeatherInputDTO{
		Localidade: cep_output.Localidade,
		ApiKey:     open_weathermap_api_key,
	}
	getWeather := usecase.NewGetWeatherUseCase(h.WeatherRepository)
	weather_output, err := getWeather.Execute(weather_dto)
	if err != nil || (weather_output.Celcius == 0 && weather_output.Fahrenheit == 0 && weather_output.Kelvin == 0) {
		http.Error(w, fmt.Sprintf("error getting weather: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(weather_output)
	if err != nil {
		http.Error(w, fmt.Sprintf("fail to convert the response to json: %v", err.Error()), http.StatusInternalServerError)
		return
	}
	span.End()
}
