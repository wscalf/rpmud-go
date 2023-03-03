package gameplay

import (
	"rpmud/gameplay/dependencies"
)

type Player struct {
	Object
	client   dependencies.Client
	room     *Room
	commands CommandSystem
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
	p.client.Write(message)
}

func (p *Player) handleInput() {
	ch := p.client.MessagesChannel()

	for input := range ch {
		err := p.commands.Execute(p, input)
		if err != nil {
			p.client.Write(err.Error()) //Should probably distinguish between input errors and server errors, server errors should definitely be logged and probably not sent to the client
		}
	}
	//When we get here, the peer is disconnected
	p.Leave()
}

func NewPlayer(client dependencies.Client, commands CommandSystem, name string) *Player {
	p := Player{}
	p.Name = name
	p.client = client
	p.commands = commands
	go p.handleInput()
	return &p
}
