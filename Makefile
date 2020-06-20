all:
	go run github.com/taciomcosta/chesstournament/cmd/webapi
test:
	go test ./...
cover:
	go test ./... -cover
