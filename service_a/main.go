package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type (
	DTOInput struct {
		Cep string `json:"cep"`
	}
	DTOOutput struct {
		TempC float64 `json:"temp_C"`
		TempF float64 `json:"temp_F"`
		TempK float64 `json:"temp_K"`
		City  string  `json:"city"`
	}
)

func main() {
	setTracing()
	http.HandleFunc("POST /", Handle)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var input DTOInput
	var output DTOOutput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !validCep(input.Cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	ctx, span := otel.Tracer("service-a").Start(r.Context(), "1 - req-to-service-b")
	defer span.End()

	output, status, err := getInfo(input.Cep, ctx)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func setTracing() {
	exporter, err := zipkin.New("http://zipkin:9411/api/v2/spans")
	if err != nil {
		log.Fatalf("Fail to create Zipkin exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("service-a"),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
}

func getInfo(cep string, ctx context.Context) (DTOOutput, int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://goapp-service-b:8081/"+cep, nil)
	if err != nil {
		return DTOOutput{}, http.StatusInternalServerError, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return DTOOutput{}, resp.StatusCode, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return DTOOutput{}, resp.StatusCode, errors.New("can not find zipcode")
	}

	var output DTOOutput
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return output, resp.StatusCode, err
	}

	return output, resp.StatusCode, err
}

func validCep(cep string) bool {
	if len(cep) != 8 {
		return false
	}
	for _, c := range cep {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
