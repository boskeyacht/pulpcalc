package simulator

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/baribari2/pulp-calculator/common/types"
	dict "github.com/baribari2/pulp-calculator/dictionary"
	neo "github.com/baribari2/pulp-calculator/neo4j"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rodaine/table"
)

// TODO: Keep track of comment decay when user changes mind
type Debate struct {
	// The Id of the tree
	Id string `json:"id"`

	// The root of the tree
	Root *Response `json:"root"`

	// The timestamps of the simulation, the first time stamp is instantiation
	Timestamps []int64 `json:"timestamps"`

	// The last score of the debate
	LastScore int64 `json:"last_score"`

	// For every time a score change eqauls zero, increase the constant in the log is mult. by
	InactiveCount int64 `json:"inactive_count"`

	// Map node Id to its children
	Responses map[string]*Response `json:"responses"`

	// The topic of the debate
	Topic string `json:"topic"`

	// The category of the debate
	Category string `json:"category"`

	// The number of participants in the debate
	RegisteredSpeakers int64 `json:"registered_speakers"`

	// The number of participants in the debate
	SupportingAudience int64 `json:"supporting_audience"`

	// The number of participants in the debate
	Commenters int64 `json:"commenters"`

	// The number of voters in the debate
	Voters int64 `json:"voters"`

	// The number of inactive participants in the debate
	InactiveParticipants int64 `json:"inactive_participants"`

	// The number of comments in the debate
	Comments int64 `json:"comments"`

	// The duration of the debate
	Duration int64 `json:"duration"`

	sync.Mutex
}

// Implements DFS to calculate the score of each node starting from the root node, and returning the
// score of the root node.
// The root node of a tree is any user action (comment, vote, etc).
//
// This function can also be used to calculate the score of a single node (like a vote without content),
// by passing a tree that contains a single node as the root.
func Calculate(cfg *types.Config /*tree *Tree,*/, response *Response) (int, error) {
	if response == nil {
		return 0, nil
	}

	// If this node has no children, it is a leaf node. In that case, return the score of this node.
	// The score of this node may involve calculations regarding confidence, and other characteristics like
	// word count, common words, links, etc.
	if len(response.GetChildren()) == 0 {
		s := response.CalculateContentAttributesScore(cfg) + response.CalculateEngagementAttributesScore()

		return s, nil
	}

	// If the node does have children, traverse through them and calculate their scores.
	var sum int
	for _, node := range response.GetChildren() {
		s, err := Calculate(cfg, node)
		if err != nil {
			log.Println(err)
			return 0, nil
		}

		sum += s
	}

	return sum, nil
}

// Simulate a thread until the given end time, with the given frequency.
// The lower the number the higher the frequency of comments.
func SimulateThread(cfg *types.Config, line *charts.Line, users int64, tick time.Duration, endTime time.Duration, freq int64) (*Debate, table.Table, table.Table, error) {
	// ticker := randtick.NewRandTickN(2)
	ticker := time.NewTicker(tick)
	ctable := table.New("Id", "Action", "Content", "Confidence", "Votes", "Time", "Score")
	ttable := table.New("Id", "Score", "Time")
	data := []opts.LineData{}
	stop := make(chan bool)
	tree := &Debate{
		Topic:     "Does everyone need a therapist?",
		Category:  "healthcare",
		Responses: make(map[string]*Response),
	}

	// Creates a new tree for insertion into neo4j
	t := neo.NewTreeDefault()
	tx, err := t.Create()
	if err != nil {
		return nil, nil, nil, err
	}

	res, err := cfg.Neo4j.WriteTransaction(tx, neo4j.WithTxTimeout(1*time.Second))
	if err != nil {
		return nil, nil, nil, err
	}

	tree.Id = res.(string)

	tc, err := dict.Gibber(25)
	if err != nil {
		return nil, nil, nil, err
	}

	// Create the root node
	tree.Root = &Response{
		Id:         t.Id,
		Action:     types.CommentResponse,
		Content:    NewContent(0, tc),
		Confidence: rand.Float64(),
		Timestamp:  time.Now().Unix(),
		Engagements: &types.Engagements{
			Reports:   []*types.Report{},
			HideCount: 0,
			Votes:     FillAllVotes(rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)),
		},
		Attributes: types.NewAttributesDefault(),
		Children:   []*Response{},
	}

	tree.Voters = int64(len(tree.Root.Engagements.Votes))

	fmt.Println(tree.Root.RootTimestamp)

	score, err := tree.Root.CalculateScore(cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	tree.Root.Score = int64(score)

	// The main simulation goroutine
	scoreChan := make(chan int64, 1)
	go run(cfg, score, tree, ticker, endTime, ttable, data, stop, scoreChan)

	// Sleep until simulation is done
	time.Sleep(endTime)
	ticker.Stop()
	stop <- true

	tree.Root.Score = <-scoreChan
	close(scoreChan)

	// Calculate the score of the thread one last time
	score, err = Calculate(cfg, tree.Root)
	if err != nil {
		return nil, nil, nil, err
	}

	data = append(data, opts.LineData{Value: score})
	tree.Root.Score = int64(score)
	tree.Root.Timestamp = time.Now().Unix()

	// Add the row to the table
	ttable.AddRow(tree.Root.Id, tree.Root.Score, tree.Root.Timestamp)

	// Add rows to the table
	for _, node := range tree.Root.Children {
		ctable.AddRow(node.Id, node.Action, node.Content.Message[:30]+"...", node.Confidence, len(node.Engagements.Votes), node.Timestamp, node.Score)
	}

	x := []interface{}{}

	end, err := strconv.Atoi(endTime.String()[:len(endTime.String())-1])
	if err != nil {
		return nil, nil, nil, err
	}

	for i := 0; i < end; i++ {
		x = append(x, i)
	}

	line.SetXAxis(x).
		AddSeries("thread score", data).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	// Write the tree to the neo4j
	t.Topic = tree.Topic
	t.Timestamps = tree.Timestamps
	t.Category = tree.Category
	t.RegisteredSpeakers = tree.RegisteredSpeakers
	t.Voters = tree.Voters
	t.Commenters = tree.Commenters
	t.Comments = tree.Comments

	tx, err = t.Update()
	if err != nil {
		return nil, nil, nil, err
	}

	_, err = cfg.Neo4j.WriteTransaction(tx, neo4j.WithTxTimeout(1*time.Second))
	if err != nil {
		return nil, nil, nil, err
	}

	return tree, ctable, ttable, nil
}

