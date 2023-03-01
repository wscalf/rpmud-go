package gameplay

type Transition interface {
	Transition() *TransitionData
	Object() *ObjectData
	Activate(player *Player)
	Describe() string
}

type TransitionData struct {
	ObjectData
	Command string
}

func (t *TransitionData) Transition() *TransitionData {
	return t
}
