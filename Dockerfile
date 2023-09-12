FROM golang:latest
WORKDIR /app
COPY . .
EXPOSE 8080
CMD go run cmd/start/main.go