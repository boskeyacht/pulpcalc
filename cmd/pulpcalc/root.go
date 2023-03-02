package pulpcalc

import (
	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
	"golang.org/x/net/dict"
)

var (
	rootCmd = &cobra.Command{
		Use:  "pulpcalc",
		Long: `pulpcalc is a tool for calculating the score of a thread over time`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cfg = &types.Config{}
)

func Execute(RedditSecretKey, TwitterAccessKey, TwitterAccessToken, TwitterApiKey, TwitterApiSecret, TwitterBearerToken string, dictClient *dict.Client, session neo4j.Session) {
	cfg.RedditSecretKey = RedditSecretKey
	cfg.TwitterAccessKey = TwitterAccessKey
	cfg.TwitterAccessToken = TwitterAccessToken
	cfg.TwitterApiKey = TwitterApiKey
	cfg.TwitterApiSecret = TwitterApiSecret
	cfg.TwitterBearerToken = TwitterBearerToken
	cfg.DictServer = dictClient
	cfg.Neo4j = session

	rootCmd.Execute()
}
