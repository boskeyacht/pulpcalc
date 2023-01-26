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
		return 20
	case CommentReply:
		return 16
	case ValidVote:
		return 7
	case InvalidVote:
		return 7
	case AbstainVote:
		return 0
	case ValidVoteWithContent:
		return 10
	case InvalidVoteWithContent:
		return 10
	case TrustedReference:
		return 6
	case DistrustedReference:
		return 0
	default:
		return 0
	}
}
