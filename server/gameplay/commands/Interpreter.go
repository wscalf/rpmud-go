package commands

import (
	"fmt"
	"rpmud/server/gameplay/commands/builtin"
	"rpmud/server/gameplay/commands/parameters"
	"rpmud/server/gameplay/world"
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
		params := cmd.Params()

		for i, param := range params {
			next := i + 1
			if next < len(params) && params[next].IsDelimiter() {
				delimiter := params[next]
				end := delimiter.Find(remaining)
				if end < 0 {
					return nil, fmt.Errorf("Expected delimiter: %s. %w", delimiter.Name(), ErrInvalidInput)
				}

				delimitedPortion := remaining[:end]

				var err error
				delimitedPortion, err = parsePortion(delimitedPortion, param, values)

				if err != nil {
					return nil, err
				}

				if len(delimitedPortion) != 0 {
					return nil, fmt.Errorf("Malformed value for parameter: %s. %w", param.Name(), ErrInvalidInput)
				}

				remaining = remaining[end:]

			} else {
				var err error
				remaining, err = parsePortion(remaining, param, values)

				if err != nil {
					return nil, err
				}
			}

		}

		return NewCmdJob(cmd, values, player, room), nil
	}

	return nil, fmt.Errorf("Unrecognized command: %s. %w", name, ErrInvalidInput)
}

func parsePortion(portion string, param parameters.Parameter, values map[string]parameters.Value) (string, error) {
	if param.Find(portion) == 0 {
		v, remaining := param.Consume(portion)

		values[param.Name()] = v

		return remaining, nil
	} else {
		if param.IsRequired() {
			return portion, fmt.Errorf("Missing required parameter: %s. %w", param.Name(), ErrInvalidInput)
		}

		return portion, nil
	}
}

func (i *Interpreter) Register(alias string, command Command) {
	i.commands[alias] = command
}

func NewInterpreter() *Interpreter {
	i := Interpreter{
		commands: make(map[string]Command),
	}

	i.Register("look", builtin.NewLook())
	i.Register("say", builtin.NewSay())

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
