package types

type SimulateThreadRequest struct {
	Tick      int64 `json:"tick"`
	EndTime   int64 `json:"end_time"`
	Frequency int64 `json:"frequency"`
}

func NewSimulateThreadRequest(tick, endTime, frequency int64) *SimulateThreadRequest {
	return &SimulateThreadRequest{
		Tick:      tick,
		EndTime:   endTime,
		Frequency: frequency,
	}
}

type GetTreeRequest struct {
	Get bool `json:"get"`
}

func NewGetTreeRequest(get bool) *GetTreeRequest {
	return &GetTreeRequest{
		Get: get,
	}
}
