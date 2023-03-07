package neo4j

import (
	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

type Response struct {
	Id string `json:"id"`

	Content string `json:"message"`

	Confidence float64 `json:"confidence"`

	Score int64 `json:"score"`

	Timestamp int64 `json:"timestamp"`

	Engagements *types.Engagements
}

func NewResponse(id string, content string, confidence float64, score int64, timestamp int64, engagements *types.Engagements) *Response {
	return &Response{
		Id:          id,
		Content:     content,
		Confidence:  confidence,
		Score:       score,
		Timestamp:   timestamp,
		Engagements: engagements,
	}
}

func NewResponseDefault() *Response {
	return &Response{
		Id:         "",
		Content:    "",
		Confidence: 0,
		Score:      0,
		Timestamp:  0,
	}
}

func (r *Response) Create() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		// TODO: engagements: $engagements
		res, err := tx.Run("CREATE (r:Response {id: $id, content: $content, confidence: $confidence, score: $score, timestamp: $timestamp}) RETURN r.id",
			map[string]interface{}{
				"id":         uuid.New().String(),
				"content":    r.Content,
				"confidence": r.Confidence,
				"score":      r.Score,
				"timestamp":  r.Timestamp,
				// "engagements": r.Engagements,
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

func (r *Response) Update() (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		// TODO: engagements: $engagements
		res, err := tx.Run("MATCH (r:Response {id: $id}) SET r.content = $content, r.confidence = $confidence, r.score = $score, r.timestamp = $timestamp RETURN r",
			map[string]interface{}{
				"id":         uuid.New().String(),
				"content":    r.Content,
				"confidence": r.Confidence,
				"score":      r.Score,
				"timestamp":  r.Timestamp,
				// "engagements": r.Engagements,
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
			return res.Record().Values[0], nil
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

		rec, err := res.Single()
		if err != nil {
			return nil, err
		}

		response := rec.Values[0].(dbtype.Node).Props

		return &Response{
			Id:         response["id"].(string),
			Content:    response["content"].(string),
			Confidence: response["confidence"].(float64),
			Score:      response["score"].(int64),
			Timestamp:  response["timestamp"].(int64),
		}, nil
	}, nil
}

func (r *Response) AddResponseOnDebate(t *Tree) (neo4j.TransactionWork, error) {
	return func(tx neo4j.Transaction) (interface{}, error) {
		res, err := tx.Run("MATCH (t:Tree {id: $debateId}), (r:Response {id: $responseId}) CREATE (d)-[:RESPONSE_TO_DEBATE]->(r)",
			map[string]interface{}{
				"debateId":   t.Id,
				"responseId": r.Id,
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
