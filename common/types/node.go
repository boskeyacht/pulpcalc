package types

// This node can represent any action in the set of defined actions, namely:
//
//	Comment (Response)
//	Comment (Reply)
//	ValidVote
//	InvalidVote
//	AbstainVote
//	ValidVoteWithContent
//	InvalidVoteWithContent
//	TrustedReference
//	DistrustedReference
//
// Certain actions also have a score, which is the number of points accumulated by generating engagement
// through said action.
type Node struct {
	Id int

	// The Id of the parent node
	ParentId int

	Confidence float64

	// Total points, accumulated from all interactions generated by this action
	Score int64

	Children []*Node

	// This will ultimately be multimedia content
	Content string

	Timestamp int64

	Action

	Exposure

	Engagements
}

// Adds a child to the adjencey list of children nodes
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

// Returns the children of this node
func (n *Node) GetChildren() []*Node {
	return n.Children
}

// Returns the Id of the parent node
func (n *Node) GetParentId() int {
	return n.ParentId
}

// Reutrns the Id of this node
func (n *Node) GetId() int {
	return n.Id
}

func (n *Node) GetTimestamp() int64 {
	return n.Timestamp
}
