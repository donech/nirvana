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
cron-wire:
	wire gen cmd/cron/inject/wire.go
.PHONY: cron-run
cron-run: cron-wire
	go run main.go cron -c /Users/piaoweixiong/donech/nirvana/cmd/app.yaml --pf=/Users/piaoweixiong/lotto/period --tp=125 --sp=125 --tn='02,07,09,10,23,27|12,08' --sn='07,09,10,23,27|02,18,12'