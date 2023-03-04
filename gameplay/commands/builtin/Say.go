package builtin

import (
	"fmt"
	"rpmud/gameplay/commands/parameters"
	"rpmud/gameplay/world"
)

type Say struct {
	params []parameters.Parameter
}

func (s *Say) Params() []parameters.Parameter {
	return s.params
}

func (l *Say) Exec(user *world.Player, r *world.Room, params map[string]parameters.Value) {
	msg := params["message"]

	r.SendToAllExcept(user, fmt.Sprintf(`%s says, "%s"`, user.Name, msg.Single()))
	user.Send(fmt.Sprintf(`You say, "%s"`, msg.Single()))
}

func NewSay() *Say {
	return &Say{
		params: []parameters.Parameter{parameters.NewFreeText("message", true)},
	}
}
