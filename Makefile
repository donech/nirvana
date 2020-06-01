PROJECT:=nirvana

.PHONY: build
build:
	go build -o $(PROJECT) main.go
.PHONY: run
run:
	go run main.go http
.PHONY: wire
wire:
	wire gen cmd/http/inject/wire.go