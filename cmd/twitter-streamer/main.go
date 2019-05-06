package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/mhutter/dfx"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := oauth1.NewConfig(mustEnv("CONSUMER_KEY"), mustEnv("CONSUMER_SECRET"))
	token := oauth1.NewToken(mustEnv("ACCESS_TOKEN"), mustEnv("ACCESS_TOKEN_SECRET"))

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		d := dfx.NewFromTweet(tweet)
		log.Printf("New Deployable --> %s: %s (from %s)\n",
			d.From, d.Content, d.SourceURL)
	}

	stream, err := client.Streams.Filter(&twitter.StreamFilterParams{
		Track:         []string{mustEnv("FILTER")},
		StallWarnings: twitter.Bool(true),
	})
	if err != nil {
		log.Fatalf("streaming tweets failed: %+v", err)
	}

	go demux.HandleChan(stream.Messages)
	log.Println("Streaming messages...")

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-interrupt)
	stream.Stop()
}

func mustEnv(name string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}

	log.Fatalf("Mandatory env var %s missing", name)
	return ""
}
