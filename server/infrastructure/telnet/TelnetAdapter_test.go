package telnet

import (
	"bytes"
	"net"
	"rpmud/server/gameplay/dependencies"
	"strings"
	"testing"

	"google.golang.org/grpc/test/bufconn"
)

func TestTelnetAdapterRoundTrip(t *testing.T) {
	t.Parallel()

	const value = "Test string\r\n"
	const valueWithoutEOL = "Test string"

	inbound, outbound, _, client := setup()

	_, err := client.Write([]byte(value))
	if err != nil {
		t.Error(err)
	}

	received := <-inbound

	if received != valueWithoutEOL {
		t.Fatalf("Expected %s but got %s", value, received)
	}

	outbound <- received

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
	inbound, _, _, client := setup()

	for i := 0; i < b.N; i++ {
		if _, err := client.Write(data); err != nil {
			b.Error(err)
		}

		<-inbound
	}
}

func TestTelnetAdapterDeclinesUnknownOption(t *testing.T) {
	t.Parallel()

	request := []byte{IAC, WILL, 150}
	_, _, _, client := setup()

	buf := make([]byte, 16)
	if _, err := client.Write(request); err != nil {
		t.Error(err)
	}

	len, err := client.Read(buf)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal([]byte{IAC, DONT, 150}, buf[:len]) {
		t.Fatalf("Unexpected result. Got: %s", buf[:len])
	}

}

func setup() (inbound chan string, outbound chan string, adapter dependencies.Client, client net.Conn) {
	buf := bufconn.Listen(2048)
	listener := TelnetListener{}
	adapters, _ := listener.Listen(buf)

	client, _ = buf.Dial()

	adapter = <-adapters
	inbound, outbound = adapter.IOChannels()

	return
}
