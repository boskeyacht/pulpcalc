package neo4j

import (
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

type Tree struct {
	// Id of the user in neo4j
	Id string `json:"id"`

	// Timestamps of the simulation
	Timestamps []int64 `json:"timestamps"`

	// Topic of the debate
	Topic string `json:"topic"`

	// Category of the debate
	Category string `json:"category"`

	// Number of registered speakers
	RegisteredSpeakers int64 `json:"registered_speakers"`

	// Number of voters
	Voters int64 `json:"voters"`

	// Number of commenters
	Commenters int64 `json:"commenters"`

	// Number of comments
	Comments int64 `json:"comments"`
}

func NewTree(id string, timestamps []int64, topic string, category string, registeredSpeakers int64, voters int64, commenters int64, comments int64) *Tree {
	return &Tree{
		Id:                 id,
		Timestamps:         timestamps,
		Topic:              topic,
		Category:           category,
		RegisteredSpeakers: registeredSpeakers,
		Voters:             voters,
		Commenters:         commenters,
		Comments:           comments,
	}
}

func NewTreeDefault() *Tree {
	return &Tree{
		Id:                 "",
		Timestamps:         []int64{},
		Topic:              "",
		Category:           "",
		RegisteredSpeakers: 0,
		Voters:             0,
		Commenters:         0,
		Comments:           0,
	}
}

// Creates a new debate/tree in neo4j
func (t *Tree) Create() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("CREATE (t:Tree {id: $id, timestamps: $timestamps, topic: $topic, category: $category, registered_speakers: $registered_speakers, voters: $voters, commenters: $commenters, comments: $comments}) RETURN t.id",
			map[string]interface{}{
				"id":                  uuid.New().String(),
				"timestamps":          t.Timestamps,
				"topic":               t.Topic,
				"category":            t.Category,
				"registered_speakers": t.RegisteredSpeakers,
				"voters":              t.Voters,
				"commenters":          t.Commenters,
				"comments":            t.Comments,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().Values[0], nil
		}

		return nil, res.Err()
	}, nil
}

// Retrieves a debate/tree from neo4j
func (t *Tree) GetTree() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}) RETURN t",
			map[string]interface{}{
				"id": t.Id,
			})

		if err != nil {
			return nil, err
		}

		rec, err := res.Single()
		if err != nil {
			return nil, err
		}

		tree := rec.Values[0].(dbtype.Node).Props

		return &Tree{
			Id:                 tree["id"].(string),
			Timestamps:         tree["timestamps"].([]int64),
			Topic:              tree["topic"].(string),
			Category:           tree["category"].(string),
			RegisteredSpeakers: tree["registered_speakers"].(int64),
			Voters:             tree["voters"].(int64),
			Commenters:         tree["commenters"].(int64),
			Comments:           tree["comments"].(int64),
		}, nil
	}, nil
}

// Deletes a debate/tree from neo4j
func (t *Tree) Delete() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}) DETACH DELETE t",
			map[string]interface{}{
				"id": t.Id,
			})

		if err != nil {
			return nil, err
		}

		return res, nil
	}, nil
}

// Updates a debate/tree in neo4j
func (t *Tree) Update() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}) SET t.timestamps = $timestamps, t.topic = $topic, t.category = $category, t.registered_speakers = $registered_speakers, t.voters = $voters, t.commenters = $commenters, t.comments = $comments RETURN t",
			map[string]interface{}{
				"id":                  t.Id,
				"timestamps":          t.Timestamps,
				"topic":               t.Topic,
				"category":            t.Category,
				"registered_speakers": t.RegisteredSpeakers,
				"voters":              t.Voters,
				"commenters":          t.Commenters,
				"comments":            t.Comments,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().Values[0], nil
		}

		return nil, res.Err()
	}, nil
}

// Increases the number of voters in a debate/tree
func (t *Tree) IncreaseVoterCount() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}) SET t.voters = t.voters + 1 RETURN t",
			map[string]interface{}{
				"id": t.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().Values[0], nil
		}

		return nil, res.Err()
	}, nil
}

// Decreases the number of voters in a debate/tree
func (t *Tree) DecreaseVoterCount() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}) SET t.voters = t.voters - 1 RETURN t",
			map[string]interface{}{
				"id": t.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().Values[0], nil
		}

		return nil, res.Err()
	}, nil
}

// Sets the number of voters in a debate/tree
func (t *Tree) SetVoterCount(count int64) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}) SET t.voters = $count RETURN t",
			map[string]interface{}{
				"id":    t.Id,
				"count": count,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().Values[0], nil
		}

		return nil, res.Err()
	}, nil
}

// Increases the number of commenters in a debate/tree
func (t *Tree) IncreaseCommenterCount() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}) SET t.commenters = t.commenters + 1 RETURN t",
			map[string]interface{}{
				"id": t.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().Values[0], nil
		}

		return nil, res.Err()
	}, nil
}

// Decreases the number of commenters in a debate/tree
func (t *Tree) DecreaseCommenterCount() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}) SET t.commenters = t.commenters - 1 RETURN t",
			map[string]interface{}{
				"id": t.Id,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().Values[0], nil
		}

		return nil, res.Err()
	}, nil
}

// Sets the number of commenters in a debate/tree
func (t *Tree) SetCommenterCount(count int64) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}) SET t.commenters = $count RETURN t",
			map[string]interface{}{
				"id":    t.Id,
				"count": count,
			})

		if err != nil {
			return nil, err
		}

		if res.Next() {
			return res.Record().Values[0], nil
		}

		return nil, res.Err()
	}, nil
}

// Adds a relationship on the debate/tree to a registered speaker
func (t *Tree) AddRegisteredSpeakerRelationship(user *User) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}), (u:User {id: $user_id}) CREATE (t)-[:REGISTERED_SPEAKER]->(u)",
			map[string]interface{}{
				"id":      t.Id,
				"user_id": user.Id,
			})

		if err != nil {
			return nil, err
		}

		return res, nil
	}, nil
}

// Adds a relationship on the debate/tree to a user that voted valid on it
func (t *Tree) AddValidVoterRelationship(user *User) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}), (u:User {id: $user_id}) CREATE (t)-[:VOTED_VALID]->(u)",
			map[string]interface{}{
				"id":      t.Id,
				"user_id": user.Id,
			})

		if err != nil {
			return nil, err
		}

		return res, nil
	}, nil
}

// Adds a relationship on the debate/tree to a user that voted invalid on it
func (t *Tree) AddInvalidVoterRelationship(user *User) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}), (u:User {id: $user_id}) CREATE (t)-[:VOTED_INVALID]->(u)",
			map[string]interface{}{
				"id":      t.Id,
				"user_id": user.Id,
			})

		if err != nil {
			return nil, err
		}

		return res, nil
	}, nil
}

// Adds a relationship on the debate/tree to a user that voted abstain on it
func (t *Tree) AddAbstainVoterRelationship(user *User) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}), (u:User {id: $user_id}) CREATE (t)-[:VOTED_ABSTAIN]->(u)",
			map[string]interface{}{
				"id":      t.Id,
				"user_id": user.Id,
			})

		if err != nil {
			return nil, err
		}

		return res, nil
	}, nil
}

// Adds a relationship on the debate/tree to a user that commented on it
func (t *Tree) AddCommenterRelationship(user *User) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $id}), (u:User {id: $user_id}) CREATE (t)-[:COMMENTED]->(u)",
			map[string]interface{}{
				"id":      t.Id,
				"user_id": user.Id,
			})

		if err != nil {
			return nil, err
		}

		return res, nil
	}, nil
}
