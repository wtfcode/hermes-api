run:
	go run cmd/main.go
build:
	go build -o exe cmd/main.go && ./exe && rm -f ./exe
test:
	go test -v ./...