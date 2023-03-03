package gameplay

type Link struct {
	Object
	Command string
	to      *Room
}

func (l *Link) Activate(player *Player) {
	player.Leave()
	player.Enter(l.to)
}
