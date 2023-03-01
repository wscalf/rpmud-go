package telnet

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"rpmud/gameplay/dependencies"
	"strings"
	"testing"
)

func TestTelnetAdapterRoundTrip(t *testing.T) {
	const value = "Test string\r\n"
	const valueWithoutEOL = "Test string"

	messages, adapter, client := setup()

	_, err := client.Write([]byte(value))
	if err != nil {
		t.Error(err)
	}

	received := <-messages

	if received != valueWithoutEOL {
		t.Fatalf("Expected %s but got %s", value, received)
	}

	adapter.Write(received)

	receiveBuffer := make([]byte, 16)
	i, err := client.Read(receiveBuffer)
	if err != nil {
		t.Error(err)
	}
	received = string(receiveBuffer[:i])

	if !strings.HasPrefix(received, value) {
		t.Fatalf("Expected %s but got %s", value, received)
	}
}

func BenchmarkAdapterRoundTrip(b *testing.B) {
	data := []byte("look Basket\r\n")
	messages, _, client := setup()

	for i := 0; i < b.N; i++ {
		client.Write(data)
		<-messages
	}
}

func TestTelnetAdapterDeclinesUnknownOption(t *testing.T) {
	request := []byte{IAC, WILL, 150}
	_, _, client := setup()

	buf := make([]byte, 16)
	client.Write(request)
	len, err := client.Read(buf)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal([]byte{IAC, DONT, 150}, buf[:len]) {
		t.Fatalf("Unexpected result. Got: %s", buf[:len])
	}

}

func setup() (messages chan string, adapter dependencies.Client, client net.Conn) {
	port := 4000 + rand.Intn(100)
	listener := TelnetListener{Port: port}
	adapters, _ := listener.Listen()

	client, _ = net.Dial("tcp", fmt.Sprintf("localhost:%d", port))

	adapter = <-adapters
	messages = adapter.MessagesChannel()

	return
}
