package tree

import (
	"math/rand"

	"github.com/baribari2/pulp-calculator/common/types"
	dict "github.com/baribari2/pulp-calculator/dictionary"
)

func setupPulpTree() (*Tree, error) {
	tree := &Tree{
		Root: &types.Node{
			Id:         0,
			Confidence: 75,
			Score:      0,
			Content:    "I think everyone needs a therapist",
			Action:     types.CommentResponse,
			Engagements: types.Engagements{
				Votes: FillAllVotes(33, 20, 2),
			},
		},
	}

	var i int
	for i < 10000 {
		g, err := dict.Gibber(25)
		if err != nil {
			return nil, err
		}

		tree.Nodes[i] = &types.Node{
			Id:      1,
			Score:   0,
			Content: g,
			Action:  types.CommentReply,
			Engagements: types.Engagements{
				Votes: FillAllVotes(rand.Intn(100), rand.Intn(100), rand.Intn(100)),
			},
		}

		i++
	}

	return tree, nil
}

func setupViralPulpTree() (*Tree, error) {
	tree := &Tree{
		Root: &types.Node{
			Id:         0,
			Confidence: 75,
			Score:      0,
			Content:    "I think everyone needs a therapist",
			Action:     types.CommentResponse,
			Engagements: types.Engagements{
				Votes: FillAllVotes(33, 20, 2),
			},
		},
	}

	var i int
	for i < 10000 {
		g, err := dict.Gibber(25)
		if err != nil {
			return nil, err
		}

		tree.Nodes[i] = &types.Node{
			Id:      1,
			Score:   0,
			Content: g,
			Action:  types.CommentReply,
			Engagements: types.Engagements{
				Votes: FillAllVotes(rand.Intn(100), rand.Intn(100), rand.Intn(100)),
			},
		}

		i++
	}

	return tree, nil
}

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
