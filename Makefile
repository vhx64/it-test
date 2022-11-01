
bin/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.49.0

prepare: bin/golangci-lint
.PHONY: prepare

lint:
	make prepare && ./bin/golangci-lint run --timeout 3m0s -c ./.golangci.yaml ./...
.PHONY: lint

codegen:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen && \
	oapi-codegen --config=.oapi-codegen.yml api/api.yml && \
	go mod tidy
.PHONY: generate-api

migrate-up:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate && \
	migrate -database "postgres://it:it@localhost:5432/it?sslmode=disable" -path resources/db up
.PHONY: migrate-up

migrate-down:
	migrate -database "postgres://it:it@localhost:5432/order?sslmode=disable" -path resources/db down
.PHONY: migrate-down

docker-up:
	docker-compose -f docker-compose.local.yaml up -d
.PHONY: docker-up

docker-down:
	docker-compose -f docker-compose.local.yaml down
.PHONY: docker-down

build:
	make prepare && go build .
.PHONY: build

start:
	make prepare && go run main.go server
.PHONY: start
