package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mhutter/dfx/queue"
)

const (
	defaultAddr = ":3185"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	addr := defaultAddr
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	s := queue.NewServer()

	s.Listen(addr)
	log.Printf("Listening on %s", addr)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-interrupt)
	s.Stop()
}
