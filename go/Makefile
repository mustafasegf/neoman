install:
	go mod tidy

run:
	go run cmd/main.go

build:
	go build -o ./bin/main cmd/main.go

run-build:
	./bin/main

test:
	go test -cover -v ./...

.PHONY: build
