package types

type User struct {
	// The id of the user
	Id string `json:"id"`

	// The votes the user has cast
	Votes []*UserVote `json:"vote"`

	// The debates the user has participated in
	Debates []string `json:"debates"`

	// The responses the user has posted
	Responses []string `json:"responses"`

	SetData map[SimulationType]interface{} `json:"set_data"`
}

func NewUser(id string) *User {
	return &User{
		Id: id,
	}
}
