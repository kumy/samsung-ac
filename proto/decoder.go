package proto

import (
	"bytes"
	"encoding/hex"
	"errors"
	"strings"
)

type Decoder interface {
	Decode(frame []byte) Message
	DecodeFromString(frame string) Message
}

type decoder struct {
}

func NewDecoder() Decoder {
	d := &decoder{}
	return d
}

func (d *decoder) Decode(frame []byte) Message {
	m := EmptyMessage
	if err := d.check(frame); err != nil {
		panic(err)
	}

	//fmt.Printf("frame: %02x %02x\n", frame[9], frame)
	m.Counter = NewCounterFromBytes(frame[9])
	p := d.extractPayload(frame)
	m.Payload = p
	return &m
}

func (d *decoder) DecodeFromString(frame string) Message {
	frame = strings.Replace(frame, " ", "", -1)
	h, err := hex.DecodeString(frame)
	if err != nil {
		panic(err)
	}
	return d.Decode(h)
}

func (d *decoder) check(frame []byte) error {
	if frame == nil {
		return errors.New("frame is nil")
	}
	if bytes.Equal(frame[0:6], HEADER[:]) {
		return errors.New("bad frame header: " + string(frame[0:6]))
	}
	if bytes.Equal(frame[len(frame)-1:], HEADER[:]) {
		return errors.New("bad frame footer: " + string(frame[len(frame)-1:]))
	}
	// TODO check checksum
	return nil
}

func (d *decoder) extractPayload(frame []byte) Payload {
	p := NewPayload(MessageType(frame[11:]))
	if frame[13] == 0 {
		return p
	}
	rs := d.extractRegisters(frame)
	p.AddRegisters(rs)
	return p
}

func (d *decoder) extractRegisters(frame []byte) []Register {
	var rs []Register
	//fmt.Printf("%02x\n", frame)
	l := frame[13]
	c := frame[14 : 14+int(l)]

	if len(c) != int(l) {
		panic("invalid register sizes")
	}

	for i := byte(0); i < l; {
		rid := c[i]
		rl := c[i+1]
		rv := c[i+2 : i+2+rl]

		r := &register{
			RegisterId:   rid,
			RegisterSize: rl,
			Value:        rv,
		}
		rs = append(rs, r)
		i += 2 + rl
	}

	return rs
}
