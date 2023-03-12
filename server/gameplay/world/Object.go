package world

import "rpmud/server/gameplay/dependencies"

type Object struct {
	ID          string
	Name        string
	Description string
	Script      *dependencies.ScriptObject
}

func (o *Object) Object() *Object {
	return o
}

func (o *Object) Describe() string {
	return o.Description
}
