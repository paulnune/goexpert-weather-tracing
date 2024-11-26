package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Conf struct {
	InputApiHttpPort                      string `mapstructure:"INPUT_API_HTTP_PORT"`
	InputApiOtelServiceName               string `mapstructure:"INPUT_API_OTEL_SERVICE_NAME"`
	OrchestratorApiPort                   string `mapstructure:"ORCHESTRATOR_API_PORT"`
	OrchestratorApiHost                   string `mapstructure:"ORCHESTRATOR_API_HOST"`
	OpenWeathermapApiKey                  string `mapstructure:"OPEN_WEATHERMAP_API_KEY"`
	OrchestratorApiServiceName            string `mapstructure:"ORCHESTRATOR_API_SERVICE_NAME"`
	OpenTelemetryCollectorExporerEndpoint string `mapstructure:"OPEN_TELEMETRY_COLLECTOR_EXPORTER_ENDPOINT"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	var err error

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	viper.BindEnv("INPUT_API_HTTP_PORT")
	viper.BindEnv("INPUT_API_OTEL_SERVICE_NAME")
	viper.BindEnv("ORCHESTRATOR_API_PORT")
	viper.BindEnv("ORCHESTRATOR_API_HOST")
	viper.BindEnv("OPEN_WEATHERMAP_API_KEY")
	viper.BindEnv("ORCHESTRATOR_API_SERVICE_NAME")
	viper.BindEnv("OPEN_TELEMETRY_COLLECTOR_EXPORTER_ENDPOINT")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("WARNING: %v\n", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, err
}
