package neo4j

import (
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

type Reference struct {
	// The reference Id
	Id string `json:"id"`

	// Internal or external reference
	Internal bool `json:"internal"`

	// The amount of times this content was 'trusted'
	Trust int64 `json:"trust"`

	// The amount of times this content was 'distrusted'
	Distrust int64 `json:"distrust"`

	// The content of the reference (most likely a link)
	Content string `json:"content"`
}

func NewReference(id string, internal bool, trust int64, distrust int64, content string) *Reference {
	return &Reference{
		Id:       id,
		Internal: internal,
		Trust:    trust,
		Distrust: distrust,
		Content:  content,
	}
}

func NewReferenceDefault() *Reference {
	return &Reference{
		Id:       "",
		Internal: false,
		Trust:    0,
		Distrust: 0,
		Content:  "",
	}
}

func (r *Reference) Create() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("CREATE (r:Reference {id: $id, internal: $internal, trust: $trust, distrust: $distrust, content: $content}) RETURN r.id", map[string]interface{}{
			"id":       uuid.New().String(),
			"internal": r.Internal,
			"trust":    r.Trust,
			"distrust": r.Distrust,
			"content":  r.Content,
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

func (r *Reference) Update() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (r:Reference {id: $id}) SET r.internal = $internal, r.trust = $trust, r.distrust = $distrust, r.content = $content RETURN r", map[string]interface{}{
			"id":       r.Id,
			"internal": r.Internal,
			"trust":    r.Trust,
			"distrust": r.Distrust,
			"content":  r.Content,
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

func (r *Reference) Delete() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (r:Reference {id: $id}) DELETE r", map[string]interface{}{
			"id": r.Id,
		})

		if err != nil {
			return nil, err
		}

		return res, nil
	}, nil
}

func (r *Reference) GetReference() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (r:Reference {id: $id}) RETURN r", map[string]interface{}{
			"id": r.Id,
		})

		if err != nil {
			return nil, err
		}

		rec, err := res.Single()
		if err != nil {
			return nil, err
		}

		reference := rec.Values[0].(dbtype.Node).Props

		return &Reference{
			Id:       reference["id"].(string),
			Internal: reference["internal"].(bool),
			Trust:    reference["trust"].(int64),
			Distrust: reference["distrust"].(int64),
			Content:  reference["content"].(string),
		}, nil
	}, nil
}

func (r *Reference) AddReferenceOnResponse(response *Response) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (r:Reference {id: $id}) MATCH (resp:Response {id: $responseId}) CREATE (resp)-[:HAS_REFERENCE]->(r) RETURN r", map[string]interface{}{
			"id":         r.Id,
			"responseId": response.Id,
		})

		if err != nil {
			return nil, err
		}

		return res, nil
	}, nil
}
