package infrastructure

type StubClient struct {
}

func (s StubClient) Init() {

}

func (s StubClient) IOChannels() (chan string, chan string) {
	return make(chan string, 10), make(chan string, 10)
}
