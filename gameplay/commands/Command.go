package commands

import (
	"rpmud/gameplay/commands/parameters"
	"rpmud/gameplay/world"
)

type Command interface {
	Params() []parameters.Parameter
	Exec(p *world.Player, r *world.Room, params map[string]parameters.Value)
}
