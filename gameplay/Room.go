package gameplay

import (
	"container/list"
	"fmt"
	"strings"
)

type Room struct {
	ObjectData
	objects     *list.List
	players     *list.List
	transitions *list.List
}

func (r *Room) Join(player *Player) {
	r.players.PushBack(player)
	player.Room = r
	r.SendToAllExcept(player, fmt.Sprintf("%s has entered the room.", player.Name))

	player.Write(r.Describe())
}

func (r *Room) Leave(player *Player) {
	r.SendToAllExcept(player, fmt.Sprintf("%s has left the room.", player.Name))
	removeItem(r.players, player)
	player.Room = nil
}

func (r *Room) SendToAll(message string) {
	for e := r.players.Front(); e != nil; e = e.Next() {
		e.Value.(*Player).Write(message)
	}
}

func (r *Room) SendToAllExcept(skip *Player, message string) {
	for e := r.players.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Player)
		if p.Name == skip.Name { //Change to id?
			continue
		}

		p.Write(message)
	}
}

func (r *Room) Find(name string) (Object, bool) {
	for e := r.players.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Player)
		if strings.HasPrefix(p.Name, name) {
			return p, true
		}
	}

	for e := r.transitions.Front(); e != nil; e = e.Next() {
		t := e.Value.(Transition)
		if strings.HasPrefix(t.Transition().Command, name) {
			return t, true
		}
	}

	return nil, false
}

func (r *Room) TryActivateTransition(p *Player, command string) bool {
	for e := r.transitions.Front(); e != nil; e = e.Next() {
		t := e.Value.(Transition)
		if strings.HasPrefix(t.Transition().Command, command) {
			t.Activate(p)
			return true
		}
	}

	return false
}

func CreateRoom(name string, description string) *Room {
	r := Room{
		objects:     list.New(),
		players:     list.New(),
		transitions: list.New(),
	}

	r.Name = name
	r.Description = description

	return &r
}

func (r *Room) Describe() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("[%s]\n\n", r.Name))
	sb.WriteString(r.Description)
	sb.WriteRune('\n')
	sb.WriteString("------------------------------------\n")
	for e := r.transitions.Front(); e != nil; e = e.Next() {
		t := e.Value.(Transition).Transition()
		sb.WriteString(fmt.Sprintf("[%s] %s\n", t.Command, t.Name))
	}
	sb.WriteString("------------------------------------\n")
	for e := r.players.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Player)
		sb.WriteString(p.Name)
		sb.WriteRune('\n')
	}
	sb.WriteString("------------------------------------\n")
	return sb.String()
}

func (from *Room) LinkTo(to *Room, command string, name string, description string) {
	t := DirectTransition{from: from, to: to}
	t.Command = command
	t.Name = name
	t.Description = description

	from.transitions.PushBack(&t)
}

func removeItem(items *list.List, toRemove Object) {
	r := toRemove.Object()
	for e := items.Front(); e != nil; e = e.Next() {
		o := e.Value.(Object).Object()
		if r.Name == o.Name {
			items.Remove(e)
			return
		}
	}
}
