package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mhutter/dfx"
	"github.com/mhutter/dfx/streamer"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var (
		consumerKey       = mustEnv("CONSUMER_KEY")
		consumerSecret    = mustEnv("CONSUMER_SECRET")
		accessToken       = mustEnv("ACCESS_TOKEN")
		accessTokenSecret = mustEnv("ACCESS_TOKEN_SECRET")
		filter            = mustEnv("FILTER")
		queueAddr         = envOr("QUEUE", "queue:3185")
	)

	conn, err := grpc.Dial(queueAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to queue at %s: %s", queueAddr, err)
	}
	defer conn.Close()
	qc := dfx.NewQueueClient(conn)
	tw := streamer.NewTwitter(consumerKey, consumerSecret, accessToken, accessTokenSecret, filter)

	ch, err := tw.Start()
	if err != nil {
		log.Fatalln("Error starting Twitter streamer:", err)
	}
	log.Printf("Listening for tweets, tracking '%s'\n", filter)
	log.Printf("Sending events to '%s'\n", queueAddr)

	go func() {
		for d := range ch {
			log.Printf("%#v", d)
			if _, err := qc.PostEvent(context.Background(), d); err != nil {
				log.Println("Could not post event: ", err)
			}
		}
	}()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-interrupt)
	tw.Stop()
	<-ch
}

func mustEnv(name string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}

	log.Fatalf("Mandatory env var %s missing", name)
	return ""
}

func envOr(name, fallback string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return fallback
}
