package gameplay

import (
	"container/list"
	"fmt"
	"strings"
)

type Room struct {
	Object
	objects *list.List
	players *list.List
	links   *list.List
}

func (r *Room) AddPlayer(player *Player) {
	r.players.PushBack(player)
	r.SendToAllExcept(player, fmt.Sprintf("%s has entered the room.", player.Name))

	player.Send(r.Describe())
}

func (r *Room) RemovePlayer(player *Player) {
	removeItem(r.players, func(a any) bool { return a.(*Player).Name == player.Name })
	r.SendToAllExcept(player, fmt.Sprintf("%s has left the room.", player.Name))
}

func (r *Room) SendToAll(message string) {
	for e := r.players.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Player)
		p.Send(message)
	}
}

func (r *Room) SendToAllExcept(skip *Player, message string) {
	for e := r.players.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Player)
		if p.Name == skip.Name { //Change to id?
			continue
		}

		p.Send(message)
	}
}

func (r *Room) FindPlayer(name string) (*Player, bool) {
	for e := r.players.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Player)
		if strings.HasPrefix(p.Name, name) {
			return p, true
		}
	}

	return nil, false
}

func (r *Room) FindLink(name string) (*Link, bool) {
	for e := r.links.Front(); e != nil; e = e.Next() {
		l := e.Value.(*Link)
		if strings.HasPrefix(l.Command, name) {
			return l, true
		}
	}

	return nil, false
}

func NewRoom(name string, description string) *Room {
	r := Room{
		objects: list.New(),
		players: list.New(),
		links:   list.New(),
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
	for e := r.links.Front(); e != nil; e = e.Next() {
		t := e.Value.(*Link)
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
	l := Link{}
	l.to = to
	l.Command = command
	l.Name = name
	l.Description = description

	from.links.PushBack(&l)
}

func seekItem(items *list.List, predicate func(any) bool) (*list.Element, bool) {
	for item := items.Front(); item != nil; item = item.Next() {
		if predicate(item.Value) {
			return item, true
		}
	}

	return nil, false
}

func removeItem(items *list.List, predicate func(any) bool) {
	if e, ok := seekItem(items, predicate); ok {
		items.Remove(e)
	}
}
