package world

import (
	"rpmud/server/gameplay/dependencies"
)

type Player struct {
	Object
	client   dependencies.Client
	inbound  chan string
	outbound chan string
	onInput  func(*Player, *Room, string)
	room     *Room
}

func (p *Player) Room() *Room {
	return p.room
}

func (p *Player) Enter(r *Room) {
	p.room = r
	p.room.AddPlayer(p)
}

func (p *Player) Leave() {
	p.room.RemovePlayer(p)
	p.room = nil
}

func (p *Player) Send(message string) {
	if p.outbound != nil {
		p.outbound <- message
	}
}

func (p *Player) handleInput() {
	for input := range p.inbound {
		p.onInput(p, p.room, input)
	}
	close(p.outbound)
	p.outbound = nil
	//When we get here, the peer is disconnected
	p.Leave()
}

func NewPlayer(client dependencies.Client, onInput func(*Player, *Room, string), name string) *Player {
	p := Player{}
	p.Name = name
	p.client = client
	p.onInput = onInput
	p.inbound, p.outbound = client.IOChannels()
	go p.handleInput()
	return &p
}
