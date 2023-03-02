package neo4j

import (
	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

type UserVote struct {
	// The type of vote (Valid, Invalid, or Abstain)
	Vote *types.VoteType `json:"vote"`

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
}

func NewUser() *User {
	return &User{}
}

// Creates a new user in neo4j
func (u *User) Create() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("CREATE (u:User {id: $id, votes: $votes, debates: $debates}) RETURN u",
			map[string]interface{}{
				"id":      uuid.New().String(),
				"votes":   u.Votes,
				"debates": u.Debates,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().GetByIndex(0), nil
		}

		return nil, res.Err()
	}, nil
}

// Updates a user in neo4j
func (u *User) Update() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (u:User {id: $id}) SET u.votes = $votes, u.debates = $debates RETURN u",
			map[string]interface{}{
				"id":      u.Id,
				"votes":   u.Votes,
				"debates": u.Debates,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().GetByIndex(0), nil
		}

		return nil, res.Err()
	}, nil
}

// Deletes a user in neo4j
func (u *User) Delete() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (u:User {id: $id}) DETACH DELETE u",
			map[string]interface{}{
				"id": u.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().GetByIndex(0), nil
		}

		return nil, res.Err()
	}, nil
}

// Retrieves a user from neo4j
func (u *User) GetUser() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (u:User {id: $id}) RETURN u",
			map[string]interface{}{
				"id": u.Id,
			})

		if err != nil {
			return nil, err
		}

		rec, err := res.Single()
		if err != nil {
			return nil, err
		}

		user := rec.GetByIndex(0).(dbtype.Node).Props

		return &User{
			Id:      user["id"].(string),
			Votes:   user["votes"].([]*UserVote),
			Debates: user["debates"].([]interface{}),
		}, nil
	}, nil
}

// Adds a relationship on the user to the response that they created
func (u *User) AddUserResponseRelationship(r *Response) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (u:User {id: $id}), (r:Response {id: $response_id}) CREATE (u)-[ur:USER_RESPONSE]->(r) RETURN ur",
			map[string]interface{}{
				"id":          u.Id,
				"response_id": r.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().GetByIndex(0), nil
		}

		return nil, res.Err()
	}, nil
}

// Adds a relationship on the user to the debate that they participated in
func (u *User) AddUserDebateRelationship(t *Tree) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (u:User {id: $id}), (d:Debate {id: $debate_id}) CREATE (u)-[ud:USER_DEBATE]->(d) RETURN ud",
			map[string]interface{}{
				"id":        u.Id,
				"debate_id": t.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().GetByIndex(0), nil
		}

		return nil, res.Err()
	}, nil
}

// Adds a relationship on the user to the response that they voted as valid
func (u *User) AddUserVotedValidResponseRelationship(r *Response) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (u:User {id: $id}), (r:Response {id: $response_id}) CREATE (u)-[uv:USER_VOTED_VALID]->(r) RETURN uv",
			map[string]interface{}{
				"id":          u.Id,
				"response_id": r.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().GetByIndex(0), nil
		}

		return nil, res.Err()
	}, nil
}

// Adds a relationship on the user to the response that they voted as invalid
func (u *User) AddUserVotedInvalidResponseRelationship(r *Response) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (u:User {id: $id}), (r:Response {id: $response_id}) CREATE (u)-[uv:USER_VOTED_INVALID]->(r) RETURN uv",
			map[string]interface{}{
				"id":          u.Id,
				"response_id": r.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().GetByIndex(0), nil
		}

		return nil, res.Err()
	}, nil
}

// Adds a relationship on the user to the response that they voted as abstain
func (u *User) AddUserVotedAbstainResponseRelationship(r *Response) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (u:User {id: $id}), (r:Response {id: $response_id}) CREATE (u)-[uv:USER_VOTED_ABSTAIN]->(r) RETURN uv",
			map[string]interface{}{
				"id":          u.Id,
				"response_id": r.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().GetByIndex(0), nil
		}

		return nil, res.Err()
	}, nil
}

// Adds a relationship on the user to the debate that they voted as valid
func (u *User) AddUserVotedValidDebateRelationship(t *Tree) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (u:User {id: $id}), (d:Debate {id: $debate_id}) CREATE (u)-[uv:USER_VOTED_VALID]->(d) RETURN uv",
			map[string]interface{}{
				"id":        u.Id,
				"debate_id": t.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().GetByIndex(0), nil
		}

		return nil, res.Err()
	}, nil
}

// Adds a relationship on the user to the debate that they voted as invalid
func (u *User) AddUserVotedInvalidDebateRelationship(t *Tree) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (u:User {id: $id}), (d:Debate {id: $debate_id}) CREATE (u)-[uv:USER_VOTED_INVALID]->(d) RETURN uv",
			map[string]interface{}{
				"id":        u.Id,
				"debate_id": t.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().GetByIndex(0), nil
		}

		return nil, res.Err()
	}, nil
}
