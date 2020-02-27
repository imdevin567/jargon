// Copyright (c) Jargon Author(s) 2020. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package adapter

import (
	"fmt"

	. "github.com/imdevin567/jargon/protocol"
)

// Direction ...
type Direction int

//
const (
	Input Direction = iota + 1
	Output
)

// AbstractAdapter ...(I)
type AbstractAdapter interface {
	Start()
	Input(c chan []byte)
	Output(c chan []byte)
	GetDirection() Direction
}

// Adapter ...(T1)
type Adapter struct {
	Direction
	Host     string
	Port     int
	Protocol Protocol
}

// NewAdapter ...
func NewAdapter(direction Direction, host string, port int, protocol Protocol) *Adapter {
	return &Adapter{
		Direction: direction,
		Host:      host,
		Port:      port,
		Protocol:  protocol,
	}
}

// Abstract interface
func (a *Adapter) Start()               {}
func (a *Adapter) Input(c chan []byte)  {}
func (a *Adapter) Output(c chan []byte) {}
func (a *Adapter) GetDirection() Direction {
	return a.Direction
}

// StartAdapter ...
func StartAdapter(a AbstractAdapter, c chan []byte) {
	a.Start()
	direction := a.GetDirection()
	if direction == Input {
		a.Input(c)
	} else if direction == Output {
		a.Output(c)
	} else {
		fmt.Println("What exactly are you trying to do here?")
	}
}
