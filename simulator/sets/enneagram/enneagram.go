package enneagram

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/baribari2/pulpcalc/chatgpt"
	"github.com/baribari2/pulpcalc/common/types"
	neo "github.com/baribari2/pulpcalc/neo4j"
	"github.com/baribari2/pulpcalc/simulator"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"mvdan.cc/xurls/v2"
)

type EnneagramSet struct {
	SimulationSize int `json:"simulation_size" yaml:"simulation_size"`

	Users []*types.User `json:"users" yaml:"users"`

	Distribution []float64 `json:"distribution" yaml:"distribution"`

	Duration time.Duration `json:"duration" yaml:"duration"`

	Topic string `json:"topic" yaml:"topic"`

	Category string `json:"category" yaml:"category"`

	Depth int `json:"depth" yaml:"depth"`
}

func NewEnneagramSet(size int, users []*types.User, distribution []float64, duration time.Duration, topic string, category string, depth int) *EnneagramSet {
	return &EnneagramSet{
		SimulationSize: size,
		Users:          users,
		Distribution:   distribution,
		Duration:       duration,
		Topic:          topic,
		Category:       category,
		Depth:          depth,
	}
}

func NewEnneagramSetDefault() *EnneagramSet {
	return &EnneagramSet{
		SimulationSize: 0,
		Users:          nil,
		Distribution:   nil,
		Duration:       0,
		Topic:          "",
		Category:       "",
		Depth:          0,
	}
}

func (e *EnneagramSet) GetSimulationType() types.SimulationType {
	return types.Enneagram
}

func (e *EnneagramSet) GetSimulationSize() int {
	return e.SimulationSize
}

func (e *EnneagramSet) GetUsers() []*types.User {
	return e.Users
}

func (e *EnneagramSet) GetDistribution() []float64 {
	return e.Distribution
}

func (e *EnneagramSet) GetDuration() time.Duration {
	return e.Duration
}

func (e *EnneagramSet) GetTopic() string {
	return e.Topic
}

func (e *EnneagramSet) GetDepth() int {
	return e.Depth
}

func (e *EnneagramSet) GetCategory() string {
	return e.Category
}

