package types

const (
	RelevanceMultiplier     = 10
	SoundnessMultiplier     = 8
	StatsBonus              = 50
	ReferencesBonus         = 40
	GrammarMistakesPenalty  = 30
	SpellingMistakesPenalty = 20
	LengthBonus             = 10
	MasteryVocabBonus       = 5
)

type Attributes struct {
	// The relevance 'score' of the content
	Relevance int64 `json:"relevance"`

	// The soundness 'score' of the content
	Soundness int64 `json:"soundness"`

	// Whether or not stats were included
	StatsIncluded bool `json:"stats_included"`

	// The amount of references (links/files)
	References int64 `json:"references"`

	// The amount of grammar mistakes
	GrammarMistakes int64 `json:"grammar"`

	// The amount of spelling mistakes
	SpellingMistakes int64 `json:"spelling"`

	// The length of the content
	Length int64 `json:"length"`

	// The amount of times this content used a word that signifies mastery in the subject
	MasteryVocab int64 `json:"mastery_vocab"`
}

func NewAttributes(relevance int64, soundness int64, statsIncluded bool, references int64, grammarMistakes int64, spellingMistakes int64, length int64, masteryVocab int64) *Attributes {
	return &Attributes{
		Relevance:        relevance,
		Soundness:        soundness,
		StatsIncluded:    statsIncluded,
		References:       references,
		GrammarMistakes:  grammarMistakes,
		SpellingMistakes: spellingMistakes,
		Length:           length,
		MasteryVocab:     masteryVocab,
	}
}

func NewAttributesDefault() *Attributes {
	return &Attributes{
		Relevance:        0,
		Soundness:        0,
		StatsIncluded:    false,
		References:       0,
		GrammarMistakes:  0,
		SpellingMistakes: 0,
		Length:           0,
		MasteryVocab:     0,
	}
}
