package protocol

import (
	"strings"
)

// Protocol ...
type Protocol int

//
const (
	AMQP Protocol = iota + 1
	HTTP
	RPC
	TCP
	UDP
	WS
)

// ToString ...
func (protocol Protocol) ToString() string {
	protocols := [6]string{
		"AMQP",
		"HTTP",
		"RPC",
		"TCP",
		"UDP",
		"WS",
	}

	if int(protocol) > len(protocols) || int(protocol) <= 0 {
		return ""
	}

	return protocols[protocol-1]
}

// FromString ...
func FromString(val string) Protocol {
	var p Protocol

	switch lower := strings.ToLower(val); lower {
	case "amqp":
		p = AMQP
	case "http":
		p = HTTP
	case "rpc":
		p = RPC
	case "tcp":
		p = TCP
	case "udp":
		p = UDP
	case "ws":
	case "websocket":
		p = WS
	default:
		panic("Invalid protocol set for adapter!")
	}

	return p
}
