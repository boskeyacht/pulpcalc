package tree

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/baribari2/pulp-calculator/common/types"
	dict "github.com/baribari2/pulp-calculator/dictionary"
	neo "github.com/baribari2/pulp-calculator/neo4j"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/rodaine/table"
)

// TODO: Keep track of comment decay when user changes mind
type Tree struct {
	// The Id of the tree
	Id string `json:"id"`

	// The root of the tree
	Root *types.Node `json:"root"`

	Timestamps []int64 `json:"timestamps"`

	LastScore int64 `json:"last_score"`

	// For every time a score change eqauls zero, increase the constant in the log is mult. by
	InactiveCount int64 `json:"inactive_count"`

	// Map node Id to its children
	Nodes map[string]*types.Node `json:"nodes"`

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

func (t *Tree) InsertNode(node *types.Node) {
	t.Lock()
	defer t.Unlock()

	if t.Nodes == nil {
		t.Nodes = make(map[string]*types.Node)
	}

	t.Nodes[string(len(t.Nodes))] = node

	if node.ParentId == "" {
		t.Root = node
		return
	}

	parent, ok := t.Nodes[node.ParentId]
	if !ok {
		return
	}

	parent.Children = append(parent.Children, node)
}

// Implements DFS to calculate the score of each node starting from the root node, and returning the
// score of the root node.
// The root node of a tree is any user action (comment, vote, etc).
//
// This function can also be used to calculate the score of a single node (like a vote without content),
// by passing a tree that contains a single node as the root.
func Calculate(cfg *types.Config /*tree *Tree,*/, node *types.Node) (int, error) {
	if node == nil {
		return 0, nil
	}

	// If this node has no children, it is a leaf node. In that case, return the score of this node.
	// The score of this node may involve calculations regarding confidence, and other characteristics like
	// word count, common words, links, etc.
	if len(node.GetChildren()) == 0 {
		s, err := CalculateScore(cfg, node)
		if err != nil {
			log.Println(err)
			return 0, nil
		}

		return s, nil
	}

	// If the node does have children, traverse through them and calculate their scores.
	var sum int
	for _, node := range node.GetChildren() {
		s, err := Calculate(cfg, node)
		if err != nil {
			log.Println(err)
			return 0, nil
		}

		sum += s
	}

	return sum, nil
}

// Calculates the score of a node, based on its action, content, vote, and confidence.
func CalculateScore(cfg *types.Config, node *types.Node) (int, error) {
	// Set the score equal to the base points of the action
	var score int = int(node.Action.BasePoints())

	if node.Score != 0 {
		score += int(node.Score)
	}

	// If the action is a vote w/o content it's a constant - return the base value
	if node.Action == types.ValidVote || node.Action == types.InvalidVote || node.Action == types.AbstainVote {
		score += int(node.Action.BasePoints())

		return score, nil
	}

	// If the action contians content, then calculate a portion of the score based on the content
	if node.Action == types.CommentResponse ||
		node.Action == types.CommentReply ||
		node.Action == types.ValidVoteWithContent ||
		node.Action == types.InvalidVoteWithContent {

		s, c, err := dict.CountCorrectAndCommonWords(cfg, node.Content)
		if err != nil {
			return 0, err
		}

		// The amount of correct words (positive characteristic) - the amount of common words (negative characteristic)
		score += (s - c)
	}

	// If the action has any votes, then calculate a portion of the score based on the votes
	for _, vote := range node.Engagements.Votes {
		switch vote {
		case types.ValidVoteType:
			score += int(types.ValidVote.BasePoints())

		case types.InvalidVoteType:
			score += int(types.InvalidVote.BasePoints())

		case types.AbstainVoteType:
			score += int(types.AbstainVote.BasePoints())

		default:
			return 0, errors.New(fmt.Sprintf("Invalid vote type: %v", vote))
		}
	}

	// Calculate a portion of the score based on the confidence
	score += int(node.Confidence * 100)

	// If the action has any references, then calculate a portion of the score based on the references

	return score, nil
}

// Simulate a thread until the given end time, with the given frequency.
// The lower the number the higher the frequency of comments.
func SimulateThread(cfg *types.Config, line *charts.Line, users int64, tick time.Duration, endTime time.Duration, freq int64) (*Tree, table.Table, table.Table, error) {
	// ticker := randtick.NewRandTickN(2)
	tree := &Tree{}
	tree.Nodes = make(map[string]*types.Node)
	ticker := time.NewTicker(tick)
	ctable := table.New("Id", "Action", "Content", "Confidence", "Votes", "Time", "Score")
	ttable := table.New("Id", "Score", "Time")
	data := []opts.LineData{}
	stop := make(chan bool)
	// Creates a new tree for insertion into neo4j
	t := neo.NewTree()
	tx, err := t.Create()
	if err != nil {
		return nil, nil, nil, err
	}

	res, err := cfg.Neo4j.WriteTransaction(tx)
	if err != nil {
		return nil, nil, nil, err
	}

	fmt.Printf("RES: %v", res)

	tc, err := dict.Gibber(25)
	if err != nil {
		return nil, nil, nil, err
	}

	// Create the root node
	tree.Root = &types.Node{
		Id:         t.Id,
		Action:     types.CommentResponse,
		Content:    tc,
		Confidence: rand.Float64(),
		Timestamp:  time.Now().Unix(),
		Engagements: types.Engagements{
			Votes: FillAllVotes(rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)),
		},
		Children: []*types.Node{},
	}

	tree.Voters = int64(len(tree.Root.Engagements.Votes))

	score, err := CalculateScore(cfg, tree.Root)
	if err != nil {
		return nil, nil, nil, err
	}

	tree.Root.Score = int64(score)

	// The main simulation goroutine
	score, err = run(cfg, score, tree, ticker, tick, ttable, data, stop)
	if err != nil {
		return nil, nil, nil, err
	}

	tree.Root.Score = int64(score)

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

	// Sleep until simulation is done
	time.Sleep(endTime)
	ticker.Stop()
	stop <- true

	// Add rows to the table
	for _, node := range tree.Root.Children {
		ctable.AddRow(node.Id, node.Action, node.Content[:30]+"...", node.Confidence, len(node.Engagements.Votes), node.Timestamp, node.Score)
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

	t.Timestamps = tree.Timestamps
	t.Topic = tree.Topic
	t.Category = tree.Category
	t.RegisteredSpeakers = tree.RegisteredSpeakers
	t.Voters = tree.Voters
	t.Commenters = tree.Commenters
	t.Comments = tree.Comments

	tx, err = t.Update()
	if err != nil {
		return nil, nil, nil, err
	}

	res, err = cfg.Neo4j.WriteTransaction(tx)
	if err != nil {
		return nil, nil, nil, err
	}

	return tree, ctable, ttable, nil
}

// Main run loop of the simulation
func run(cfg *types.Config, score int, tree *Tree, ticker *time.Ticker, tick time.Duration, table table.Table, data []opts.LineData, stopChan chan bool) (int, error) {
	// Adds comments until the ticker runs out
	for {
		select {
		case <-ticker.C:
			tree.Timestamps = append(tree.Timestamps, time.Now().Unix())

			// After five seconds stop increasing the score of the root node
			if time.Now().UnixNano() > time.Unix(tree.Timestamps[0], 0).UnixNano()+tick.Nanoseconds() {
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
					return 0, err
				}

				// Create a new node
				node := &types.Node{
					Id:         fmt.Sprintln(len(tree.Root.Children) + 1),
					Action:     types.CommentResponse,
					Content:    m,
					Confidence: rand.Float64(),
					ParentId:   tree.Root.Id,
					Timestamp:  time.Now().Unix(),
					Engagements: types.Engagements{
						Votes: FillAllVotes(rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)),
					},
					Children: []*types.Node{},
				}

				s, _ := CalculateScore(cfg, node)

				node.Score = int64(s)
				tree.Root.Children = append(tree.Root.Children, node)

				// Adds a new response to neo4j
				r := neo.NewResponse()
				r.Content = node.Content
				r.Confidence = node.Confidence
				r.Score = int64(s)
				r.Timestamp = node.Timestamp
				r.Engagements = node.Engagements

				tx, err := r.Create()
				if err != nil {
					return 0, err
				}

				// Add the response to the database
				_, err = cfg.Neo4j.WriteTransaction(tx)
				if err != nil {
					return 0, err
				}

				tree.Commenters++

				user := types.NewUser("")
				user.Responses = append(user.Responses, node)
				user.Debates = append(user.Debates, tree.Root.Id)

				// log.Println("User", user.Id, "commented on thread", tree.Root.Id, "with content", node.Content, "and score", node.Score, "and confidence", node.Confidence, "and votes", node.Engagements.Votes, "and timestamp", node.Timestamp)

				u := neo.NewUser()
				tx, err = u.Create()
				if err != nil {
					return 0, err
				}

				// Add the user to the database
				_, err = cfg.Neo4j.WriteTransaction(tx)
				if err != nil {
					return 0, err
				}

				tx, err = u.AddUserResponseRelationship(r)
				if err != nil {
					return 0, err
				}

				_, err = cfg.Neo4j.WriteTransaction(tx)
				if err != nil {
					return 0, err
				}

				tx, err = u.AddUserDebateRelationship(&neo.Tree{Id: tree.Id})
				if err != nil {
					return 0, err
				}

				_, err = cfg.Neo4j.WriteTransaction(tx)
				if err != nil {
					return 0, err
				}

				// Calculate the score of the thread
				score, err = Calculate(cfg, tree.Root)
				if err != nil {
					return 0, err
				}

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
			return score, nil
		}
	}
}
