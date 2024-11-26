package entity

type CEPRepositoryInterface interface {
	Get(string) ([]byte, error)
	Convert([]byte) (*CEP, error)
	IsValid(string) bool
}

type WeatherRepositoryInterface interface {
	Get(string, string) ([]byte, error)
	ConvertToWeatherResponse([]byte) (*WeatherResponse, error)
	ConvertToWeather(*WeatherResponse) (*Weather, error)
}
