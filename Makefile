.PHONY: test
test:
	go clean -testcache
	go test ./... -v

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	go run ./cmd/main.go -env dev

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: mUp
mUp:
	migrate -path ./migrates/postgres -database 'postgres://yMetric:QWBOVFDMRC@localhost:5432/postgres?sslmode=disable' up

.PHONY: mDown
mDown:
	migrate -path ./migrates/postgres -database 'postgres://yMetric:QWBOVFDMRC@localhost:5432/postgres?sslmode=disable' down

.DEFAULT_GOAL := run