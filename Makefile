default: lint clean build

lint:
	golangci-lint run

clean:
	rm -rf dist

build:
	go build -o dist/gosign cmd/main.go
	go run cmd/main.go dist/gosign

test:
	go test ./...
