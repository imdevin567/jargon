package adapter

import (
	"bufio"
	"fmt"
	"net"

	. "github.com/imdevin567/jargon/protocol"
)

// UDPAdapter ... (T2)
type UDPAdapter struct {
	*Adapter
	Delimiter byte
	Conn      net.Conn
}

// NewUDPAdapter ...
func NewUDPAdapter(direction Direction, host string, port int, delimiter byte) *UDPAdapter {
	adapter := NewAdapter(direction, host, port, UDP)
	return &UDPAdapter{
		Adapter:   adapter,
		Delimiter: delimiter,
	}
}

// Start ...
func (udp *UDPAdapter) Start() {
	if udp.Direction == Input {
		udp.listen()
	} else if udp.Direction == Output {
		udp.connect()
	} else {
		fmt.Println("Seriously, what are you doing?")
	}
}

// Input = create server ...
func (udp *UDPAdapter) Input(c chan []byte) {
	reader := bufio.NewReader(udp.Conn)
	for {
		msg, err := reader.ReadBytes(udp.Delimiter)
		if err != nil {
			// TODO: Handle error
		}

		c <- msg
	}
}

// Output = connect to server ...
func (udp *UDPAdapter) Output(c chan []byte) {
	for {
		msg := <-c
		_, err := udp.Conn.Write(msg)
		if err != nil {
			// TODO: Handle error
		}
	}
}

// connect ...
func (udp *UDPAdapter) connect() {
	conn, err := net.Dial("udp", fmt.Sprintf("%v:%v", udp.Host, udp.Port))
	if err != nil {
		// TODO: Handle error
		return
	}

	udp.Conn = conn
}

// listen ...
func (udp *UDPAdapter) listen() {
	server, err := net.Listen("udp", fmt.Sprintf("%v:%v", udp.Host, udp.Port))
	if err != nil {
		// TODO: Handle error
		return
	}

	conn, err := server.Accept()
	if err != nil {
		// TODO: Handle error
		return
	}

	udp.Conn = conn
}
