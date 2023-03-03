package gameplay

import (
	"errors"
	"fmt"
	"strings"
)

type CommandSystem interface {
	Execute(player *Player, command string) error
}

type HardcodedCommandSystem struct {
}

func (s *HardcodedCommandSystem) Execute(player *Player, command string) error {
	room := player.Room()
	if link, hit := room.FindLink(command); hit {
		link.Activate(player)
		return nil
	}

	words := strings.Split(command, " ")
	if words != nil {
		switch words[0] {
		case "look":
			switch len(words) {
			case 1:
				player.Send(room.Describe())
				return nil
			case 2:
				var o *Object
				hit := false
				name := words[1]

				if p, hit := room.FindPlayer(name); hit {
					o = &p.Object
				} else if l, hit := room.FindLink(name); hit {
					o = &l.Object
				}

				if hit {
					player.Send(o.Describe())
					return nil
				} else {
					return errors.New("Cannot find: " + words[1])
				}
			default:
				return errors.New("look command should have zero or one arguments")
			}
		case "say":
			message := command[4:]
			room.SendToAll(fmt.Sprintf("%s says, \"%s\"", player.Name, message))
			return nil
		}
	}

	return errors.New("command not recognized")
}
