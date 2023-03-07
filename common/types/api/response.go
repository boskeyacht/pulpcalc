package types

type SimulateThreadResponse struct {
	// The root of the tree
	Root interface{} `json:"root"`

	// The timestamps of the simulation
	Timestamps []int64 `json:"timestamps"`

	// The last score of the simulation
	LastScore int64 `json:"last_score"`

	// The amount of time the score change was 0
	// In other words, 0 engagements
	InactiveCount int64 `json:"inactive_count"`

	// Map node Id to its children
	Nodes map[string]interface{} `json:"nodes"`
}

func NewSimulateThreadResponse(root interface{}, timestamps []int64, lastScore, inactiveCount int64, nodes map[string]interface{}) *SimulateThreadResponse {
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
	Root interface{} `json:"root"`

	// Map node Id to its children
	Nodes map[string]interface{} `json:"nodes"`
}
