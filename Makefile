all:
	go run github.com/taciomcosta/chesstournament/cmd/webapi
test:
	go test ./...
cover:
	go test ./... -cover
cover-open:
	go test ./... -coverprofile=cover.out
	go tool cover -html=cover.out
