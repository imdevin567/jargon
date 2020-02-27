package adapter

import (
	"bufio"
	"fmt"
	"net"

	. "github.com/imdevin567/jargon/protocol"
)

// TCPAdapter ... (T2)
type TCPAdapter struct {
	*Adapter
	Delimiter byte
	Conn      net.Conn
}

// NewTCPAdapter ...
func NewTCPAdapter(direction Direction, host string, port int, delimiter byte) *TCPAdapter {
	adapter := NewAdapter(direction, host, port, TCP)
	return &TCPAdapter{
		Adapter:   adapter,
		Delimiter: delimiter,
	}
}

// Start ...
func (tcp *TCPAdapter) Start() {
	if tcp.Direction == Input {
		tcp.listen()
	} else if tcp.Direction == Output {
		tcp.connect()
	} else {
		fmt.Println("Seriously, what are you doing?")
	}
}

// Input = create server ...
func (tcp *TCPAdapter) Input(c chan []byte) {
	defer tcp.Conn.Close()
	reader := bufio.NewReader(tcp.Conn)
	for {
		msg, err := reader.ReadBytes(tcp.Delimiter)
		if err != nil {
			// TODO: Handle error
		}

		c <- msg
	}
}

// Output = connect to server ...
func (tcp *TCPAdapter) Output(c chan []byte) {
	defer tcp.Conn.Close()
	for {
		msg := <-c
		_, err := tcp.Conn.Write(msg)
		if err != nil {
			// TODO: Handle error
		}
	}
}

// connect ...
func (tcp *TCPAdapter) connect() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", tcp.Host, tcp.Port))
	if err != nil {
		// TODO: Handle error
		return
	}

	tcp.Conn = conn
}

// listen ...
func (tcp *TCPAdapter) listen() {
	server, err := net.Listen("tcp", fmt.Sprintf("%v:%v", tcp.Host, tcp.Port))
	if err != nil {
		// TODO: Handle error
		return
	}

	conn, err := server.Accept()
	if err != nil {
		// TODO: Handle error
		return
	}

	tcp.Conn = conn
}
