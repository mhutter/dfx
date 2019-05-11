package streamer

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/mhutter/dfx"
)

// Twitter is a streaming client that listens to new tweets on the twitter
// streaming API. Tweets are filtered using the "Filter" struct field.
type Twitter struct {
	Filter string

	client *twitter.Client
	stream *twitter.Stream
}

// QueueClient is a client that can send out events
type QueueClient interface{}

// NewTwitter wires up the client and returns a new Twitter instance
func NewTwitter(
	consumerKey,
	consumerSecret,
	accessToken,
	accessTokenSecret,
	filter string,
) *Twitter {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	return &Twitter{
		Filter: filter,
		client: twitter.NewClient(httpClient),
	}
}

// Start starts streaming tweets and returns the chan which will contain
// Deployables derived from tweets.
func (t *Twitter) Start() (<-chan *dfx.Deployable, error) {
	ch := make(chan *dfx.Deployable, 10)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		d := dfx.NewFromTweet(tweet)
		ch <- d
	}

	stream, err := t.client.Streams.Filter(&twitter.StreamFilterParams{
		Track: []string{t.Filter},
	})
	if err != nil {
		return nil, err
	}

	go demux.HandleChan(stream.Messages)
	t.stream = stream
	return ch, nil
}

// Stop stops the streaming client
func (t *Twitter) Stop() {
	t.stream.Stop()
}
