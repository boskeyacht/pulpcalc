package types

import "github.com/baribari2/pulp-calculator/common/types"

type SimulateThreadResponse struct {
	// The root of the tree
	Root *types.Node `json:"root"`

	// The timestamps of the simulation
	Timestamps []int64 `json:"timestamps"`

	// The last score of the simulation
	LastScore int64 `json:"last_score"`

	// The amount of time the score change was 0
	// In other words, 0 engagements
	InactiveCount int64 `json:"inactive_count"`

	// Map node Id to its children
	Nodes map[int]*types.Node `json:"nodes"`
}

func NewSimulateThreadResponse(root *types.Node, timestamps []int64, lastScore, inactiveCount int64, nodes map[int]*types.Node) *SimulateThreadResponse {
	return &SimulateThreadResponse{
		Root:          root,
		Timestamps:    timestamps,
		LastScore:     lastScore,
		InactiveCount: inactiveCount,
		Nodes:         nodes,
	}
}

type GetTreeResponse struct {
	// The root of the tree
	Root *types.Node `json:"root"`

	// Map node Id to its children
	Nodes map[int]*types.Node `json:"nodes"`
}
