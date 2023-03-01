package gameplay

type DirectTransition struct {
	TransitionData
	from *Room
	to   *Room
}

func (t *DirectTransition) Data() TransitionData {
	return t.TransitionData
}

func (t *DirectTransition) Activate(player *Player) {
	t.from.Leave(player)
	t.to.Join(player)
}
