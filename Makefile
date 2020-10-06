all: build

build:
	go build -o ./bin/results ./cmd/results

test:
	go test -race -cover -count=1 ./...