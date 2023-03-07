package types

type Reason int64

const (
	ReasonHarmfulToOthers Reason = iota
	ReasonAbuseOfPlatform
)

type Report struct {
	// The id of the post that has been reported
	ReportedId string

	// The reason for reporting this post
	Reason
}
