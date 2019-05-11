.PHONY: test
test:
	go test -v -race -cover ./...

all: twitter-streamer queue

twitter-streamer: proto *.go cmd/twitter-streamer/*.go streamer/*.go
	go build -v -o build/twitter-streamer ./cmd/twitter-streamer
queue: proto *.go cmd/queue/*.go queue/*.go
	go build -v -o build/queue ./cmd/queue

.PHONY: dev-twitter-streamer
dev-twitter-streamer:
	gin --immediate --build cmd/twitter-streamer --port 12345
.PHONY: dev-queue
dev-queue:
	gin --immediate --build cmd/queue

proto: queue.pb.go

queue.pb.go: queue.proto
	protoc queue.proto --go_out=plugins=grpc:.

.PHONY: clean
clean:
	rm -rf build gin-bin

.PHONY: vet lint
vet:
	go vet -race ./...
lint:
	golint  ./...
