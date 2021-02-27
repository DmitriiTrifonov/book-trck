.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: test
test:
	go test ./...

.PHONY:build
build:
	go build -o ./bin/boot-trck ./cmd/main.go