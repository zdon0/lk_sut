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
	golangci-lint run ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: prepare
prepare: tidy lint test docs

.PHONY: compose-up
compose-up:
	docker compose up --build -d

.PHONY: compose-down
compose-down:
	docker compose down --remove-orphans