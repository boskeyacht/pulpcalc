package types

import (
	"golang.org/x/net/dict"
)

type Config struct {
	RedditSecretKey     string
	RedditAccessToken   string
	TwitterAccessKey    string
	TwitterAccessSecret string
	TwitterAccessToken  string
	TwitterBearerToken  string
	DictServer          *dict.Client
}

func NewConfig(redditSecret, twitterAccess, twitterSecret, twitterBearer string) *Config {
	return &Config{
		RedditSecretKey:     redditSecret,
		RedditAccessToken:   "",
		TwitterAccessKey:    twitterAccess,
		TwitterAccessSecret: twitterSecret,
		TwitterAccessToken:  "",
		TwitterBearerToken:  twitterBearer,
	}
}
