package dependencies

type Client interface {
	Init()
	MessagesChannel() chan string
	Write(output string)
}
