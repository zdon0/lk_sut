include .env
export

.PHONY: test
test:
	go test -race -v ./...

.PHONY: docs
docs:
	swag init -g cmd/main/main.go --parseInternal --parseDependency

.PHONY: lint
lint:
	golangci-lint run --fix ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: prepare
prepare: tidy lint test docs

.PHONY: up
up:
	docker compose up --build --wait

.PHONY: down
down:
	docker compose down
