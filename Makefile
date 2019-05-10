.PHONY: test
test:
	go test -v -race -cover ./...

.PHONY: all
all: twitter-streamer

twitter-streamer: *.go cmd/twitter-streamer/*.go
	go build -v -o twitter-streamer ./cmd/twitter-streamer

.PHONY: clean
clean:
	rm -f twitter-streamer gin-bin
