package commands

import (
	"rpmud/server/gameplay/commands/parameters"
	"rpmud/server/gameplay/world"
)

type Command interface {
	Params() []parameters.Parameter
	Exec(p *world.Player, r *world.Room, params map[string]parameters.Value)
}
