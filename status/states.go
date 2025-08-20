package status

type State string

const (
	LISTENING = "listening"
	TALKING   = "talking"
	STOPPING  = "stopping"
)
