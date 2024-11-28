FROM golang:1.22.3 as build-b
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goapp

FROM alpine:latest
WORKDIR /app
COPY --from=build-b /app/goapp .
EXPOSE 8080
CMD ["./goapp"]