package types

type SimulationType int

const (
	Enneagram SimulationType = iota
	Age
)

func (s SimulationType) String() string {
	switch s {
	case Enneagram:
		return "Enneagram"
	case Age:
		return "Age"
	default:
		return "Unknown"
	}
}
