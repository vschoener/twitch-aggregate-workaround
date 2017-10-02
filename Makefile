all: build exec

build:
	go install ./...

clean:
	go clean

exec_webserver:
	$(GOPATH)/bin/webserver

exec_aggregation:
	$(GOPATH)/bin/aggregation

install:
	go get gopkg.in/yaml.v2
	go get github.com/bsphere/le_go
