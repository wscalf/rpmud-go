package telnet

import (
	"fmt"
	"net"
	"rpmud/core/contract"
)

type TelnetListener struct {
	Port int
}

func (tl TelnetListener) Listen() (chan contract.ClientAdapter, error) {
	ch := make(chan contract.ClientAdapter)
	l, err := net.ListenTCP("tcp", &net.TCPAddr{Port: tl.Port})
	if err != nil {
		return nil, err
	}

	go handleConnections(l, ch)
	return ch, nil
}

func handleConnections(listener *net.TCPListener, ch chan contract.ClientAdapter) {
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
