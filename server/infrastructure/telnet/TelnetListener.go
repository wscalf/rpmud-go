package telnet

import (
	"fmt"
	"net"
	"rpmud/server/gameplay/dependencies"
)

type TelnetListener struct {
}

func (tl *TelnetListener) ListenTCP(port int) (chan dependencies.Client, error) {
	l, err := net.ListenTCP("tcp", &net.TCPAddr{Port: port})
	if err != nil {
		return nil, err
	}

	return tl.Listen(l)
}

func (tl *TelnetListener) Listen(ls net.Listener) (chan dependencies.Client, error) {
	ch := make(chan dependencies.Client)

	go handleConnections(ls, ch)
	return ch, nil
}

func handleConnections(listener net.Listener, ch chan dependencies.Client) {
	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		adapter := createTelnetAdapter(c)
		adapter.Init()
		ch <- adapter
	}
}
