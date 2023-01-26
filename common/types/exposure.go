package types

type Exposure struct {
	// The total amount of time the post has been posted
	Time int64

	// The amount of people this post has reached
	Impressions int64

	// The amount of people who have returned o this psot
	Revisits int64
}
