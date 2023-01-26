package tree

import (
	"testing"

	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/baribari2/pulp-calculator/tree/util"
)

func TestTree(t *testing.T) {
	_ = &Tree{
		Root: &types.Node{
			Id:         0,
			Confidence: 75,
			Score:      0,
			Content:    "I think everyone needs a therapist",
			Action:     types.CommentResponse,
			Engagements: types.Engagements{
				Votes: util.FillAllVotes(33, 20, 2),
			},
		},
		Nodes: map[int]*types.Node{
			1: &types.Node{
				Id:         1,
				Confidence: 63,
				Score:      0,
				Content:    "I disagree with you, I think everyone has their own way of dealing with their problems",
				Action:     types.CommentReply,
				Engagements: types.Engagements{
					Votes: util.FillAllVotes(77, 13, 0),
				},
			},
			2: &types.Node{
				Id:         1,
				Confidence: 63,
				Score:      0,
				Content:    "I disagree with you, I think everyone has their own way of dealing with their problems",
				Action:     types.CommentReply,
				Engagements: types.Engagements{
					Votes: util.FillAllVotes(77, 13, 0),
				},
			},
			3: &types.Node{
				Id:         1,
				Confidence: 63,
				Score:      0,
				Content:    "I disagree with you, I think everyone has their own way of dealing with their problems",
				Action:     types.CommentReply,
				Engagements: types.Engagements{
					Votes: util.FillAllVotes(77, 13, 0),
				},
			},
			4: &types.Node{
				Id:         1,
				Confidence: 63,
				Score:      0,
				Content:    "I disagree with you, I think everyone has their own way of dealing with their problems",
				Action:     types.CommentReply,
				Engagements: types.Engagements{
					Votes: util.FillAllVotes(77, 13, 0),
				},
			},
			5: &types.Node{
				Id:         1,
				Confidence: 63,
				Score:      0,
				Content:    "I disagree with you, I think everyone has their own way of dealing with their problems",
				Action:     types.CommentReply,
				Engagements: types.Engagements{
					Votes: util.FillAllVotes(77, 13, 0),
				},
			},
			6: &types.Node{
				Id:         1,
				Confidence: 63,
				Score:      0,
				Content:    "I disagree with you, I think everyone has their own way of dealing with their problems",
				Action:     types.CommentReply,
				Engagements: types.Engagements{
					Votes: util.FillAllVotes(77, 13, 0),
				},
			},
		},
	}
}
