package tree

import (
	"math/rand"
	"testing"

	"github.com/baribari2/pulp-calculator/common/types"
)

func TestTree(t *testing.T) {
	cfg := initConfig()
	tree, err := setupPulpTree()
	if err != nil {
		t.Errorf("Failed to setup tree: %v", err.Error())
	}

	score, err := Calculate(cfg, tree.Root)
	if err != nil {
		t.Errorf("Failed to calculate score: %v", err.Error())
	}

	t.Logf("Pulp Thread Score: %v", score)
}

func TestViralTree(t *testing.T) {
	cfg := initConfig()
	tree, err := setupViralPulpTree()
	if err != nil {
		t.Errorf("Failed to setup tree: %v", err.Error())
	}

	score, err := Calculate(cfg, tree.Root)
	if err != nil {
		t.Errorf("Failed to calculate score: %v", err.Error())
	}

	t.Logf("Viral Thread Score: %v", score)
}

func TestSingleComment(t *testing.T) {
	cfg := initConfig()
	tree := &Tree{
		Root: &types.Node{
			Id:         1,
			Confidence: 63,
			Score:      0,
			Action:     types.CommentReply,
			Content:    "This is a comment that should meet the minimun character count, because it's longggggggggggggggggggggg",
			Engagements: types.Engagements{
				Votes: FillAllVotes(100, 10, 5),
			},
		},
	}

	score, err := Calculate(cfg, tree.Root)
	if err != nil {
		t.Errorf("Failed to calculate comment: %v", err.Error())
	}

	t.Logf("Single Comment Score: %v", score)
}

func TestRandomSingleComment(t *testing.T) {
	cfg := initConfig()
	tree := &Tree{
		Root: &types.Node{
			Id:         1,
			Confidence: rand.Float64(),
			Score:      0,
			Action:     types.CommentReply,
			Content:    "This is a comment that should meet the minimun character count, because it's longggggggggggggggggggggg",
			Engagements: types.Engagements{
				Votes: FillAllVotes(100, 10, 5),
			},
		},
	}

	score, err := Calculate(cfg, tree.Root)
	if err != nil {
		t.Errorf("Failed to calculate comment: %v", err.Error())
	}

	t.Logf("Single Comment Score: %v", score)
}

func TestSingleShortComment(t *testing.T) {
	cfg := initConfig()
	tree := &Tree{
		Root: &types.Node{
			Id:         1,
			Confidence: 63,
			Score:      0,
			Action:     types.CommentReply,
			Content:    "This is a comment that shouldn't meet the minimun character count",
			Engagements: types.Engagements{
				Votes: FillAllVotes(33, 20, 2),
			},
		},
	}

	score, err := Calculate(cfg, tree.Root)
	if err != nil {
		t.Errorf("Failed to calculate comment: %v", err.Error())
	}

	if score != 0 {
		t.Errorf("Single Short Comment Score: %v", score)
	}

}
