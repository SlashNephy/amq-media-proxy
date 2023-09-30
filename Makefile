build:
	go build ./cmd/server

build-batch-download:
	go build ./cmd/batch-download

start:
	go run ./cmd/server

start-batch-download:
	go run ./cmd/batch-download

dev:
	go run github.com/cosmtrek/air

install-tools:
	go install go.uber.org/mock/mockgen
	go install github.com/utgwkk/bulkmockgen/cmd/bulkmockgen

generate: install-tools
	go generate ./...

test: generate
	go test -v ./...
