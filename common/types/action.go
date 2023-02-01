package types

type Action int64

const (
	Topic Action = iota
	CommentResponse
	CommentReply
	ValidVote
	InvalidVote
	AbstainVote
	ValidVoteWithContent
	InvalidVoteWithContent
	TrustedReference
	DistrustedReference
)

// TODO: Rearrange from largest to smallest
func (a Action) BasePoints() int64 {
	switch a {
	case CommentResponse:
		return 10
	case CommentReply:
		return 8
	case ValidVote:
		return 4
	case InvalidVote:
		return 4
	case AbstainVote:
		return 0
	case ValidVoteWithContent:
		return 5
	case InvalidVoteWithContent:
		return 5
	case TrustedReference:
		return 3
	case DistrustedReference:
		return 0
	default:
		return 0
	}
}
