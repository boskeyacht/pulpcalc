package types

type Reference struct {
	// Internal or external reference
	Internal bool `json:"internal"`

	// The amount of times this content was 'trusted'
	Trust int64 `json:"trust"`

	// The amount of times this content was 'distrusted'
	Distrust int64 `json:"distrust"`

	Content string `json:"content"`
}

func NewReference(internal bool, trust int64, distrust int64, content string) *Reference {
	return &Reference{
		Internal: internal,
		Trust:    trust,
		Distrust: distrust,
		Content:  content,
	}
}

func NewReferenceDefault() *Reference {
	return &Reference{
		Internal: false,
		Trust:    0,
		Distrust: 0,
		Content:  "",
	}
}
