test:
	go fmt $(go list ./... | grep -v /vendor/)
	go vet $(go list ./... | grep -v /vendor/)
	go test -v -race $(go list ./... | grep -v /vendor/)

build: bin
	go build -o bin ./...

bin:
	@mkdir bin