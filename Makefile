build:
	@go build -o bin/journal-be cmd/main.go 

test:
	@go test -v ./...

run: build
	@./bin/journal-be

db-status:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING="root:password@/journal_be" goose -dir="db/migrations" status

up:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING="root:password@/journal_be" goose -dir="db/migrations" up
