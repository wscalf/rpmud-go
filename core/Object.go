package core

type ObjectData struct {
	ID          string
	Name        string
	Description string
}

type Object interface {
	Object() *ObjectData
	Describe() string
}

func (o *ObjectData) Object() *ObjectData {
	return o
}

func (o *ObjectData) Describe() string {
	return o.Description
}
