package usecase

import (
	"errors"
	"strings"

	"github.com/paulnune/goexpert-weather/orchestrator-api/internal/entity"
)

type WeatherInputDTO struct {
	Localidade string
	ApiKey     string
}

type WeatherOutputDTO struct {
	City       string  `json:"city"`
	Celcius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

type GetWeatherUseCase struct {
	WeatherRepository entity.WeatherRepositoryInterface
}

func NewGetWeatherUseCase(weather_repository entity.WeatherRepositoryInterface) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		WeatherRepository: weather_repository,
	}
}

func (w *GetWeatherUseCase) Execute(input WeatherInputDTO) (WeatherOutputDTO, error) {
	if input.Localidade == "" {
		return WeatherOutputDTO{}, errors.New("missing input: Localidade")
	}

	if input.ApiKey == "" {
		return WeatherOutputDTO{}, errors.New("missing input: ApiKey")
	}

	weather_resp, err := w.WeatherRepository.Get(input.Localidade, input.ApiKey)
	if err != nil || strings.Contains(string(weather_resp), "city not found") {
		return WeatherOutputDTO{}, errors.New("fail to get weather")
	}

	weather_response, err := w.WeatherRepository.ConvertToWeatherResponse(weather_resp)
	if err != nil {
		return WeatherOutputDTO{}, errors.New("fail to convert weather response: %s" + err.Error())
	}

	weather_dto, err := w.WeatherRepository.ConvertToWeather(weather_response)
	if err != nil {
		return WeatherOutputDTO{}, errors.New("fail to convert weather")
	}

	dto := WeatherOutputDTO{
		City:       input.Localidade,
		Celcius:    weather_dto.Celcius,
		Fahrenheit: weather_dto.Fahrenheit,
		Kelvin:     weather_dto.Kelvin,
	}

	return dto, nil
}
