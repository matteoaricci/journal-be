build:
	@go build -o bin/journal-be cmd/main.go 

test:
	@go test -v ./...

run: build
	@./bin/journal-be