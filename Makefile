build:
	go build ./cmd/server

start:
	go run ./cmd/server

dev:
	go run github.com/cosmtrek/air

install-tools:
	go install go.uber.org/mock/mockgen
	go install github.com/utgwkk/bulkmockgen/cmd/bulkmockgen

generate: install-tools
	go generate ./...

test: generate
	go test -v ./...
