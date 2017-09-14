all: build exec

exec: build
	./twitch

run:
	go run

build:
	go build

clean:
	go clean

install:
	go get gopkg.in/yaml.v2
