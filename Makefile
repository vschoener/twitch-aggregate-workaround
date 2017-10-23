DB = ws_aggregate_storage
DB_CONTAINER = storage
DOCKER_COMPOSE = docker-compose
DOCKER = docker

all: re

gobuild:
	go install ./...

gobuild_linux:
	GOOS=linux GOARCH=amd64 go build ./... -o

goclean:
	go clean

exec_webserver:
	$(GOPATH)/bin/webserver

exec_aggregation:
	$(GOPATH)/bin/aggregation

exec_test:
	$(GOPATH)/bin/test

exec_auth:
	$(GOPATH)/bin/auth

goinstall:
	go get gopkg.in/yaml.v2
	go get github.com/bsphere/le_go
	go get github.com/go-sql-driver/mysql

build: clean
	$(DOCKER_COMPOSE) build

start:
	$(DOCKER_COMPOSE) up -d

stop:
	$(DOCKER_COMPOSE) kill || true

clean: stop
	$(DOCKER_COMPOSE) rm -afv
	$(DOCKER_COMPOSE) down -v

re: stop clean build start

rm: clean

install:
	@make install-db

install-db:
	docker-compose exec $(DB_CONTAINER) bash -c "mysql -u root -proot  < /docker-entrypoint-initdb.d/schema.sql"

ps:
	$(DOCKER_COMPOSE) ps
