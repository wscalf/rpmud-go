package telnet

import (
	"fmt"
	"net"
	"rpmud/gameplay/dependencies"
)

type TelnetListener struct {
	Port int
}

func (tl TelnetListener) Listen() (chan dependencies.Client, error) {
	ch := make(chan dependencies.Client)
	l, err := net.ListenTCP("tcp", &net.TCPAddr{Port: tl.Port})
	if err != nil {
		return nil, err
	}

	go handleConnections(l, ch)
	return ch, nil
}

func handleConnections(listener *net.TCPListener, ch chan dependencies.Client) {
	for {
		c, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			return
		}

		adapter := createTelnetAdapter(c)
		adapter.Init()
		ch <- adapter
	}
}
