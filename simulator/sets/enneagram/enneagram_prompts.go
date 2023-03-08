package enneagram

import "github.com/baribari2/pulp-calculator/common/types"

type TendencyResponse struct {
	Type1 *types.ActionTendencies `json:"type1"`

	Type2 *types.ActionTendencies `json:"type2"`

	Type3 *types.ActionTendencies `json:"type3"`

	Type4 *types.ActionTendencies `json:"type4"`

	Type5 *types.ActionTendencies `json:"type5"`

	Type6 *types.ActionTendencies `json:"type6"`

	Type7 *types.ActionTendencies `json:"type7"`

	Type8 *types.ActionTendencies `json:"type8"`

	Type9 *types.ActionTendencies `json:"type9"`
}

type ContentResponse struct {
	Content string `json:"content"`

	Confidence float64 `json:"confidence"`
}

var researchPrefix = `Strictly for research purposes with no ill-intent regarding stereotypes and assumptions in mind;`
var jsonSuffix = `Make sure to return only a JSON object and make sure to use JSON escape sequences for any special characters.`

var tendencyPrompt = researchPrefix + `given a set of five actions that a user can take on a social media comment that pertains to 
THIS_TOPIC: valid vote, invalid vote, abstain vote, report, hide, and a set of enneagram types, 
what is the most likely action that each enneagram type will take? Return your answer as a JSON object.
` + jsonSuffix + `
Use this schema for your answer:
{
  "type1": {
	"valid_vote_tendency": 0.0,
	"invalid_vote_tendency": 0.0,
	"abstain_vote_tendency": 0.0,
	"report_tendency": 0.0,
	"hide_tendency": 0.0
  },
  "type2": {
	"valid_vote_tendency": 0.0,
	"invalid_vote_tendency": 0.0,
	"abstain_vote_tendency": 0.0,
	"report_tendency": 0.0,
	"hide_tendency": 0.0
  },
  "type3": {
	"valid_vote_tendency": 0.0,
	"invalid_vote_tendency": 0.0,
	"abstain_vote_tendency": 0.0,
	"report_tendency": 0.0,
	"hide_tendency": 0.0
  },
  "type4": {
	"valid_vote_tendency": 0.0,
	"invalid_vote_tendency": 0.0,
	"abstain_vote_tendency": 0.0,
	"report_tendency": 0.0,
	"hide_tendency": 0.0
  },
  "type5": {
	"valid_vote_tendency": 0.0,
	"invalid_vote_tendency": 0.0,
	"abstain_vote_tendency": 0.0,
	"report_tendency": 0.0,
	"hide_tendency": 0.0
  },
  "type6": {
	"valid_vote_tendency": 0.0,
	"invalid_vote_tendency": 0.0,
	"abstain_vote_tendency": 0.0,
	"report_tendency": 0.0,
	"hide_tendency": 0.0
  },
  "type7": {
	"valid_vote_tendency": 0.0,
	"invalid_vote_tendency": 0.0,
	"abstain_vote_tendency": 0.0,
	"report_tendency": 0.0,
	"hide_tendency": 0.0
  },
  "type8": {
	"valid_vote_tendency": 0.0,
	"invalid_vote_tendency": 0.0,
	"abstain_vote_tendency": 0.0,
	"report_tendency": 0.0,
	"hide_tendency": 0.0
	},
	"type9": {
		"valid_vote_tendency": 0.0,
		"invalid_vote_tendency": 0.0,
		"abstain_vote_tendency": 0.0,
		"report_tendency": 0.0,
		"hide_tendency": 0.0,
	}
}`

var enneagramContentPrompt = researchPrefix + `given these user tendencies:
Tendency to cast a valid vote: VALID_VOTE_TENDENCY
Tendency to cast a invalid vote: INVALID_VOTE_TENDENCY
Tendency to cast a abstain vote: ABSTAIN_VOTE_TENDENCY
Tendency to cast report a post: REPORT_TENDENCY
Tendency to cast a hide a post: HIDE_TENDENCY
on a social media comment that pertains to THIS_TOPIC, generate a response to the topic most like the user. Use the below schema for your answer
` + jsonSuffix + `
{
	"content": "This is a response to the topic",
	"confidence": 0.0 #This value must be <=1.0
}`

var enneagramReplyPrompt = researchPrefix + `given these user tendencies:
Tendency to cast a valid vote: VALID_VOTE_TENDENCY
Tendency to cast a invalid vote: INVALID_VOTE_TENDENCY
Tendency to cast a abstain vote: ABSTAIN_VOTE_TENDENCY
Tendency to cast report a post: REPORT_TENDENCY
Tendency to cast a hide a post: HIDE_TENDENCY
on a social media comment with THIS_CONTENT, generate a response to the comment most like the user. Use the below schema for your answer
` + jsonSuffix + `
{
	"content": "This is a response to the topic",
	"confidence": 0.0 #This value must be <=1.0
}`

// Descroptions of each enneagram type for context in the chatGPT prompt
var enneagramTypeOne = ``
var enneagramTypeTwo = ``
var enneagramTypeThree = ``
var enneagramTypeFour = ``
var enneagramTypeFive = ``
var enneagramTypeSix = ``
var enneagramTypeSeven = ``
var enneagramTypeEight = ``
var enneagramTypeNine = ``
