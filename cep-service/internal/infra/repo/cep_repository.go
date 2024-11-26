package repo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type CEPRepository struct {
	OrchestratorApiHost string
	OrchestratorApiPort string
}

func NewCEPRepository(orchestrator_api_host string, orchestrator_api_port string) *CEPRepository {
	return &CEPRepository{
		OrchestratorApiHost: orchestrator_api_host,
		OrchestratorApiPort: orchestrator_api_port,
	}
}

func (r *CEPRepository) IsValid(cep_address string) bool {
	check, _ := regexp.MatchString("^[0-9]{8}$", cep_address)
	return (len(cep_address) == 8 && cep_address != "" && check)
}

func (r *CEPRepository) Get(cep_address string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(
		"http://%s:%s/cep/%s",
		r.OrchestratorApiHost,
		r.OrchestratorApiPort,
		cep_address),
		nil,
	)
	if err != nil {
		log.Printf("Fail to create the request: %v", err)
		return err
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport,
			otelhttp.WithSpanNameFormatter(func(_ string, req *http.Request) string {
				return "get-cep-temp"
			}),
		),
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Fail to make the request: %v", err)
		return err
	}
	defer resp.Body.Close()

	ctx_err := ctx.Err()
	if ctx_err != nil {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			log.Printf("Max timeout reached: %v", err)
			return err
		}
	}

	return nil
}
