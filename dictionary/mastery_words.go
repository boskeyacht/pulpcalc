package dictionary

type MasteryWordMap map[string]bool

var (
	politicsMasteryWords = MasteryWordMap{
		"abortion": true,
		"biden":    true,
		"harris":   true,
		"trump":    true,
		"feminism": true,
		"racism":   true,
	}

	MasteryWords = map[string]MasteryWordMap{
		"politics": politicsMasteryWords,
	}
)

func IsMasteryWord(category string, word string) bool {
	if masteryWords, ok := MasteryWords[category]; ok {
		if _, ok := masteryWords[word]; ok {
			return true
		}
	}

	return false
}
