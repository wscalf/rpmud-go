package gameplay

import (
	"rpmud/server/gameplay/commands"
	"rpmud/server/gameplay/dependencies"
	"rpmud/server/gameplay/world"
)

type Game struct {
	players     []*world.Player
	interpreter *commands.Interpreter
	jobQueue    chan commands.Job
	start       *world.Room
}

func (g *Game) Start() {
	go g.runJobs()
}

func (g *Game) Join(client dependencies.Client) {
	inbound, outbound := client.IOChannels()

	outbound <- "Enter player name:"
	p := world.NewPlayer(client, g.processCommand, <-inbound)

	g.jobQueue <- JoinJob{
		player: p,
		room:   g.start,
		game:   g,
	}

}

func (g *Game) processCommand(p *world.Player, r *world.Room, command string) {
	if job, err := g.interpreter.Prepare(command, p, r); err == nil {
		g.jobQueue <- job
	} else {
		p.Send(err.Error())
	}
}

func (g *Game) runJobs() {
	for job := range g.jobQueue {
		job.Run()
	}
}

func NewGame(start *world.Room, interpreter *commands.Interpreter) *Game {
	return &Game{
		players:     make([]*world.Player, 16),
		interpreter: interpreter,
		jobQueue:    make(chan commands.Job),
		start:       start,
	}
}

type JoinJob struct {
	player *world.Player
	room   *world.Room
	game   *Game
}

func (j JoinJob) Run() error {
	j.game.players = append(j.game.players, j.player)

	j.player.Enter(j.room)
	return nil
}
