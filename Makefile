

all: format test

format:
	go fmt

test:
	go test ./...


testVerbose:
	go test . -v
