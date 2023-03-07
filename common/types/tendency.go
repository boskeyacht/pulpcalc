package types

type ActionTendencies struct {
	ValidVoteTendency float64 `json:"valid_vote_tendency"`

	InvalidVoteTendency float64 `json:"invalid_vote_tendency"`

	AbstainVoteTendency float64 `json:"abstain_vote_tendency"`

	ReportTendency float64 `json:"report_tendency"`

	HideTendency float64 `json:"hide_tendency"`
}
