install:
	go mod tidy

clean:
	go mod tidy
	go fmt ./...
	${GOPATH}/bin/goimports -w .

start:
	go run cmd/main.go

testing:
	go test -v ./...
	${GOPATH}/bin/golangci-lint run