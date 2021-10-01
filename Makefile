all: sample
	@go test -cover ./...

sample:
	@cd cmd/server && go build -o ../../bin/server

