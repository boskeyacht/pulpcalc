package simulator

import (
	"math/rand"

	"github.com/baribari2/pulp-calculator/common/types"
)

func FillValidVotes(num int) []types.VoteType {
	var votes []types.VoteType
	for i := 0; i < num; i++ {
		votes = append(votes, types.ValidVoteType)
	}
	return votes
}

func FillInvalidVotes(num int) []types.VoteType {
	var votes []types.VoteType
	for i := 0; i < num; i++ {
		votes = append(votes, types.InvalidVoteType)
	}
	return votes
}

func FillAbstainVotes(num int) []types.VoteType {
	var votes []types.VoteType
	for i := 0; i < num; i++ {
		votes = append(votes, types.AbstainVoteType)
	}
	return votes
}

func FillAllVotes(valNum, invalNum, abstainNum int) []types.VoteType {
	var votes []types.VoteType
	for i := 0; i < rand.Intn(valNum); i++ {
		votes = append(votes, types.ValidVoteType)
	}
	for i := 0; i < rand.Intn(invalNum); i++ {
		votes = append(votes, types.InvalidVoteType)
	}
	for i := 0; i < rand.Intn(abstainNum); i++ {
		votes = append(votes, types.AbstainVoteType)
	}
	return votes
}

func FillReports(reports int) []*types.Report {
	var r []*types.Report
	for i := 0; i < reports; i++ {
		r = append(r, &types.Report{
			ReportedId: "reported",
			Reason:     types.ReasonHarmfulToOthers,
		})
	}
	return r
}

func MakeNewReport(reason types.Reason, rId string) types.Report {
	return types.Report{
		ReportedId: rId,
		Reason:     reason,
	}
}

func MapVotesToUsers() {

}
