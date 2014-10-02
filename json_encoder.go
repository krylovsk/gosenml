package gosenml

import (
	"encoding/json"
)

type JsonEncoder struct{}

func NewJsonEncoder() *JsonEncoder {
	return &JsonEncoder{}
}

func (self *JsonEncoder) DecodeMessage(data []byte) (Message, error) {
	m := Message{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return Message{}, err
	}

	if err = m.Validate(); err != nil {
		return m, err
	}
	return m, nil
}

func (self *JsonEncoder) EncodeMessage(m *Message) ([]byte, error) {
	if err := m.Validate(); err != nil {
		return []byte{}, err
	}

	b, err := json.Marshal(m)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

func (self *JsonEncoder) EncodeEntry(e *Entry) ([]byte, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

func (self *JsonEncoder) DecodeEntry(data []byte) (Entry, error) {
	e := Entry{}
	err := json.Unmarshal(data, &e)
	if err != nil {
		return Entry{}, err
	}
	return e, nil
}
