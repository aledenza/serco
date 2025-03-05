-include .env
export

# commands
lint:
	@golangci-lint run

test:
	@go test -p 1 -v ./tests/... 

format:
	@golines -w -m=120

build:
	@go build -a -gcflags=all="-B" -ldflags="-w -s" -o app

# docker: control services
stop:
	@docker-compose stop

down:
	@docker-compose down -v --remove-orphans

new.migrate:
	@sql-migrate new -env="postgres" ${name}

migrate:
	@sql-migrate up -env="postgres"

rollback:
	@sql-migrate down -env="postgres" -limit=1

migrations_list:
	@sql-migrate status -env="postgres"

# docker: built containers
build.database:
	@docker-compose build database

build.test:
	@docker-compose build test

# docker: exec
database.up:
	@docker-compose run -p 5432:5432 -d database

# docker: commands
dc.test: stop down build.database build.test
	@docker-compose up -d database;
	@docker-compose run test make lint;
	@docker-compose run test make test';
	@docker-compose stop;