package dfx

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

const (
	srcTwitter = "twitter"
)

// NewFromTweet creates a new Deployable from the given Tweet
func NewFromTweet(tweet *twitter.Tweet) *Deployable {
	surl := fmt.Sprintf(
		"https://twitter.com/%s/status/%s",
		tweet.User.ScreenName,
		tweet.IDStr,
	)
	return &Deployable{
		Content:   tweet.Text,
		From:      tweet.User.ScreenName,
		SourceUrl: surl,
		Source:    srcTwitter,
	}
}
