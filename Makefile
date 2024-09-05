run: build
	@./bin/api

build:
	@echo "Executing go build"
	go build -v -buildmode=pie -ldflags "-X main.version=0.1.0"
	@echo "Binary ready"


test:
	@go test -v ./...