package dictionary

import (
	"strings"

	"github.com/baribari2/pulp-calculator/common/types"
)

var (
	COMMON_WORDS = map[string]bool{
		"the":     true,
		"be":      true,
		"to":      true,
		"of":      true,
		"and":     true,
		"a":       true,
		"in":      true,
		"that":    true,
		"have":    true,
		"I":       true,
		"it":      true,
		"for":     true,
		"not":     true,
		"on":      true,
		"he":      true,
		"as":      true,
		"you":     true,
		"do":      true,
		"at":      true,
		"this":    true,
		"but":     true,
		"his":     true,
		"by":      true,
		"from":    true,
		"they":    true,
		"we":      true,
		"say":     true,
		"her":     true,
		"she":     true,
		"or":      true,
		"an":      true,
		"will":    true,
		"my":      true,
		"one":     true,
		"all":     true,
		"would":   true,
		"there":   true,
		"their":   true,
		"what":    true,
		"so":      true,
		"up":      true,
		"out":     true,
		"if":      true,
		"about":   true,
		"who":     true,
		"get":     true,
		"which":   true,
		"go":      true,
		"me":      true,
		"when":    true,
		"make":    true,
		"can":     true,
		"like":    true,
		"time":    true,
		"no":      true,
		"just":    true,
		"him":     true,
		"know":    true,
		"take":    true,
		"people":  true,
		"into":    true,
		"year":    true,
		"your":    true,
		"good":    true,
		"some":    true,
		"could":   true,
		"them":    true,
		"see":     true,
		"other":   true,
		"than":    true,
		"then":    true,
		"now":     true,
		"look":    true,
		"only":    true,
		"come":    true,
		"its":     true,
		"over":    true,
		"think":   true,
		"also":    true,
		"back":    true,
		"after":   true,
		"use":     true,
		"two":     true,
		"how":     true,
		"our":     true,
		"work":    true,
		"first":   true,
		"well":    true,
		"way":     true,
		"even":    true,
		"new":     true,
		"want":    true,
		"because": true,
		"any":     true,
		"these":   true,
		"give":    true,
		"day":     true,
		"most":    true,
		"us":      true,
	}
)

// Counts and returns the number of words in the english dictionary
//
// TODO: Filter out punctuation
func CountCorrectWords(cfg *types.Config, msg string) (int, error) {
	var correctWords int

	for _, word := range strings.Split(msg, " ") {
		// Clean the word in preparation for the query
		word := FilterPunctuation(word)

		// Clean the word in preparation for the query
		def, err := cfg.DictServer.Define("!", word)
		if err != nil {
			return 0, err
		}

		if len(def) > 0 {
			correctWords++
		}
	}

	return correctWords, nil
}

// Counts and returns both words in the common word map, and words in the english dictionary
func CountCorrectAndCommonWords(cfg *types.Config, msg string) (correctWords, commonWords int, err error) {
	for _, word := range strings.Split(msg, " ") {
		// Clean the word in preparation for the query
		word := FilterPunctuation(word)

		// Clean the word in preparation for the query
		def, err := cfg.DictServer.Define("!", word)
		if err != nil {
			return 0, 0, err
		}

		if len(def) > 0 {
			correctWords++
		}

		// Check if the word exits in the common word map
		if _, exists := COMMON_WORDS[word]; exists {
			commonWords++
		}
	}

	return correctWords, commonWords, nil
}

// Filters out punctuation from a string
func FilterPunctuation(msg string) string {
	t := strings.ReplaceAll(msg, ".", "")
	t = strings.ReplaceAll(t, ",", "")
	t = strings.ReplaceAll(t, "!", "")
	t = strings.ReplaceAll(t, "?", "")
	t = strings.ReplaceAll(t, ":", "")
	t = strings.ReplaceAll(t, ";", "")
	t = strings.ReplaceAll(t, "\"", "")
	t = strings.ReplaceAll(t, "'", "")
	t = strings.ReplaceAll(t, "(", "")
	t = strings.ReplaceAll(t, ")", "")
	t = strings.ReplaceAll(t, "[", "")
	t = strings.ReplaceAll(t, "]", "")
	t = strings.ReplaceAll(t, "{", "")
	t = strings.ReplaceAll(t, "}", "")
	t = strings.ReplaceAll(t, "-", "")
	t = strings.ReplaceAll(t, "_", "")
	t = strings.ReplaceAll(t, "=", "")
	t = strings.ReplaceAll(t, "+", "")
	t = strings.ReplaceAll(t, "*", "")
	t = strings.ReplaceAll(t, "/", "")
	t = strings.ReplaceAll(t, "\\", "")
	t = strings.ReplaceAll(t, "|", "")
	t = strings.ReplaceAll(t, "`", "")
	t = strings.ReplaceAll(t, "~", "")
	t = strings.ReplaceAll(t, "@", "")
	t = strings.ReplaceAll(t, "#", "")
	t = strings.ReplaceAll(t, "$", "")
	t = strings.ReplaceAll(t, "%", "")
	t = strings.ReplaceAll(t, "^", "")
	t = strings.ReplaceAll(t, "&", "")
	t = strings.ReplaceAll(t, "<", "")
	t = strings.ReplaceAll(t, ">", "")

	return t
}
