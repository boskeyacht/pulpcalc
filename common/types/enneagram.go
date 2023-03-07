package types

type EnneagramData struct {
	PersonalityType int `json:"personality_type"`

	Tendencies *ActionTendencies `json:"action_tendency"`
}
