package types

type TweetResponse struct {
	Id      string   `json:"id"`
	Text    string   `json:"text"`
	Author  string   `json:"author_id"`
	EditIds []string `json:"edit_history_tweet_ids"`
	Meta
}
