PROJECT:=nirvana

.PHONY: build
build:
	go build -o $(PROJECT) main.go

.PHONY: run
run:
	go run main.go server

	