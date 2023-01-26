package types

type Meta struct {
	Count         int    `json:"count"`
	NewestId      string `json:"newest_id"`
	OldestId      string `json:"oldest_id"`
	NextToken     string `json:"next_token"`
	PreviousToken string `json:"previous_token"`
}
