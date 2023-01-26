package util

import (
	"github.com/baribari2/pulp-calculator/common/types"
)

func FillValidVotes(num int) []types.VoteType {
	var votes []types.VoteType
	for i := 0; i < num; i++ {
		votes = append(votes, types.ValidVoteType)
	}
	return votes
}

func FillInalidVotes(num int) []types.VoteType {
	var votes []types.VoteType
	for i := 0; i < num; i++ {
		votes = append(votes, types.ValidVoteType)
	}
	return votes
}

func FillAbstainVotes(num int) []types.VoteType {
	var votes []types.VoteType
	for i := 0; i < num; i++ {
		votes = append(votes, types.ValidVoteType)
	}
	return votes
}

func FillAllVotes(valNum, invalNum, abstainNum int) []types.VoteType {
	var votes []types.VoteType
	for i := 0; i < valNum; i++ {
		votes = append(votes, types.ValidVoteType)
	}
	for i := 0; i < invalNum; i++ {
		votes = append(votes, types.InvalidVoteType)
	}
	for i := 0; i < abstainNum; i++ {
		votes = append(votes, types.AbstainVoteType)
	}
	return votes
}
