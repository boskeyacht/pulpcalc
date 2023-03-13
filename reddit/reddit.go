package reddit

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"

	"github.com/baribari2/pulpcalc/common/types"
	"github.com/baribari2/pulpcalc/simulator"
	"github.com/broothie/qst"
)

// raw_json=1
func AppOnlyRequest(cfg *types.Config) (string, error) {
	res, err := qst.Post(
		"https://www.reddit.com/api/v1/access_token",
		qst.Header("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(cfg.RedditAppId+":"+cfg.RedditSecretKey))),
		qst.Header("Content-Type", "application/json"),
		qst.BodyJSON(map[string]string{
			"grant_type": "client_credentials",
			"raw_json":   "1",
		}),
	)

	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", errors.New("Failed to execute reddit app-only OAUTH request")
	}

	var t struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}

	err = json.NewDecoder(res.Body).Decode(&t)
	if err != nil {
		return "", err
	}

	cfg.RedditBearerToken = t.AccessToken

	return t.AccessToken, nil
}

func SearchCMVSubreddit(cfg *types.Config, query string) (*simulator.Debate, error) {
	res, err := qst.Get(
		"https://reddit.com/r/changemyview/search",
		qst.QueryValue("limit", "25"),
		qst.QueryValue("restrict_sr", "on"),
		qst.QueryValue("sort", "relevance"),
		qst.QueryValue("t", "all"),
		qst.QueryValue("q", query),
		qst.Header("Authorization", "Bearer "+cfg.RedditBearerToken),
	)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Failed to execute reddit search request")
	}

	type z struct {
		Data map[string]interface{} `json:"data"`
	}

	r := &z{}

	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func GetTrendingThread(*types.Config) (*simulator.Debate, error) {

	return nil, nil
}
