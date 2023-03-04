package builtin

import (
	"fmt"
	"rpmud/gameplay/commands/parameters"
	"rpmud/gameplay/world"
)

type Look struct {
	params []parameters.Parameter
}

func (l *Look) Params() []parameters.Parameter {
	return l.params
}

func (l *Look) Exec(user *world.Player, r *world.Room, params map[string]parameters.Value) {
	if v, ok := params["at"]; ok {
		name := v.Single()
		if p, ok := r.FindPlayer(name); ok {
			user.Send(p.Describe())
			p.Send(fmt.Sprintf("%s looked at you.", user.Name))
		} else if l, ok := r.FindLink(name); ok {
			user.Send(l.Describe())
		}
	} else {
		user.Send(r.Describe())
	}
}

func NewLook() *Look {
	return &Look{
		params: []parameters.Parameter{
			parameters.NewName("at", false),
		},
	}
}
