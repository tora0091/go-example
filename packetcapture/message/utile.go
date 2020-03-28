package message

import (
	"bytes"
	"encoding/gob"
)

type Target struct {
	Name string
	Age  int8
	Area string
	Job  string
}

func NewTarget() *Target {
	return &Target{}
}

func (t *Target) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buf).Encode(t)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t *Target) Unmarshal(data []byte) error {
	buf := bytes.NewBuffer(data)
	err := gob.NewDecoder(buf).Decode(t)
	if err != nil {
		return err
	}
	return nil
}
