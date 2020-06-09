PROJECT:=nirvana

.PHONY: build
build: wire
	go build -o $(PROJECT) main.go
.PHONY: run
run:
	go run main.go http
.PHONY: wire
wire:
	wire gen cmd/http/inject/wire.go
.PHONY: swag
swag:
	@echo "generate gin swagger doc."
	swag init --output=internal/entry/gin/docs
