PROJECT:=nirvana

.PHONY: build
build: wire
	go build -o $(PROJECT) main.go
.PHONY: run
run: wire
	go run main.go http
.PHONY: wire
wire:
	wire gen cmd/http/inject/wire.go
.PHONY: swag
swag:
	@echo "generate gin swagger doc."
	swag init --output=internal/entry/gin/docs
.PHONY: test
	go test ./...
.PHONY: grpc-wire
grpc-wire:
	wire gen cmd/grpc/inject/wire.go
.PHONY: grpc-run
grpc-run: grpc-wire
	go run main.go grpc