func (e *EnneagramSet) RunSimulation(wg *sync.WaitGroup, cfg *types.Config, debateChan chan *simulator.Debate, errChan chan error) {
	ticker := time.NewTicker(time.Second)
	users := []*types.User{}
	stop := make(chan bool)
	debate := &simulator.Debate{
		Topic:     e.Topic,
		Category:  e.Category,
		Responses: make(map[string]*simulator.Response),
	}

	tx, err := neo.NewTree("", []int64{}, debate.Topic, debate.Category, 0, 0, 0, 0).Create()
	if err != nil {
		errChan <- err
		wg.Done()
	}

	res := writeOrPanic(cfg.Neo4j, tx)
	debate.Id = res.(string)

	// Given this set of actions, this enneagram type, and this debate topic,
	// what is this user most likely to do
	r, err := chatgpt.SendChatRequest(cfg.OpenAIClient, strings.Replace(tendencyPrompt, "THIS_TOPIC", e.Topic, 1))
	if err != nil {
		fmt.Println(err.Error())

		errChan <- err
	}

	tendencies := &TendencyResponse{}
	err = json.Unmarshal([]byte(r), tendencies)
	if err != nil {
		fmt.Println(err.Error())

		errChan <- err
	}

	// Here, we assume that the array has a length of 9, for the 9 enneagram types,
	// each index maps to an enneagram type (0 is 1, 1 is 2, etc.)
	for i, dist := range e.Distribution {
		// This multiplication is what creates users based on the distribution
		for j := 0; j < int(float64(e.SimulationSize)*dist); j++ {
			user := &types.User{
				Id:        "",
				Votes:     nil,
				Debates:   nil,
				Responses: nil,
				SetData: map[types.SimulationType]interface{}{
					types.Enneagram: &types.EnneagramData{
						PersonalityType: types.PersonalityType(i),
						Tendencies:      matchUserTendency(i, tendencies),
					},
				},
			}

			users = append(users, user)

			tx, err := neo.NewUser(user.Id, user.Votes, user.Debates, user.SetData).Create()
			if err != nil {
				fmt.Println(err.Error())
				errChan <- err
			}

			res := writeOrPanic(cfg.Neo4j, tx)
			user.Id = res.(string)
		}
	}

	go func(us []*types.User) { // Until time is up generate a random piece of content based on a random chosen user
	outer:
		for {
			select {
			case <-ticker.C:
				randUser := us[rand.Intn(len(users)-1)]

				p := strings.Replace(enneagramContentPrompt, "THIS_TOPIC", e.Topic, 1)
				p = strings.Replace(p, "VALID_VOTE_TENDENCY", fmt.Sprintf("%v", randUser.SetData[types.Enneagram].(*types.EnneagramData).Tendencies.ValidVoteTendency), 1)
				p = strings.Replace(p, "INVALID_VOTE_TENDENCY", fmt.Sprintf("%v", randUser.SetData[types.Enneagram].(*types.EnneagramData).Tendencies.InvalidVoteTendency), 1)
				p = strings.Replace(p, "ABSTAIN_VOTE_TENDENCY", fmt.Sprintf("%v", randUser.SetData[types.Enneagram].(*types.EnneagramData).Tendencies.AbstainVoteTendency), 1)
				p = strings.Replace(p, "REPORT_TENDENCY", fmt.Sprintf("%v", randUser.SetData[types.Enneagram].(*types.EnneagramData).Tendencies.ReportTendency), 1)
				p = strings.Replace(p, "HIDE_TENDENCY", fmt.Sprintf("%v", randUser.SetData[types.Enneagram].(*types.EnneagramData).Tendencies.HideTendency), 1)

				res, err := chatgpt.SendChatRequest(cfg.OpenAIClient, p)
				if err != nil {
					fmt.Println(err.Error())
					errChan <- err
				}

				content := &ContentResponse{}
				err = json.Unmarshal([]byte(res), &content)
				if err != nil {
					fmt.Println(err.Error())
					errChan <- err
				}

				response := &simulator.Response{
					Timestamp: time.Now().Unix(),
					ParentId:  debate.Id,
					// RootTimestamp: debate.Root.Timestamp,
					Attributes:  &types.Attributes{},
					Children:    make([]*simulator.Response, 0),
					Engagements: &types.Engagements{},
					Content:     simulator.NewContent(0, content.Content),
					Confidence:  content.Confidence,
					Score:       types.CommentResponse.BasePoints(),
				}

				debate.Responses[response.Id] = response

				resp := neo.NewResponse("", response.Content.Message, response.Confidence, response.Score, response.Timestamp, response.Engagements)
				userM := neo.NewUser(randUser.Id, nil, nil, nil)

				tx, err := resp.Create()
				if err != nil {
					fmt.Println(err.Error())
					errChan <- err
				}

				res = writeOrPanic(cfg.Neo4j, tx).(string)
				response.Id = res
				resp.Id = res

				rx := xurls.Relaxed()
				urls := rx.FindAllString(response.Content.Message, -1)

				if len(urls) > 0 {
					for _, url := range urls {
						ref := neo.NewReference("", false, 0, 0, url)

						tx, err := ref.Create()
						if err != nil {
							fmt.Println(err.Error())
							errChan <- err
						}

						res := writeOrPanic(cfg.Neo4j, tx).(string)
						ref.Id = res

						tx, err = ref.AddReferenceOnResponse(resp)
						if err != nil {
							fmt.Println(err.Error())
							errChan <- err
						}

						writeOrPanic(cfg.Neo4j, tx)
					}
				}

				tx, err = userM.AddUserResponseRelationship(resp)
				if err != nil {
					fmt.Println(err.Error())
					errChan <- err
				}

				writeOrPanic(cfg.Neo4j, tx)

				// ...
				generateEngagement(cfg, response, us, e.Depth)

				score := response.CalculateContentAttributesScore(cfg) + response.CalculateEngagementAttributesScore()
				response.Score += int64(score)

				debate := neo.NewTree(debate.Id, nil, "", "", 0, 0, 0, 0)
				tx, err = resp.AddResponseOnDebate(debate)
				if err != nil {
					fmt.Println(err.Error())
					errChan <- err
				}

				writeOrPanic(cfg.Neo4j, tx)

				tx, err = userM.AddUserDebateRelationship(debate)
				if err != nil {
					fmt.Println(err.Error())
					errChan <- err
				}

				writeOrPanic(cfg.Neo4j, tx)

				fmt.Println("user id ", randUser.Id)
				fmt.Println("user personality type ", randUser.SetData[types.Enneagram].(*types.EnneagramData).PersonalityType)
				fmt.Println("timestamp ", response.Timestamp)
				fmt.Println("content ", response.Content.Message)
				fmt.Println("score ", response.Score)
				fmt.Println("confidence ", response.Confidence)
				fmt.Println()

			case <-stop:
				break outer
			}
		}
	}(users)

	time.Sleep(e.Duration)
	ticker.Stop()
	stop <- true

	debateChan <- debate
	wg.Done()
}

