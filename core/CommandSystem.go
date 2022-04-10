package core

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
	room := player.Room
	if room.TryActivateTransition(player, command) {
		return nil
	}

	words := strings.Split(command, " ")
	if words != nil {
		switch words[0] {
		case "look":
			switch len(words) {
			case 1:
				player.Write(room.Describe())
				return nil
			case 2:
				o, hit := room.Find(words[1])
				if hit {
					player.Write(o.Describe())
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
