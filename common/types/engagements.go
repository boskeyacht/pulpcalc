package types

const (
	HarmfulToOthersPenalty = 1
	AbuseOfPlatformPenalty = .80
	HidePenalty            = .50
	VoteMultiplier         = .20
	DistancePenalty        = .20
	TimingPenalty          = .20
)

type Engagements struct {
	// The amount of times this content was reported
	// Rank 1
	Reports []*Report `json:"report"`

	// The amount of times this content was hidden
	// Rank 2
	HideCount int `json:"hide"`

	// Valid, Invalid, or Abstain with the option of content
	// Rank 3
	Votes []VoteType `json:"votes"`

	// Responses to the content
	// Rank 4
	Response int64 `json:"response"`
}

func NewEngagements(reports []*Report, hideCount int, votes []VoteType, response int64) *Engagements {
	return &Engagements{
		Reports:   reports,
		HideCount: hideCount,
		Votes:     votes,
		Response:  response,
	}
}

func NewEngagementsDefault() *Engagements {
	return &Engagements{}
}
