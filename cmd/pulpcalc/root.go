package pulpcalc

import (
	"github.com/baribari2/pulp-calculator/common/types"
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

func Execute(RedditSecretKey, RedditAccessToken, TwitterAccessKey, TwitterAccessToken, TwitterApiKey, TwitterApiSecret, TwitterBearerToken string, dictClient *dict.Client) {
	cfg.RedditAccessToken = RedditAccessToken
	cfg.RedditSecretKey = RedditSecretKey
	cfg.TwitterAccessKey = TwitterAccessKey
	cfg.TwitterAccessToken = TwitterAccessToken
	cfg.TwitterApiKey = TwitterApiKey
	cfg.TwitterApiSecret = TwitterApiSecret
	cfg.TwitterBearerToken = TwitterBearerToken
	cfg.DictServer = dictClient

	rootCmd.Execute()
}
