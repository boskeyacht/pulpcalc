package types

type Report struct {
	// The id of the post that has been reported
	ReportedId int

	// The reason for reporting this post
	Reason string
}
