build:
	@go build -o bin/gobank

run:
	@air

test:
	@go test -v ./...
