package commands

import (
	"rpmud/server/gameplay/commands/parameters"
	"rpmud/server/gameplay/world"
)

type Job interface {
	Run() error
}

type CmdJob struct {
	command Command
	params  map[string]parameters.Value
	player  *world.Player
	room    *world.Room
}

func (j CmdJob) Run() error {
	j.command.Exec(j.player, j.room, j.params)

	return nil
}

func NewCmdJob(cmd Command, params map[string]parameters.Value, player *world.Player, room *world.Room) CmdJob {
	return CmdJob{
		command: cmd,
		params:  params,
		player:  player,
		room:    room,
	}
}

type LinkJob struct {
	player *world.Player
	link   *world.Link
}

func (j LinkJob) Run() error {
	j.link.Activate(j.player)

	return nil
}

func NewLinkJob(player *world.Player, link *world.Link) LinkJob {
	return LinkJob{
		player: player,
		link:   link,
	}
}
