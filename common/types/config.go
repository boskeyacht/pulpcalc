package types

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"golang.org/x/net/dict"
)

type Config struct {
	RedditSecretKey     string
	TwitterAccessKey    string
	TwitterAccessSecret string
	TwitterAccessToken  string
	TwitterApiKey       string
	TwitterApiSecret    string
	TwitterBearerToken  string
	NeoEndpoint         string
	NeoUser             string
	NeoPassword         string
	Neo4j               neo4j.Session
	DictServer          *dict.Client
}

func NewConfig(redditSecretKey, twitterAccess, twitterAccessToken, twitterApiKey, twitterSecret, twitterBearer, neoEndpoint, neoUser, neoPass string) *Config {
	return &Config{
		RedditSecretKey:     redditSecretKey,
		TwitterAccessKey:    twitterAccess,
		TwitterAccessSecret: twitterSecret,
		TwitterAccessToken:  twitterAccessToken,
		TwitterBearerToken:  twitterBearer,
		NeoEndpoint:         neoEndpoint,
		NeoUser:             neoUser,
		NeoPassword:         neoPass,
	}
}
