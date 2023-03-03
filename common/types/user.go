package types

type UserVote struct {
	// The type of vote (Valid, Invalid, or Abstain)
	Vote *VoteType `json:"vote"`

	// The Id of the debate that was voted on
	DebateId int64 `json:"debate_id"`

	// The ID of the post that was voted on
	PostId int64 `json:"post_id"`
}

type User struct {
	// The id of the user
	Id string `json:"id"`

	// The votes the user has cast
	Votes []*UserVote `json:"vote"`

	// The debates the user has participated in
	Debates []interface{} `json:"debates"`

	// The responses the user has posted
	Responses []*Node `json:"responses"`
}

func NewUser(id string) *User {
	return &User{
		Id: id,
	}
}
