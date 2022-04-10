package core

import (
	"rpmud/core/contract"
)

type Player struct {
	ObjectData
	contract.ClientAdapter
	Room     *Room
	commands CommandSystem
}

func (p *Player) handleInput() {
	ch := p.MessagesChannel()

	for input := range ch {
		err := p.commands.Execute(p, input)
		if err != nil {
			p.Write(err.Error()) //Should probably distinguish between input errors and server errors, server errors should definitely be logged and probably not sent to the client
		}
	}
	//When we get here, the peer is disconnected
	p.Room.Leave(p)
}

func CreatePlayer(adapter contract.ClientAdapter, commands CommandSystem, name string) *Player {
	p := Player{}
	p.Name = name
	p.ClientAdapter = adapter
	p.commands = commands
	go p.handleInput()
	return &p
}
