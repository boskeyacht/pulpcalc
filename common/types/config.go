package types

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/sashabaranov/go-openai"
	"github.com/sausheong/goreplicate"
	"golang.org/x/net/dict"
)

type Config struct {
	RedditAppId         string
	RedditSecretKey     string
	RedditBearerToken   string
	TwitterAccessKey    string
	TwitterAccessSecret string
	TwitterAccessToken  string
	TwitterApiKey       string
	TwitterApiSecret    string
	TwitterBearerToken  string
	NeoEndpoint         string
	NeoUser             string
	NeoPassword         string
	OpenAIKey           string
	ReplicateKey        string
	DictEndpoint        string
	ReplicateClient     *goreplicate.Client
	Neo4j               neo4j.Session
	OpenAIClient        *openai.Client
	DictServer          *dict.Client
}

func getEnv(e string) string {
	env, exists := os.LookupEnv(e)
	if !exists {
		log.Fatalf("Environment variable %s not found", env)
	}

	return env
}

func getOptionalEnv(e string) string {
	env, _ := os.LookupEnv(e)

	return env
}

func NewConfig(redditAppId, redditSecretKey, twitterAccess, twitterAccessToken, twitterApiKey, twitterSecret, twitterBearer, neoEndpoint, neoUser, neoPass, openAIKey, replicateKey, dictEndpoint string) *Config {
	return &Config{
		RedditAppId:         redditAppId,
		RedditSecretKey:     redditSecretKey,
		TwitterAccessKey:    twitterAccess,
		TwitterAccessSecret: twitterSecret,
		TwitterAccessToken:  twitterAccessToken,
		TwitterBearerToken:  twitterBearer,
		NeoEndpoint:         neoEndpoint,
		NeoUser:             neoUser,
		NeoPassword:         neoPass,
		OpenAIKey:           openAIKey,
		ReplicateKey:        replicateKey,
		DictEndpoint:        dictEndpoint,
	}
}

func InitConfig() *Config {
	// Try to load .env file in the current working directory.
	// Repeat for parent directories until succeeding, or reaching root.
	cwd, _ := os.Getwd()
	log.Printf("Loading environment in %s", cwd)
	for cwd != "." && cwd != "/" {
		p := filepath.Join(cwd, ".env")
		err := godotenv.Load(p)
		cwd = filepath.Dir(cwd)
		if err == nil {
			log.Printf("Loaded env file: %s", p)
			break
		}
	}

	cfg := NewConfig(
		getEnv("REDDIT_APP_ID"),
		getEnv("REDDIT_SECRET_KEY"),
		getEnv("TWITTER_ACCESS_KEY"),
		getEnv("TWITTER_ACCESS_SECRET"),
		getEnv("TWITTER_API_KEY"),
		getEnv("TWITTER_API_SECRET"),
		getEnv("TWITTER_BEARER_TOKEN"),
		getEnv("NEO_ENDPOINT"),
		getEnv("NEO_USER"),
		getEnv("NEO_PASSWORD"),
		getEnv("OPENAI_KEY"),
		getEnv("REPLICATE_KEY"),
		getOptionalEnv("DICT_SERVER"),
	)

	return cfg
}
