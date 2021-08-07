TEST_DB_DSN:= "postgres://book-trck:book-trck@localhost:5432/book_trck?sslmode=disable"

.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o ./bin/boot-trck ./cmd/main.go

.PHONY: test-db-local
test-db-local:
	docker run --rm --name book-trck-pg -e POSTGRES_USER=book-trck \
	-e POSTGRES_PASSWORD=book-trck -e POSTGRES_DB=book_trck -d -p 5432:5432  postgres:11

.PHONY: migrations-up
migrations-up:
	./bin/goose -dir ./scripts/migrations/ postgres $(TEST_DB_DSN) up

.PHONY: migrations-down
migrations-down:
	./bin/goose -dir ./scripts/migrations/ postgres $(TEST_DB_DSN) down

.PHONY: install-goose
install-goose:
	go get -u github.com/pressly/goose/v3/cmd/goose