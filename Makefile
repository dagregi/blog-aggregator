build:
	go build -o bin/blog-aggregator

run: build
	./bin/./blog-aggregator

test:
	go test ./...
