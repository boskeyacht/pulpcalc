package tree

import (
	"log"
	"os"

	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/joho/godotenv"
)

var (
	RedditSecretKey     string
	TwitterAccessKey    string
	TwitterAccessSecret string
	TwitterApiKey       string
	TwitterApiSecret    string
	TwitterBearerToken  string
	NeoEndpoint         string
	NeoUser             string
	NeoPassword         string
	DictServer          string
)

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

func initConfig() *types.Config {
	godotenv.Load("../.env")

	RedditSecretKey = getEnv("REDDIT_SECRET_KEY")
	TwitterAccessKey = getEnv("TWITTER_ACCESS_KEY")
	TwitterAccessSecret = getEnv("TWITTER_ACCESS_SECRET")
	TwitterApiKey = getEnv("TWITTER_API_KEY")
	TwitterApiSecret = getEnv("TWITTER_API_SECRET")
	TwitterBearerToken = getEnv("TWITTER_BEARER_TOKEN")
	NeoEndpoint = getEnv("NEO_ENDPOINT")
	NeoUser = getEnv("NEO_USER")
	NeoPassword = getEnv("NEO_PASSWORD")
	DictServer = getOptionalEnv("DICT_SERVER")

	return types.NewConfig(RedditSecretKey, TwitterAccessKey, TwitterAccessSecret, TwitterApiKey, TwitterApiSecret, NeoEndpoint, NeoUser, NeoPassword, TwitterBearerToken)
}
