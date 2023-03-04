package dependencies

type ClientListener interface {
	Listen() chan Client
}
