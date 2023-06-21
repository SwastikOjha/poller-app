tidy:
	go mod tidy

test:
	go test -coverprofile .coverage.txt ./...

build:
	go build  -o ./cmd/main ./cmd/main.go

clean:
	rm -rf ./cmd/main

run:
	go run ./cmd/main.go