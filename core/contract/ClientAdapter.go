package contract

type ClientAdapter interface {
	Init()
	MessagesChannel() chan string
	Write(output string)
}
