package gosenml

import "fmt"

// Data model as described in
// http://tools.ietf.org/html/draft-jennings-senml-10

// Root variable
type Message struct {
	BaseName  string  `json:"bn,omitempty"`
	BaseTime  int64   `json:"bt,omitempty"`
	BaseUnits string  `json:"bu,omitempty"`
	Version   int     `json:"ver"`
	Entries   []Entry `json:"e"`
}

// Measurement of Parameter Entry
type Entry struct {
	Name         string  `json:"n,omitempty"`
	Units        string  `json:"u,omitempty"`
	Value        float64 `json:"v"`
	StringValue  string  `json:"sv"`
	BooleanValue bool    `json:"bv"`
	Sum          float64 `json:"s,omitempty"`
	Time         int64   `json:"t,omitempty"`
	UpdateTime   int64   `json:"ut,omitempty"`
}

func NewMessage(entries ...Entry) *Message {
	return &Message{
		Version: 1.0,
		Entries: entries,
	}
}

// Makes a deep copy of the message
func (self *Message) copy() Message {
	mc := *self
	entries := make([]Entry, len(self.Entries))
	copy(entries, self.Entries)
	mc.Entries = entries
	return mc
}

// Validates a message
func (self *Message) Validate() error {
	if len(self.Entries) == 0 {
		return fmt.Errorf("Invalid Message: entries must be non-empty")
	}
	// TODO: more validation
	return nil
}

// Returns a copy with all Enties expanded ("self-contained")
func (self *Message) Expand() Message {
	m := self.copy()

	for i, e := range m.Entries {
		// BaseName
		e.Name = m.BaseName + e.Name

		// BaseTime
		e.Time = m.BaseTime + e.Time

		// BaseUnits
		if e.Units == "" {
			e.Units = m.BaseUnits
		}
		m.Entries[i] = e
	}
	m.BaseName = ""
	m.BaseTime = 0
	m.BaseUnits = ""
	return m
}

// Returns a copy with all Entries compacted (common data put into Message)
func (self *Message) Compact() Message {
	m := self.copy()
	// TODO
	// BaseName
	// BaseTime
	// BaseUnits
	return m
}

type SenmlEncoder interface {
	EncodeMessage(*Message) ([]byte, error)
	DecodeMessage([]byte) (Message, error)
	EncodeEntry(*Entry) ([]byte, error)
	DecodeEntry([]byte) (Entry, error)
}
