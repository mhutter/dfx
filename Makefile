twitter-streamer: *.go cmd/twitter-streamer/*.go
	go build -v -o twitter-streamer ./cmd/twitter-streamer

.PHONY: test
test:
	go test -v -race -cover ./...
