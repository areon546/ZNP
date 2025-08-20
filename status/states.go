package status

type State int

const (
	INIT State = iota
	HEAR
	TALK
	STOP
)

func Lookup(s State) string {
	switch s {
	case INIT:
		return "Initialising Conversation"
	case HEAR:
		return "Hearing Message"
	case TALK:
		return "Talking Message"
	case STOP:
		return "Ending conversation"
	default:
		return "UNKNOWN"
	}
}
