watch:
	@go run github.com/rafaelsq/wtc

generate-mock:
	go generate ./...

docker-build:
	docker build -t go-clean .

docker-run: docker-build
	docker run -it --rm --name go-clean -p 9000:9000 go-clean 

compose-down:
	docker compose down --remove-orphans

compose:
	docker compose up -d --build

open-adminer:
	@open "http://localhost:8081/?pgsql=postgres&username=root&db=go-clean"

open-jaeger:
	@open "http://localhost:16686"

add-migration:
	migrate create -ext sql -dir internal/migrations -seq $(name)

execute-migrations:
	go run cmd/migrate/main.go $(args)

linter:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run ./...