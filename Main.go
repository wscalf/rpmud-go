package main

import (
	"fmt"
	"rpmud/core"
	"rpmud/core/contract"
	"rpmud/telnet"
)

func main() {
	listener := telnet.TelnetListener{Port: 4000}
	clients, err := listener.Listen()
	commands := &core.HardcodedCommandSystem{}
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

func doLogin(c contract.ClientAdapter, commands core.CommandSystem, start *core.Room) {
	c.Write("Enter player name:")
	p := core.CreatePlayer(c, commands, <-c.MessagesChannel())
	start.Join(p)
}

func createWorld() *core.Room {
	start := core.CreateRoom("Start Room", "This is the starting room! You made it!")
	extra := core.CreateRoom("Overflow", "You've reached the overflow room!")

	start.LinkTo(extra, "north", "Overflow", "A doorway with a curtain over it and a sign that says 'Overflow'")
	extra.LinkTo(start, "south", "Start Room", "A doorway with a curtain over it")

	return start
}
