build:
	golangci-lint run
	go build 

test:
	go test -cover ./...