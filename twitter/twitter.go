package twitter

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	types "github.com/baribari2/pulp-calculator/common/types/twitter"
	"github.com/broothie/qst"
)

func AppOnlyRequest(key, secret string) (string, error) {
	res, err := qst.Post(
		fmt.Sprintf("%v/oauth2/token", "https://api.twitter.com"),
		qst.Header("Authorization", fmt.Sprintf("Basic %v", b64.StdEncoding.EncodeToString([]byte(key+":"+secret)))),
		qst.Header("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8"),
		qst.QueryValue("grant_type", "client_credentials"),
	)

	if err != nil {
		return "", err
	}

	var auth struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
	}

	err = json.NewDecoder(res.Body).Decode(&auth)
	if err != nil {
		return "", err
	}

	if auth.TokenType != "bearer" {
		return "", errors.New("Authentication response does not contain token of type 'bearer'")
	}

	return auth.AccessToken, nil
}

func SearchTweets(token, query string) ([]types.TweetResponse, error) {
	res, err := qst.Get(
		fmt.Sprintf("%v/2/tweets/search/recent", "https://api.twitter.com"),
		qst.Header("Authorization", fmt.Sprintf("Bearer %v", token)),
		qst.QueryValue("query", query),
		qst.QueryValue("expansions", "author_id"),
		qst.QueryValue("media.fields", "url"),
		qst.QueryValue("tweet.fields", "attachments,public_metrics"),
	)

	log.Printf("Tweets res: %v", res)

	if err != nil {
		log.Printf("Failed to search tweets: %v", err.Error())
		return nil, err
	}

	type tr []types.TweetResponse
	var tweet struct {
		tr `json:"data"`
	}

	err = json.NewDecoder(res.Body).Decode(&tweet)
	if err != nil {
		log.Printf("Failed to decode tweets: %v", err.Error())
		return nil, err
	}

	log.Printf("Tweets res: %v", tweet)

	return tweet.tr, nil
}
