
lint:
	make prepare && ./bin/golangci-lint run --timeout 3m0s -c ./.golangci.yaml ./...
.PHONY: lint

codegen:
	go get -d github.com/deepmap/oapi-codegen/cmd/oapi-codegen && \
	oapi-codegen --config=.oapi-codegen.yml api/api.yml && \
	go mod tidy
.PHONY: generate-api

migrate-up:
	migrate -database "postgres://it:it@localhost:5432/it?sslmode=disable" -path resources/db up
.PHONY: migrate-up

migrate-down:
	migrate -database "postgres://it:it@localhost:5432/order?sslmode=disable" -path resources/db down
.PHONY: migrate-down

build:
	make prepare && go build .
.PHONY: build

start:
	make prepare && go run main.go server
.PHONY: start
