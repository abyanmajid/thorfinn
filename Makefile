build:
	chmod +x scripts/build.sh
	./scripts/build.sh

tv:
	go mod tidy
	go mod vendor

infra-postgres:
	docker compose -f docker/compose.postgres.yml up -d

infra-mysql:
	docker compose -f docker/compose.mysql.yml up -d

infra-sqlite:
	docker compose -f docker/compose.sqlite.yml up -d