package main

import (
	"log"

	// "github.com/baribari2/pulp-calculator/reddit"
	"github.com/baribari2/pulp-calculator/twitter"
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

	// err := reddit.AppOnlyRequest(cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	tweets, err := twitter.SearchTweets(cfg.TwitterBearerToken, "elon")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", tweets)
}
