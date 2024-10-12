build:
	@go build -o bin/api cmd/main.go

run: build
	@./bin/api