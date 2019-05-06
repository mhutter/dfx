package dfx

import "github.com/dghubble/go-twitter/twitter"

// Source from where a deployable came from
type Source string

const (
	srcTwitter Source = "twitter"
)

// Deployable is some payload that can be deployed into a DFT Pod
type Deployable struct {
	Content   string `json:"content"`
	From      string `json:"from"`
	SourceURL string `json:"source_url"`
	Source    Source `json:"source"`
}

// NewFromTweet creates a new Deployable from the given Tweet
func NewFromTweet(tweet *twitter.Tweet) *Deployable {
	return &Deployable{
		Content:   tweet.Text,
		From:      tweet.User.ScreenName,
		SourceURL: tweet.Source,
		Source:    srcTwitter,
	}
}
