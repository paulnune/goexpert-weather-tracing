package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/paulnune/goexpert-weather/orchestrator-api/internal/entity"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WeatherRepository struct {
	client HTTPClient
}

func NewWeatherRepository(client HTTPClient) *WeatherRepository {
	return &WeatherRepository{
		client: client,
	}
}

func (w *WeatherRepository) Get(city string, api_key string) ([]byte, error) {
	city = strings.Replace(city, " ", "%20", -1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf(
			"http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&temperature.unit=celsius",
			city,
			api_key,
		),
		nil,
	)

	if err != nil {
		log.Printf("Fail to create the request: %v", err)
		return nil, err
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Fail to make the request: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	ctx_err := ctx.Err()
	if ctx_err != nil {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			log.Printf("Max timeout reached: %v", err)
			return nil, err
		}
	}

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Fail to read the response: %v", err)
		return nil, err
	}

	if strings.Contains(string(resp), "Invalid API key") {
		return nil, err
	}

	return resp, nil
}

func (w *WeatherRepository) ConvertToWeatherResponse(weather_response []byte) (*entity.WeatherResponse, error) {
	var weather_res entity.WeatherResponse
	err := json.Unmarshal(weather_response, &weather_res)
	if err != nil {
		log.Printf("Fail to decode the response: %v", err)
		return nil, err
	}

	return &weather_res, nil
}

func (w *WeatherRepository) ConvertToWeather(weather_response *entity.WeatherResponse) (*entity.Weather, error) {
	weather := entity.Weather{}
	weather.MakeTemperatureConversions(weather_response.Main.Temp)
	return &weather, nil
}
