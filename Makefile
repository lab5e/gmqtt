# Basic Makefile for Golang project
# Includes GRPC Gateway, Protocol Buffers
SERVICE		?= $(shell basename `go list`)
VERSION		?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || cat $(PWD)/.version 2> /dev/null || echo v0)
PACKAGE		?= $(shell go list)
PACKAGES	?= $(shell go list ./...)
FILES		?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")
BUILD_DIR	?= build

.PHONY: help clean fmt lint vet test test-cover build all

default: help

# show this help
help:
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*#\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*#/#/' | column -t -c 2 -s '#'

sample:
	@cd cmd/server && go build -o ../../bin/server

# clean, format, build and unit test
all: clean-all fmt test sample

# Print useful environment variables to stdout
env:
	echo $(CURDIR)
	echo $(SERVICE)
	echo $(PACKAGE)
	echo $(VERSION)

# go clean
clean:
	go clean

# remove all generated artifacts and clean all build artifacts
clean-all:
	go clean -i ./...
	rm -fr build

# fetch and install all required tools
tools:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/golang/lint/golint
	go get github.com/golang/mock/mockgen@v1.4.4

# format the go source files
fmt:
	goimports -w $(FILES)

# run go lint on the source files
lint:
	golint $(PACKAGES)

 # run go vet on the source files
vet:
	go vet ./...

# generate godocs and start a local documentation webserver on port 8085
doc:
	godoc -http=:8085 -index

# update golang dependencies
update-dependencies:
	go mod tidy

# generate mock code
generate-mocks:
	@mockgen -source=persistence/queue/elem.go -destination=./persistence/queue/elem_mock.go -package=queue -self_package=github.com/lab5e/lmqtt/queue
	@mockgen -source=persistence/queue/queue.go -destination=./persistence/queue/queue_mock.go -package=queue -self_package=github.com/lab5e/lmqtt/queue
	@mockgen -source=persistence/session/session.go -destination=./persistence/session/session_mock.go -package=session -self_package=github.com/lab5e/lmqtt/session
	@mockgen -source=persistence/subscription/subscription.go -destination=./persistence/subscription/subscription_mock.go -package=subscription -self_package=github.com/lab5e/lmqtt/subscription
	@mockgen -source=persistence/unack/unack.go -destination=./persistence/unack/unack_mock.go -package=unack -self_package=github.com/lab5e/lmqtt/unack
	@mockgen -source=pkg/packets/packets.go -destination=./pkg/packets/packets_mock.go -package=packets -self_package=github.com/lab5e/lmqtt/packets
	@mockgen -source=retained/interface.go -destination=./retained/interface_mock.go -package=retained -self_package=github.com/lab5e/lmqtt/retained
	@mockgen -source=pkg/gmqtt/client.go -destination=./server/client_mock.go -package=server -self_package=github.com/lab5e/lmqtt/server
	@mockgen -source=pkg/gmqtt/persistence.go -destination=./server/persistence_mock.go -package=server -self_package=github.com/lab5e/lmqtt/server
	@mockgen -source=pkg/gmqtt/plugin.go -destination=./server/plugin_mock.go -package=server -self_package=github.com/lab5e/lmqtt/server
	@mockgen -source=pkg/gmqtt/server.go -destination=./server/server_mock.go -package=server -self_package=github.com/lab5e/lmqtt/server
	@mockgen -source=pkg/gmqtt/service.go -destination=./server/service_mock.go -package=server -self_package=github.com/lab5e/lmqtt/server
	@mockgen -source=pkg/gmqtt/stats.go -destination=./server/stats_mock.go -package=server -self_package=github.com/lab5e/lmqtt/server
	@mockgen -source=pkg/gmqtt/topic_alias.go -destination=./server/topic_alias_mock.go -package=server -self_package=github.com/lab5e/lmqtt/server


go-generate:
	go generate ./...

# generate mocks and run short tests
test: generate-mocks
	go test -v -race ./...

# run benchmark tests
test-bench: generate-mocks
	go test -bench ./...

test-cover:
	rm -fr coverage
	mkdir coverage
	go list -f '{{if gt (len .TestGoFiles) 0}}"go test -covermode count -coverprofile {{.Name}}.coverprofile -coverpkg ./... {{.ImportPath}}"{{end}}' ./... | xargs -I {} bash -c {}
	echo "mode: count" > coverage/cover.out
	grep -h -v "^mode:" *.coverprofile >> "coverage/cover.out"
	rm *.coverprofile
	go tool cover -html=coverage/cover.out -o=coverage/cover.html

test-all: test test-bench test-cover

