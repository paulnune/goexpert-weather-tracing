
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

### Subindo os serviços
Utilize o comando a seguir para subir toda a infraestrutura necessária:

```bash
make infra-up
```

### Chamando os serviços
- Para enviar uma consulta ao `cep-service`:
  ```bash
  make svc-a
  ```
- Para consultar o `weather-service` com base no CEP:
  ```bash
  make svc-b
  ```

### Derrubando a infraestrutura
Para parar e remover os containers criados, use:
```bash
make infra-down
```

### Limpando recursos Docker/Podman
Para remover containers, imagens e volumes não utilizados, execute:
```bash
make docker-clean-up
```

## Evidências 📷

### Imagem 1: Tela do Zipkin exibindo os traces do `service-b`
![Imagem 1](1.png)
Esta imagem mostra a visualização de um trace no `service-b`, com spans detalhados para identificar tempos de execução das chamadas.

### Imagem 2: Detalhamento de spans do `service-b`
![Imagem 2](2.png)
Esta imagem apresenta o detalhamento dos spans internos do `service-b`, incluindo o início (`service-b-start`) e chamadas específicas como `get-location` e `get-weather`.

### Imagem 3: Trace do `service-a` chamando o `service-b`
![Imagem 3](3.png)
Nesta imagem, vemos o trace do `service-a` ao realizar uma chamada para o `service-b`, exibindo um único span com tempo total.

### Imagem 4: Listagem de traces no Zipkin
![Imagem 4](4.png)
Aqui está a visão geral de vários traces no Zipkin, exibindo a duração e os serviços envolvidos em cada trace.

## Estrutura do Projeto 📂

```
.
├── docker-compose.yaml
├── Makefile
├── README.md
├── service_a
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── service_b
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    └── main.go
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