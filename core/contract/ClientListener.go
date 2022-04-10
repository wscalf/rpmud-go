package contract

type ClientListener interface {
	Listen() chan ClientAdapter
}
