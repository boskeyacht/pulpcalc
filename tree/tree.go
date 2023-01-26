package tree

import (
	"github.com/baribari2/pulp-calculator/common/types"
	dict "github.com/baribari2/pulp-calculator/dictionary"
)

// TODO: Decay for comment replies
// TODO: Return error msgs instead of 0's
type Tree struct {

	// The root of the tree
	Root *types.Node

	// Map node Id to its children
	Nodes map[int]*types.Node
}

// Implements DFS to calculate the score of each node starting from the root node, and returning the
// score of the root node.
// The root node of a tree is any user action (comment, vote, etc).
//
// This function can also be used to calculate the score of a single node (like a vote without content),
// by passing a tree that contains a single node as the root.
func Calculate(cfg *types.Config /*tree *Tree,*/, node *types.Node) int {
	if node == nil {
		return 0
	}

	// If this node has no children, it is a leaf node. In that case, return the score of this node.
	// The score of this node may involve calculations regarding confidence, and other characteristics like
	// word count, common words, links, etc.
	if len(node.GetChildren()) == 0 {
		return CalculateScore(cfg, node)
	}

	// If the node does have children, traverse through them and calculate their scores.
	var sum int
	for _, node := range node.GetChildren() {
		sum += Calculate(cfg, node)
	}

	return sum
}

// Calculates the score of a node, based on its action, content, vote, and confidence.
//
// TODO: Check underflow
// TODO: Can votes generate content?
// TODO: Vote:Confidence ratio?
// TODO: Measure strength through words
// TODO: References
func CalculateScore(cfg *types.Config, node *types.Node) int {
	// Set the score equal to the base points of the action
	var score int = int(node.Action.BasePoints())

	// If the action is a vote w/o content it's a constant - return the base value
	if node.Action == types.ValidVote || node.Action == types.InvalidVote || node.Action == types.AbstainVote {
		score += int(node.Action.BasePoints())

		return score
	}

	// If the action contians content, then calculate a portion of the score based on the content
	if node.Action == types.CommentResponse ||
		node.Action == types.CommentReply ||
		node.Action == types.ValidVoteWithContent ||
		node.Action == types.InvalidVoteWithContent {

		s, c, err := dict.CountCorrectAndCommonWords(cfg, node.Content)
		if err != nil {
			return 0
		}

		// The amount of correct words (positive characteristic) - the amount of common words (negative characteristic)
		score += (s - c)
	}

	// If the action has any votes, then calculate a portion of the score based on the votes
	for _, vote := range node.Engagements.Votes {
		switch vote {
		case types.ValidVoteType:
			score += int(types.ValidVote.BasePoints())

		case types.InvalidVoteType:
			score += int(types.InvalidVote.BasePoints())

		case types.AbstainVoteType:
			score += int(types.AbstainVote.BasePoints())

		default:
			return 0
		}
	}

	// Calculate a portion of the score based on the confidence
	score += int(node.Confidence) / 100

	// If the action has any references, then calculate a portion of the score based on the references

	return score
}
