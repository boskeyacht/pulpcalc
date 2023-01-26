package reddit

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/baribari2/pulp-calculator/tree"
	"github.com/broothie/qst"
	"github.com/google/uuid"
)

// raw_json=1
func AppOnlyRequest(cfg *types.Config) error {
	id := uuid.New().String()

	log.Printf("Len: %v", len(id))

	res, err := qst.Post(
		"https://www.reddit.com/api/v1/access_token",
		// qst.QueryValue("grant_type", "client_credentials"),
		qst.BodyJSON(map[string]string{
			"grant_type": "client_credentials",
			// "device_id":  id,
			// "client_id":     cfg.RedditAccessKey,
			// "client_secret": cfg.RedditSecretKey,
		}),
	)

	log.Printf("\x1b[32m%s\x1b[0m%v", "Request:", res.Request)
	log.Printf("\x1b[33m%s\x1b[0m%v", "Response:", res)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("Failed to execute reddit app-only OAUTH request")
	}

	var t struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}

	err = json.NewDecoder(res.Body).Decode(&t)
	if err != nil {
		return err
	}

	cfg.RedditAccessToken = t.AccessToken

	log.Printf("Key: %v", cfg.RedditAccessToken)

	return nil
}

func GetTrendingThread(*types.Config) (*tree.Tree, error) {

	return nil, nil
}
