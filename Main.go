package main

import (
	"fmt"
	"rpmud/gameplay"
	"rpmud/gameplay/commands"
	"rpmud/gameplay/world"
	"rpmud/infrastructure/telnet"
)

func main() {
	listener := telnet.TelnetListener{}
	clients, err := listener.ListenTCP(4000)

	if err != nil {
		fmt.Println(err)
		return
	}

	interpreter := commands.NewInterpreter()
	game := gameplay.NewGame(createWorld(), interpreter)
	game.Start()

	fmt.Println("Server up and listening")

	for {
		c := <-clients
		go game.Join(c)
	}
}

func createWorld() *world.Room {
	start := world.NewRoom("Start Room", "This is the starting room! You made it!")
	extra := world.NewRoom("Overflow", "You've reached the overflow room!")

	start.LinkTo(extra, "north", "Overflow", "A doorway with a curtain over it and a sign that says 'Overflow'")
	extra.LinkTo(start, "south", "Start Room", "A doorway with a curtain over it")

	return start
}
