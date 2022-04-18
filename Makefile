default: build

build:
	go build -o bin/go-fdup cmd/go-fdup/main.go

run:
	go run cmd/go-fdup/main.go

test:
	go test -v ./...

fmt:
	go fmt ./...

lint:
	go vet ./...

clean:
	go clean
	rm -r bin/*
