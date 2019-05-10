package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mhutter/dfx/streamer"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var (
		consumerKey       = mustEnv("CONSUMER_KEY")
		consumerSecret    = mustEnv("CONSUMER_SECRET")
		accessToken       = mustEnv("ACCESS_TOKEN")
		accessTokenSecret = mustEnv("ACCESS_TOKEN_SECRET")
		filter            = mustEnv("FILTER")
	)

	tw := streamer.NewTwitter(consumerKey, consumerSecret, accessToken, accessTokenSecret, filter)

	ch, err := tw.Start()
	if err != nil {
		log.Fatalln("Error starting Twitter streamer:", err)
	}
	log.Println("Streaming messages from twitter")

	go func() {
		for d := range ch {
			log.Printf("%#v", d)
		}
	}()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-interrupt)
	tw.Stop()
}

func mustEnv(name string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}

	log.Fatalf("Mandatory env var %s missing", name)
	return ""
}
