package main

import (
	"fmt"
	"rpmud/gameplay"
	"rpmud/gameplay/dependencies"
	"rpmud/infrastructure/telnet"
)

func main() {
	listener := telnet.TelnetListener{}
	clients, err := listener.ListenTCP(4000)
	commands := &gameplay.HardcodedCommandSystem{}
	if err != nil {
		fmt.Println(err)
		return
	}

	start := createWorld()

	fmt.Println("Server up and listening")

	for {
		c := <-clients
		go doLogin(c, commands, start)
	}
}

func doLogin(c dependencies.Client, commands gameplay.CommandSystem, start *gameplay.Room) {
	c.Write("Enter player name:")
	p := gameplay.NewPlayer(c, commands, <-c.MessagesChannel())
	p.Enter(start)
}

func createWorld() *gameplay.Room {
	start := gameplay.CreateRoom("Start Room", "This is the starting room! You made it!")
	extra := gameplay.CreateRoom("Overflow", "You've reached the overflow room!")

	start.LinkTo(extra, "north", "Overflow", "A doorway with a curtain over it and a sign that says 'Overflow'")
	extra.LinkTo(start, "south", "Start Room", "A doorway with a curtain over it")

	return start
}
