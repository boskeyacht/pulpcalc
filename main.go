package main

import (
	"log"

	cmd "github.com/baribari2/pulp-calculator/cmd/pulpcalc"
	"golang.org/x/net/dict"
)

func main() {
	cfg := initConfig()

	if DictServer == "" {
		DictServer = "dict://dict.dict.org/"
	} else {
		s, err := dict.Dial("tcp", DictServer)
		if err != nil {
			log.Fatal(err)
		}

		cfg.DictServer = s
	}

	cmd.Execute(cfg.RedditAccessToken, cfg.RedditSecretKey, cfg.TwitterAccessKey, cfg.TwitterAccessToken, cfg.TwitterApiKey, cfg.TwitterApiSecret, cfg.TwitterBearerToken, cfg.DictServer)
}
