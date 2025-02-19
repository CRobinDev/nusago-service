include .env

build:
	@echo "Build app..."
	@go build -o ./bin/nusago-service cmd/app/main.go

run: build
	@echo "Starting app..."
	@./bin/nusago-service

