package gameplay

type Object struct {
	ID          string
	Name        string
	Description string
}

func (o *Object) Object() *Object {
	return o
}

func (o *Object) Describe() string {
	return o.Description
}