func generateEngagement(cfg *types.Config, r *simulator.Response, users []*types.User, depth int) {
	engagement := &types.Engagements{}

	if depth == 0 {
		return
	}

	randUser := users[rand.Intn(len(users)-1)]
	um := neo.NewUser(randUser.Id, nil, nil, nil)
	rm := neo.NewResponse(r.Id, "", 0, 0, 0, nil)

	p := strings.Replace(enneagramReplyPrompt, "THIS_CONTENT", r.Content.Message, 1)
	p = strings.Replace(p, "VALID_VOTE_TENDENCY", fmt.Sprintf("%v", randUser.SetData[types.Enneagram].(*types.EnneagramData).Tendencies.ValidVoteTendency), 1)
	p = strings.Replace(p, "INVALID_VOTE_TENDENCY", fmt.Sprintf("%v", randUser.SetData[types.Enneagram].(*types.EnneagramData).Tendencies.InvalidVoteTendency), 1)
	p = strings.Replace(p, "ABSTAIN_VOTE_TENDENCY", fmt.Sprintf("%v", randUser.SetData[types.Enneagram].(*types.EnneagramData).Tendencies.AbstainVoteTendency), 1)
	p = strings.Replace(p, "REPORT_TENDENCY", fmt.Sprintf("%v", randUser.SetData[types.Enneagram].(*types.EnneagramData).Tendencies.ReportTendency), 1)
	p = strings.Replace(p, "HIDE_TENDENCY", fmt.Sprintf("%v", randUser.SetData[types.Enneagram].(*types.EnneagramData).Tendencies.HideTendency), 1)

	res, err := chatgpt.SendChatRequest(cfg.OpenAIClient, p)
	if err != nil {
		return
	}

	content := &ContentResponse{}
	err = json.Unmarshal([]byte(res), &content)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	response := &simulator.Response{
		Timestamp: time.Now().Unix(),
		ParentId:  r.Id,
		// RootTimestamp: debate.Root.Timestamp,
		Attributes:  &types.Attributes{},
		Children:    make([]*simulator.Response, 0),
		Engagements: &types.Engagements{},
		Content:     simulator.NewContent(0, content.Content),
		Confidence:  content.Confidence,
		Score:       types.CommentResponse.BasePoints(),
	}

	response.Score += int64(response.CalculateContentAttributesScore(cfg))

	vv := simulator.FillValidVotes(rand.Intn(len(users) - 1))
	iv := simulator.FillInvalidVotes(rand.Intn((len(users) - 1) - len(vv)))
	av := simulator.FillAbstainVotes(rand.Intn((len(users) - 1) - (len(vv) + len(iv))))

	engagement.Votes = append(engagement.Votes, vv...)
	engagement.Votes = append(engagement.Votes, iv...)
	engagement.Votes = append(engagement.Votes, av...)

	engagement.Reports = simulator.FillReports(rand.Intn(int(float64(len(users)) * 0.03)))

	engagement.Responses = int64(depth)

	fmt.Println("engagement ", engagement)

	r.Engagements = engagement

	rs := neo.NewResponse(response.Id, response.Content.Message, response.Confidence, response.Score, response.Timestamp, response.Engagements)
	tx, err := rs.Create()
	if err != nil {
		return
	}

	res = writeOrPanic(cfg.Neo4j, tx).(string)
	response.Id = res
	rs.Id = res

	rx := xurls.Relaxed()
	urls := rx.FindAllString(response.Content.Message, -1)

	if len(urls) > 0 {
		for _, url := range urls {
			ref := neo.NewReference("", false, 0, 0, url)

			tx, err := ref.Create()
			if err != nil {
				return
			}

			res := writeOrPanic(cfg.Neo4j, tx).(string)
			ref.Id = res

			tx, err = ref.AddReferenceOnResponse(rs)
			if err != nil {
				return
			}

			writeOrPanic(cfg.Neo4j, tx)
		}
	}

	fmt.Println("response ", response)

	tx, err = um.AddUserResponseRelationship(rs)
	if err != nil {
		return
	}

	writeOrPanic(cfg.Neo4j, tx)

	tx, err = rm.AddReplyOnResponse(rs)
	if err != nil {
		return
	}

	writeOrPanic(cfg.Neo4j, tx)

	generateEngagement(cfg, response, users, depth-1)
}

func matchUserTendency(enneagramType int, tendencies *TendencyResponse) *types.ActionTendencies {
	switch enneagramType {
	case 0:
		return tendencies.Type1
	case 1:
		return tendencies.Type2
	case 2:
		return tendencies.Type3
	case 3:
		return tendencies.Type4
	case 4:
		return tendencies.Type5
	case 5:
		return tendencies.Type6
	case 6:
		return tendencies.Type7
	case 7:
		return tendencies.Type8
	case 8:
		return tendencies.Type9
	default:
		return nil
	}
}

// Having to handle the neo4j.WriteTransaction error everywhere gets clunky, and would rather not ignore the error since
// the neo4j nodes that are created are dependent upon each other. Might be better to panic anyways.
func writeOrPanic(session neo4j.Session, tx neo4j.TransactionWork) interface{} {
	res, err := session.WriteTransaction(tx, neo4j.WithTxTimeout(1*time.Second))
	if err != nil {
		panic(err)
	}

	return res
}
