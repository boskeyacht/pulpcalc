package neo4j

import (
	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Response struct {
	Id int `json:"id"`

	Content string `json:"message"`

	Confidence float64 `json:"confidence"`

	Score int64 `json:"score"`

	Timestamp int64 `json:"timestamp"`

	types.Engagements
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Create() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("CREATE (r:Response {id: $id, content: $content, r.confidence = $confidence, score: $score, timestamp: $timestamp, engagements: $engagements}) RETURN r",
			map[string]interface{}{
				"id":          uuid.New().String(),
				"content":     r.Content,
				"confidence":  r.Confidence,
				"score":       r.Score,
				"timestamp":   r.Timestamp,
				"engagements": r.Engagements,
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

func (r *Response) Update() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (r:Response {id: $id}) SET r.content = $content, r.confidence = $confidence, r.score = $score, r.timestamp = $timestamp, r.engagements = $engagements RETURN r",
			map[string]interface{}{
				"id":          uuid.New().String(),
				"content":     r.Content,
				"confidence":  r.Confidence,
				"score":       r.Score,
				"timestamp":   r.Timestamp,
				"engagements": r.Engagements,
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

func (r *Response) Delete() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (r:Response {id: $id}) DELETE r",
			map[string]interface{}{
				"id": r.Id,
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

func (r *Response) GetResponse() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (r:Response {id: $id}) RETURN r",
			map[string]interface{}{
				"id": r.Id,
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
