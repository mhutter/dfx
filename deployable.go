package dfx

import "github.com/dghubble/go-twitter/twitter"

const (
	srcTwitter = "twitter"
)

// NewFromTweet creates a new Deployable from the given Tweet
func NewFromTweet(tweet *twitter.Tweet) *Deployable {
	return &Deployable{
		Content:   tweet.Text,
		From:      tweet.User.ScreenName,
		SourceUrl: tweet.Source,
		Source:    srcTwitter,
	}
}
