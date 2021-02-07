run:
	go run ./cmd/main.go

test:
	go test ./...

build:
	go build -o ./bin/boot-trck ./cmd/main.go