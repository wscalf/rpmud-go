package dependencies

type Client interface {
	Init()
	IOChannels() (chan string, chan string)
}
