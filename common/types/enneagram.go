package types

type PersonalityType int

func (p PersonalityType) String() string {
	switch p {
	case 0:
		return "Enneagram Type 1"
	case 1:
		return "Enneagram Type 2"
	case 2:
		return "Enneagram Type 3"
	case 3:
		return "Enneagram Type 4"
	case 4:
		return "Enneagram Type 5"
	case 5:
		return "Enneagram Type 6"
	case 6:
		return "Enneagram Type 7"
	case 7:
		return "Enneagram Type 8"
	case 8:
		return "Enneagram Type 9"
	default:
		return "Unknown"
	}
}

type EnneagramData struct {
	PersonalityType PersonalityType `json:"personality_type"`

	Tendencies *ActionTendencies `json:"action_tendency"`
}
