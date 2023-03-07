package types

type VoteType int8

const (
	ValidVoteType = iota
	InvalidVoteType
	AbstainVoteType
)

type UserVote struct {
	// The type of vote (Valid, Invalid, or Abstain)
	Vote *VoteType `json:"vote"`

	// The Id of the debate that was voted on
	DebateId int64 `json:"debate_id"`

	// The ID of the post that was voted on
	PostId int64 `json:"post_id"`
}
