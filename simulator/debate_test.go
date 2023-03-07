package simulator

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/baribari2/pulp-calculator/common/types"
	dict "github.com/baribari2/pulp-calculator/dictionary"
)

func TestTree(t *testing.T) {
	cfg := types.InitConfig()
	tree, err := setupPulpDebate()
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
	cfg := types.InitConfig()
	tree, err := setupViralPulpDebate()
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
	cfg := types.InitConfig()
	tree := &Debate{
		Root: &Response{
			Id:         "1",
			Confidence: 63,
			Score:      0,
			Action:     types.CommentReply,
			Content:    NewContent(0, "This is a comment that should meet the minimun character count, because it's longggggggggggggggggggggg"),
			Engagements: &types.Engagements{
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
	cfg := types.InitConfig()
	tree := &Debate{
		Root: &Response{
			Id:         "1",
			Confidence: rand.Float64(),
			Score:      0,
			Action:     types.CommentReply,
			Content:    NewContent(0, "This is a comment that should meet the minimun character count, because it's longggggggggggggggggggggg"),
			Engagements: &types.Engagements{
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
	cfg := types.InitConfig()
	tree := &Debate{
		Root: &Response{
			Id:         "1",
			Confidence: 63,
			Score:      0,
			Action:     types.CommentReply,
			Content:    NewContent(0, "This is a comment that shouldn't meet the minimun character count"),
			Engagements: &types.Engagements{
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

func setupPulpDebate() (*Debate, error) {
	tree := &Debate{
		Root: &Response{
			Id:         "0",
			Confidence: 75,
			Score:      0,
			Content:    NewContent(0, "I think everyone needs a therapist"),
			Action:     types.CommentResponse,
			Engagements: &types.Engagements{
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

		tree.Responses[fmt.Sprintln(i)] = &Response{
			Id:      "1",
			Score:   0,
			Content: NewContent(0, g),
			Action:  types.CommentReply,
			Engagements: &types.Engagements{
				Votes: FillAllVotes(rand.Intn(100), rand.Intn(100), rand.Intn(100)),
			},
		}

		i++
	}

	return tree, nil
}

func setupViralPulpDebate() (*Debate, error) {
	tree := &Debate{
		Root: &Response{
			Id:         "0",
			Confidence: 75,
			Score:      0,
			Content:    NewContent(0, "I think everyone needs a therapist"),
			Action:     types.CommentResponse,
			Engagements: &types.Engagements{
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

		tree.Responses[fmt.Sprintln(i)] = &Response{
			Id:      "1",
			Score:   0,
			Content: NewContent(0, g),
			Action:  types.CommentReply,
			Engagements: &types.Engagements{
				Votes: FillAllVotes(rand.Intn(100), rand.Intn(100), rand.Intn(100)),
			},
		}

		i++
	}

	return tree, nil
}
