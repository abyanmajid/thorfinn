build:
	env GOOS=linux CGO_ENABLED=0 go build -o bin/thorfinn ./cmd

dev:
	go run ./cmd/main.go -dev=true

tv:
	go mod tidy && go mod vendor

migrate-up:
	goose -dir sql/migrations postgres postgresql://thorfinn:thorfinn@localhost:5433/thorfinn up

migrate-down:
	goose -dir sql/migrations postgres postgresql://thorfinn:thorfinn@localhost:5433/thorfinn down