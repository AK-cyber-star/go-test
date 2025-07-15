build:
	@go build -o ./bin/main ./cmd/api/main.go

run: build
	@./bin/main

test:
	@go test -v ./internal/...

clean:
	@rm -rf ./bin
