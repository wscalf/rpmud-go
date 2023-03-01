package telnet

import (
	"fmt"
	"net"
	"strings"
)

type Option byte

var knownOptions = map[Option]bool{}

type Status int

const (
	Off Status = iota
	Enabling
	On
	Disabling
)

const (
	IAC  byte = 255
	GA   byte = 249
	NOP  byte = 241
	DO   byte = 253
	WILL byte = 251
	DONT byte = 254
	WONT byte = 252
)

const (
	CR byte = 13
	LF byte = 10
)

type TelnetAdapter struct {
	ch            chan string
	socket        *net.TCPConn
	sb            *strings.Builder
	serverOptions map[Option]Status
	clientOptions map[Option]Status
}

func createTelnetAdapter(socket *net.TCPConn) TelnetAdapter {
	adapter := TelnetAdapter{make(chan string), socket, &strings.Builder{}, make(map[Option]Status), make(map[Option]Status)}
	return adapter
}

func (t TelnetAdapter) Init() {
	go t.handleInput()
}

func (t TelnetAdapter) Write(output string) {
	t.socket.Write([]byte(output))
	t.sendCommand(CR, LF, IAC, GA)
}

func (t TelnetAdapter) sendCommand(bytes ...byte) {
	t.socket.Write(bytes)
}

func (t TelnetAdapter) MessagesChannel() chan string {
	return t.ch
}

func (t TelnetAdapter) handleBuffer(data []byte) {
	start := 0

	for i := 0; i < len(data); i++ {
		el := data[i]
		switch el {
		case IAC:
			t.sb.Write(data[start:i])
			i += t.handleCommand(data[i:])
			start = i
		case CR: //Not checking for the LF for now
			if t.sb.Len() == 0 {
				t.ch <- string(data[start:i])
			} else {
				t.sb.Write(data[start:i])
				t.ch <- t.sb.String()
				t.sb.Reset()
			}
			i += 2 //Skip CRLF
			start = i
		}
	}

	if start < len(data) {
		t.sb.Write(data[start:])
	}
}

func (t TelnetAdapter) requestServerOption(option Option) {
	t.serverOptions[option] = Enabling
	t.sendCommand(IAC, WILL, byte(option))
}

func (t TelnetAdapter) requestClientOption(option Option) {
	t.clientOptions[option] = Enabling
	t.sendCommand(IAC, DO, byte(option))
}

func (t TelnetAdapter) handleCommand(data []byte) int {
	//Process the command and return the number of bytes consumed
	switch data[1] {
	case DO, DONT, WILL, WONT:
		t.handleOptionCommand(data[1], Option(data[2])) //TODO: if a command is split across buffers, this will panic. May need to be able to leave some 'expectations' after processing a buffer so we can anticipate in the next buffer
		return 3
	case GA, NOP:
		return 2
	}
	return 0
}

func (t TelnetAdapter) handleOptionCommand(command byte, option Option) {
	fmt.Printf("Received option command: %d %d\n", option, command)
	switch command {
	case DO:
		if knownOptions[option] {
			switch t.serverOptions[option] {
			case Off:
				t.sendCommand(IAC, WILL, byte(option))
				fallthrough
			case Enabling:
				t.serverOptions[option] = On
			case Disabling:
				t.sendCommand(IAC, WONT, byte(option))
			}
		} else {
			t.sendCommand(IAC, WONT, byte(option))
		}
	case DONT:
		if t.serverOptions[option] == On || t.serverOptions[option] == Off {
			t.sendCommand(IAC, WONT, byte(option))
		}
		t.serverOptions[option] = Off
	case WILL:
		if knownOptions[option] { //Distinguish supported client and server options?
			switch t.clientOptions[option] {
			case Off:
				t.sendCommand(IAC, DO, byte(option))
				fallthrough
			case Enabling:
				t.clientOptions[option] = On
			case Disabling:
				t.sendCommand(IAC, DONT, byte(option))
			}
		} else {
			t.sendCommand(IAC, DONT, byte(option))
		}
	case WONT:
		if t.clientOptions[option] == On || t.clientOptions[option] == Off {
			t.sendCommand(IAC, DONT, byte(option))
		}

		t.clientOptions[option] = Off
	}
}

func (t TelnetAdapter) handleInput() {
	defer t.socket.Close()
	defer close(t.ch)

	buffer := make([]byte, 1024)

	for {
		len, err := t.socket.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		t.handleBuffer(buffer[:len])
	}
}
