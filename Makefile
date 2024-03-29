run:
	@go run cmd/webserver/main.go
test:
	@go test -v ./...

migrate:
	@migrate -path=sql/migrations -database "mysql://root:secret@tcp(localhost:3306)/blogs_api?parseTime=true" up

.PHONY: migrate
