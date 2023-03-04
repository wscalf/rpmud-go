package commands

import (
	"fmt"
	"rpmud/gameplay/commands/builtin"
	"rpmud/gameplay/commands/parameters"
	"rpmud/gameplay/world"
	"strings"
)

type Interpreter struct {
	commands map[string]Command
}

func (i *Interpreter) Prepare(line string, player *world.Player, room *world.Room) (Job, error) {
	name, remaining := consumeCommand(line)

	if link, hit := room.FindLink(name.Single()); hit {
		return NewLinkJob(player, link), nil
	}

	if cmd, hit := i.commands[name.Single()]; hit {
		values := make(map[string]parameters.Value)
		for _, param := range cmd.Params() {
			if param.Find(remaining) == 0 {
				values[param.Name()], remaining = param.Consume(remaining)
			} else {
				if param.IsRequired() {
					return nil, fmt.Errorf("Missing required parameter: %s. %w", param.Name(), ErrInvalidInput)
				}
				continue
			}
		}

		return NewCmdJob(cmd, values, player, room), nil
	}

	return nil, fmt.Errorf("Unrecognized command: %s. %w", name, ErrInvalidInput)
}

func (i *Interpreter) Register(alias string, command Command) {
	i.commands[alias] = command
}

func NewInterpreter() *Interpreter {
	i := Interpreter{
		commands: make(map[string]Command),
	}

	i.Register("look", builtin.NewLook())

	return &i
}

func consumeCommand(command string) (parameters.Value, string) {
	if idx := strings.Index(command, " "); idx != -1 {
		v := parameters.NewSingleValue(command[:idx])
		return v, strings.TrimSpace(command[idx:])
	} else {
		v := parameters.NewSingleValue(command)
		return v, ""
	}
}
