package config

import (
	"github.com/imdevin567/jargon/protocol"

	"github.com/imdevin567/jargon/adapter"
	"github.com/imdevin567/jargon/bus"
)

// Config ...
var Config Jargonfile

// Jargonfile ...
type Jargonfile struct {
	Routes []Route
}

// Route ...
type Route struct {
	Name   string
	Input  genericAdapter
	Output genericAdapter
}

// genericAdapter ...
type genericAdapter struct {
	Adapter     string
	Host        string
	Port        int
	Delimiter   string
	Path        string
	ContentType string
}

// CreateEntries ...
func (jf *Jargonfile) CreateEntries() {
	for _, route := range jf.Routes {
		input := jf.createAdapter(adapter.Input, route.Input)
		output := jf.createAdapter(adapter.Output, route.Output)
		bus.Instance.AddEntry(route.Name, input, output)
	}
}

func (jf *Jargonfile) createAdapter(direction adapter.Direction, route genericAdapter) adapter.AbstractAdapter {
	switch p := protocol.FromString(route.Adapter); p {
	case protocol.HTTP:
		return adapter.NewHTTPAdapter(direction, route.Host, route.Port, route.Path, route.ContentType)
	case protocol.TCP:
		return adapter.NewTCPAdapter(direction, route.Host, route.Port, []byte(route.Delimiter)[0])
	case protocol.UDP:
		return adapter.NewUDPAdapter(direction, route.Host, route.Port, []byte(route.Delimiter)[0])
	case protocol.WS:
		return adapter.NewWSAdapter(direction, route.Host, route.Port, route.Path)
	}

	return nil
}
