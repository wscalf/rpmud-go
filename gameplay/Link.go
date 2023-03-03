package gameplay

type Link struct { //May evolve back into an interface later with this becoming a HardLink vs a ScriptLink
	Object
	Command string
	to      *Room
}

func (l *Link) Activate(player *Player) {
	player.Leave()
	player.Enter(l.to)
}
