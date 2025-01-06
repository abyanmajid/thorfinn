build:
	env GOOS=linux CGO_ENABLED=0 go build -o bin/clyde-novus ./cmd

start-compose-postgres:
	docker-compose -f orchestration/compose.postgres.yml -p clyde up -d

start-cp: start-compose-postgres

start-compose-mysql:
	docker-compose -f orchestration/compose.mysql.yml -p clyde up -d

start-cm: start-compose-mysql

stop-compose-postgres:
	docker-compose -f orchestration/compose.postgres.yml down --volumes --remove-orphans

stop-cp: stop-compose-postgres

stop-compose-mysql:
	docker-compose -f orchestration/compose.mysql.yml down --volumes --remove-orphans

stop-cm: stop-compose-mysql

tv:
	go mod tidy && go mod vendor

migrate-up:
	goose -dir sql/migrations postgres postgresql://clyde:clyde@localhost:5432/clyde up

migrate-down:
	goose -dir sql/migrations postgres postgresql://clyde:clyde@localhost:5432/clyde down

test:
	go test ./...

typegen:
	protoc -I=./protobufs --go_out=./shared/go --go_opt=paths=source_relative ./protobufs/types.proto
	protoc -I=./protobufs \
		--plugin=protoc-gen-ts=/home/abyanmajid/.nvm/versions/node/v22.3.0/bin/protoc-gen-ts \
		--ts_out=./shared/ts \
		./protobufs/types.proto