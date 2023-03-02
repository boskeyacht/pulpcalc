package types

type Engagements struct {
	// The amount of times this content was reported
	// Rank 1
	Report int64 `json:"report"`

	// The amount of times this content was hidden
	// Rank 2
	Hide int64 `json:"hide"`

	// Valid, Invalid, or Abstain with the option of content
	// Rank 3
	Votes []VoteType `json:"votes"`

	// Responses to the content
	// Rank 4
	Response *Node `json:"response"`
}

func NewEngagements() *Engagements {
	return &Engagements{}
}
