package types

type Reference struct {
	// Internal or external reference
	Internal bool `json:"internal"`

	// The amount of times this content was 'trusted'
	Trust int64 `json:"trust"`

	// The amount of times this content was 'distrusted'
	Distrust int64 `json:"distrust"`
}