// Main run loop of the simulation
func run(cfg *types.Config, score int, tree *Debate, ticker *time.Ticker, endTime time.Duration, table table.Table, data []opts.LineData, stopChan chan bool, scoreChan chan int64) {
	// Adds comments until the ticker runs out
outer:
	for {
		select {
		case <-ticker.C:
			tree.Timestamps = append(tree.Timestamps, time.Now().Unix())

			// After five seconds stop increasing the score of the root node
			if time.Now().UnixNano() > time.Unix(tree.Timestamps[0], 0).UnixNano()+(endTime.Nanoseconds()/2) {
				// If the score hasn't changed...
				if tree.Root.Score == tree.Root.Score {
					tree.InactiveCount++

					d := CalculateDecay(int(tree.InactiveCount))
					if d != 0 {
						tree.Root.Score = int64(float64(tree.Root.Score) * d)
					}

				} else {
					tree.Root.Score = int64(score)
					tree.Root.Timestamp = time.Now().Unix()
				}
			} else {

				// Generate content for the node
				m, err := dict.Gibber(25)
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				// Create a new node
				node := &Response{
					Id:            fmt.Sprintln(len(tree.Root.Children) + 1),
					RootTimestamp: tree.Root.Timestamp,
					Action:        types.CommentResponse,
					Content:       NewContent(0, m),
					Confidence:    rand.Float64(),
					ParentId:      tree.Root.Id,
					Timestamp:     time.Now().Unix(),
					Engagements: &types.Engagements{
						Reports:   []*types.Report{},
						HideCount: 0,
						Votes:     FillAllVotes(rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)),
					},
					Attributes: types.NewAttributesDefault(),
					Children:   []*Response{},
				}

				s := node.CalculateContentAttributesScore(cfg) + node.CalculateEngagementAttributesScore()

				node.Score = int64(s)
				tree.Root.Children = append(tree.Root.Children, node)

				// Adds a new response to neo4j
				r := neo.NewResponse("", node.Content.Message, node.Confidence, int64(s), node.Timestamp, node.Engagements)

				tx, err := r.Create()
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				// Add the response to the database
				res, err := cfg.Neo4j.WriteTransaction(tx, neo4j.WithTxTimeout(1*time.Second))
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				r.Id = res.(string)

				tx, err = r.AddResponseOnDebate(&neo.Tree{Id: tree.Id})
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				_, err = cfg.Neo4j.WriteTransaction(tx, neo4j.WithTxTimeout(1*time.Second))
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				tree.Commenters++

				user := types.NewUser("")
				user.Responses = append(user.Responses, r.Id)
				user.Debates = append(user.Debates, tree.Root.Id)

				u := neo.NewUserDefault()
				tx, err = u.Create()
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				// Add the user to the database
				res, err = cfg.Neo4j.WriteTransaction(tx, neo4j.WithTxTimeout(1*time.Second))
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				u.Id = res.(string)

				tx, err = u.AddUserResponseRelationship(r)
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				_, err = cfg.Neo4j.WriteTransaction(tx, neo4j.WithTxTimeout(1*time.Second))
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				tx, err = u.AddUserDebateRelationship(&neo.Tree{Id: tree.Id})
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				_, err = cfg.Neo4j.WriteTransaction(tx, neo4j.WithTxTimeout(1*time.Second))
				if err != nil {
					fmt.Println(err.Error())

					return
				}

				// Calculate the score of the thread
				// score, err = Calculate(cfg, tree.Root)
				// if err != nil {
				// 	fmt.Println(err.Error())

				// 	return
				// }

				// Calculate the change in score and increment the counter if the change is zero
				if int64(score) == tree.Root.Score {
					tree.InactiveCount++

					d := CalculateDecay(int(tree.InactiveCount))
					score = int(float64(score) * d)

				} else {
					tree.Root.Score = int64(score)
					tree.Root.Timestamp = time.Now().Unix()
				}
			}

			data = append(data, opts.LineData{Value: tree.Root.Score})

			// Add the row to the table
			table.AddRow(tree.Root.Id, tree.Root.Score, tree.Root.Timestamp)

		case <-stopChan:
			scoreChan <- int64(score)
			break outer
		}
	}
}

func CalculateDecay(degree int) float64 {
	return math.Log(float64(degree)) / 3
}
