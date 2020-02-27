package bus

import (
	"github.com/imdevin567/jargon/adapter"
)

// Entry ...
type Entry struct {
	Name     string
	From, To adapter.AbstractAdapter
	c        chan []byte
}

// Bus ...
type Bus struct {
	Entries []*Entry
}

// Instance ...
var Instance = &Bus{
	Entries: make([]*Entry, 0),
}

// NewEntry ...
func NewEntry(name string, from, to adapter.AbstractAdapter) *Entry {
	return &Entry{
		Name: name,
		From: from,
		To:   to,
		c:    make(chan []byte),
	}
}

// AddEntry ...
func (b *Bus) AddEntry(name string, from, to adapter.AbstractAdapter) {
	entry := NewEntry(name, from, to)
	Instance.Entries = append(Instance.Entries, entry)
}

// Start ...
func (b *Bus) Start() {
	for _, e := range b.Entries {
		go func(entry *Entry) {
			go adapter.StartAdapter(entry.From, entry.c)
			go adapter.StartAdapter(entry.To, entry.c)
		}(e)
	}

	for {

	}
}
