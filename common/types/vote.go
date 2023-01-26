package types

type VoteType int8

const (
	ValidVoteType = iota
	InvalidVoteType
	AbstainVoteType
)
