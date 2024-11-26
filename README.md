
# GoExpert Weather com Tracing 🌦️

Projeto desenvolvido em Go para consulta de clima atual com base em um CEP. O sistema retorna a temperatura em graus Celsius, Fahrenheit e Kelvin. Agora, também inclui **Tracing** para observabilidade, integrando ferramentas como OpenTelemetry e Zipkin. Desenvolvido pelo **Paulo Nunes**.

## Funcionalidades 📋

- Receber um CEP válido de 8 dígitos.
- Consultar a API ViaCEP para identificar a localização do CEP.
- Utilizar a API WeatherAPI para consultar a temperatura na localização encontrada.
- Converter e retornar a temperatura nos formatos Celsius, Fahrenheit e Kelvin.
- Tracing distribuído com OpenTelemetry para facilitar a análise de desempenho.

## Requisitos 📦

- Docker ou Podman e Docker ou Podman Compose instalados.
- Configuração do ambiente com as variáveis:
  - `WEATHER_API_KEY`: Chave da API WeatherAPI para consulta de clima.

## Exemplos de uso 🛠️

### Com `curl`

#### `cep-service`

```bash
curl "http://localhost:8081/validate?cep=01001000"
# Saída esperada:
# {"cep":"01001-000","logradouro":"Praça da Sé","localidade":"São Paulo","uf":"SP"}
```

#### `weather-service`

```bash
curl "http://localhost:8082/weather?cep=01001000"
# Saída esperada:
# {"temp_C":22.1,"temp_F":71.78,"temp_K":295.25}
```

## Como executar o projeto 🚀

### Utilizando Docker Compose

1. Suba os containers com:
   ```bash
   docker-compose up --build
   ```

2. Os serviços estarão disponíveis em:
   - `cep-service`: [http://localhost:8081](http://localhost:8081)
   - `weather-service`: [http://localhost:8082](http://localhost:8082)

### Utilizando Podman Compose

1. Suba os containers com:
   ```bash
   podman-compose up --build
   ```

2. Acesse os mesmos endpoints indicados acima.

### Localmente

1. Configure as variáveis de ambiente:
   ```bash
   export WEATHER_API_KEY="sua_chave_api_weather"
   ```

2. Execute cada serviço individualmente:

- Para o `cep-service`:
   ```bash
   cd cep-service
   go mod tidy
   go run cmd/app/main.go
   ```

- Para o `weather-service`:
   ```bash
   cd weather-service
   go run cmd/app/main.go
   ```

## Estrutura do Projeto 📂

```
.
├── cep-service
│   ├── cmd
│   │   └── app
│   │       └── main.go
│   ├── Dockerfile
│   └── internal
│       ├── entity
│       │   ├── cep.go
│       │   └── interface.go
│       ├── infra
│       │   ├── repo
│       │   │   └── cep_repository.go
│       │   └── web
│       │       ├── cep_handler.go
│       │       ├── status_handler.go
│       │       └── webserver
│       │           ├── starter.go
│       │           └── webserver.go
│       └── usecase
│           ├── get_cep.go
│           └── validate_cep.go
├── configs
│   └── config.go
├── docker-compose.yaml
├── go.mod
├── go.sum
├── o11y
│   ├── otel-collector-config.yaml
│   └── prometheus.yaml
├── pkg
│   └── otel
│       └── otel_provider.go
├── README.md
└── weather-service
    ├── cmd
    │   └── app
    │       └── main.go
    ├── Dockerfile
    └── internal
        ├── entity
        │   ├── cep.go
        │   ├── interface.go
        │   └── weather.go
        ├── infra
        │   ├── repo
        │   │   ├── cep_repository.go
        │   │   └── weather_repository.go
        │   └── web
        │       ├── cep_handler.go
        │       ├── status_handler.go
        │       └── webserver
        │           ├── starter.go
        │           └── webserver.go
        └── usecase
            ├── get_cep.go
            ├── get_weather.go
            └── validate_cep.go
```

## Testes automatizados ✅

1. Configure o ambiente:
   ```bash
   go mod tidy
   ```

2. Execute os testes:
   ```bash
   go test ./internal/repository/... ./internal/usecase/... -v
   ```

## 👨‍💻 Autor

**Paulo Henrique Nunes Vanderley**  
- 🌐 [Site Pessoal](https://www.paulonunes.dev/)  
- 🌐 [GitHub](https://github.com/paulnune)  
- ✉️ Email: [paulo.nunes@live.de](mailto:paulo.nunes@live.de)  
- 🚀 Aluno da Pós **GoExpert 2024** pela [FullCycle](https://fullcycle.com.br)

---

## 🎉 Agradecimentos

Este repositório foi desenvolvido com muita dedicação para a **Pós GoExpert 2024**. Agradeço à equipe da **FullCycle** por proporcionar uma experiência incrível de aprendizado! 🚀