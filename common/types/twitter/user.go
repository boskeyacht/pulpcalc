package types

type BasicUsernameSearchResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	// Protected   bool   `json:"protected"`
	// Description string `json:"description"`
	// URL         string `json:"url"`
}
