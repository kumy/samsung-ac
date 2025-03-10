package proto

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/fatih/color"
	"strings"
)

func appendBytes(slice []byte, buf ...byte) []byte {
	out := make([]byte, len(slice)+len(buf))
	copy(out, slice)
	copy(out[len(slice):], buf)
	return out
}

type Message interface {
	SetPayload(payload Payload)
	String() string
	StringRaw() string
	StringHex() string
	Bytes() []byte
	GetPayload() Payload
	IsAck() bool
	SetCounter(c Counter)
	GetCounter() Counter
	CounterNext()
	HasRegisters([]byte) bool
}

type message struct {
	Start          [3]byte
	PayloadsLength int
	Counter        Counter
	Separator      byte
	Payload        Payload
	Checksum       byte
	End            byte
}

func NewEmptyMessage(c Counter) Message {
	m := EmptyMessage
	m.Counter = c
	return &m
}

func computeXorChecksum(data []byte) byte {
	var checksum byte = 0x00 // Initial value of checksum

	// XOR each byte in the data array with the current checksum
	for _, b := range data {
		checksum ^= b
	}

	return checksum
}

func (m *message) SetPayload(payload Payload) {
	m.Payload = payload
}

func (m *message) SetCounter(c Counter) {
	m.Counter = c
}

func (m *message) CounterNext() {
	if m.Counter == nil {
		return
	}
	m.Counter.Next()
}

func (m *message) GetCounter() Counter {
	return m.Counter
}

func (m *message) GetPayload() Payload {
	return m.Payload
}

func (m *message) Bytes() []byte {
	var out []byte
	if m.Counter == nil {
		m.Counter = NewCounter(0)
	}
	c := m.Counter.GetCurrent()
	out = appendBytes(out, []byte{0x00, 0x00, 0x00, 0x00, 0x00}...)
	out = appendBytes(out, c)
	out = appendBytes(out, m.Separator)
	out = appendBytes(out, m.Payload.Bytes()...)
	out = appendBytes(out, m.Checksum)
	out = appendBytes(out, m.End)
	msgLen := byte(len(out))

	var final []byte
	final = appendBytes(final, m.Start[:]...)
	final = appendBytes(final, msgLen)
	final = appendBytes(final, out...)

	ci := len(final) - 2
	final[ci] = computeXorChecksum(final[:ci])

	//fmt.Printf("%x\n", final)
	//fmt.Printf("%x %02x %02x %02x %02x %02x %02x %02x %02x %02x\n\n", final[:2], final[2], final[3], final[4:10], final[10], final[11:13], final[13], final[14:ci], final[ci], final[ci+1:])
	return final
}

func (m *message) String() string {
	var out bytes.Buffer

	t := m.Payload.GetType()
	out.WriteString(hex.EncodeToString(t[:]))
	out.WriteString(" ")
	//if m.Payload.GetType()[1]%2 == 0 {
	//	out.WriteString("  ❄️  ")
	//} else {
	//	out.WriteString("  🛜  ")
	//}

	for _, r := range m.Payload.GetRegisters() {
		out.WriteString(r.StringTxt(m.Payload.GetType()[0]))
		out.WriteString(" ")
	}
	return out.String()
}

func (m *message) StringRaw() string {
	c := m.Bytes()
	ci := len(c) - 2
	if c[13] == 0x00 {
		return fmt.Sprintf("%02x %02x %02x %02x %02x %02x %02x %02x",
			c[:3], c[3], c[4:10], c[10],
			c[11:13], c[13],
			c[ci], c[ci+1:])
	}
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%02x %02x %02x %02x ", c[:3], c[3], c[4:10], c[10]))
	rig := color.New(color.FgGreen)
	if c[12]%2 == 0 {
		rig = color.New(color.FgHiGreen)
	}
	out.WriteString(rig.Sprintf("%02x ", c[11:13]))
	out.WriteString(fmt.Sprintf("%02x ", c[13]))

	for _, r := range m.Payload.GetRegisters() {
		out.WriteString(r.String())
		out.WriteString(" ")
	}
	out.WriteString(fmt.Sprintf("%02x %02x", c[ci], c[ci+1:]))

	return out.String()
}

//func (m *message) StringRaw() string {
//	c := m.Bytes()
//	ci := len(c) - 2
//	if c[13] == 0x00 {
//		return fmt.Sprintf("%02x %02x %02x %02x %02x %02x %02x %02x",
//			c[:3], c[3], c[4:10], c[10],
//			c[11:13], c[13],
//			c[ci], c[ci+1:])
//
//	}
//	return fmt.Sprintf("%02x %02x %02x %02x %02x %02x %02x %02x %02x",
//		c[:3], c[3], c[4:10], c[10],
//		c[11:13], c[13], c[14:14+c[13]],
//		c[ci], c[ci+1:])
//}

func (m *message) StringHex() string {
	c := m.Bytes()
	return fmt.Sprintf("%x", c)
}

func (m *message) IsAck() bool {
	return m.Payload.IsAck()
}

func (m *message) HasRegisters(registerIds []byte) bool {
	for _, r := range registerIds {
		if !m.Payload.HasRegister(r) {
			log.Debugf("Didn't found register ID: %02x", r)
			return false
		}
	}
	return true
}

type Payload interface {
	AddRegister(register Register)
	AddRegisters(registers []Register)
	SetRegisters(registers []Register)
	SetRegistersFromString(registers string)
	String() string
	Bytes() []byte
	GetAck() Payload
	IsAck() bool
	GetType() MessageType
	GetRegisters() []Register
	HasRegister(register byte) bool
}

type payload struct {
	Type          MessageType
	PayloadLength int
	Registers     []Register
}

func NewEmptyPayload() Payload {
	return &payload{
		Type:          MessageType{},
		PayloadLength: 0,
		Registers:     []Register{},
	}
}

func NewPayload(t MessageType) Payload {
	return &payload{
		Type:          MessageType{t[0], t[1]},
		PayloadLength: 0,
		Registers:     []Register{},
	}
}

func NewPayloadFromString(t string) Payload {
	if len(t) > 4 {
		panic("length is too long")
	}
	h, err := hex.DecodeString(t)
	if err != nil {
		panic(err)
	}

	return NewPayload(MessageType(h))
}

func (p *payload) AddRegister(register Register) {
	p.Registers = append(p.Registers, register)
}

func (p *payload) AddRegisters(registers []Register) {
	for _, r := range registers {
		p.AddRegister(r)
	}
}

func (p *payload) SetRegisters(registers []Register) {
	p.Registers = registers
}

func (p *payload) SetRegistersFromString(registers string) {
	p.Registers = []Register{}
	registers = strings.Replace(registers, " ", "", -1)
	h, err := hex.DecodeString(registers)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(h)
	scanner := bufio.NewScanner(buf)
	scanner.Split(ScanRegister)
	for scanner.Scan() {
		line := scanner.Bytes()
		r := NewRegisterRaw(line)
		p.AddRegister(r)
	}
}

func (p *payload) Bytes() []byte {
	var final []byte
	final = appendBytes(final, p.Type[:]...)

	var out []byte
	for _, r := range p.Registers {
		out = append(out, r.Bytes()...)
	}
	lengthByte := byte(len(out))

	final = appendBytes(final, lengthByte)
	final = appendBytes(final, out...)
	return final
}

func (p *payload) String() string {
	c := p.Bytes()
	return hex.EncodeToString(c)
}

func (p *payload) IsAck() bool {
	return p.Type[1]%2 != 0
}

func (p *payload) GetAck() Payload {
	r := *p
	r.Type = MessageType{p.Type[0], p.Type[1] + 1}
	return &r
}

func (p *payload) GetType() MessageType {
	return p.Type
}

func (p *payload) GetRegisters() []Register {
	return p.Registers
}

func (p *payload) HasRegister(register byte) bool {
	for _, r := range p.Registers {
		if r.GetId() == register {
			return true
		}
	}
	return false
}

type Register interface {
	GetId() byte
	GetLen() int
	GetValue() []byte
	String() string
	StringRaw() string
	StringTxt(group byte) string
	Bytes() []byte
}

type register struct {
	RegisterId   byte
	RegisterSize byte
	Value        []byte
}

func NewRegisterFromString(id byte, value string) Register {
	valueB, err := hex.DecodeString(value)
	if err != nil {
		panic(fmt.Errorf("HEX decode error: (%02x) %s", id, value))
	}
	return NewRegister(id, valueB)
}

func NewRegister(id byte, value []byte) Register {
	length := len(value)
	return &register{
		RegisterId:   id,
		RegisterSize: byte(length),
		Value:        value,
	}
}

func NewRegisterRaw(b []byte) Register {
	id := b[0]
	s := b[1]
	if len(b) == 0 {
		return nil
	}
	if len(b) < 2 {
		return nil
	}
	if len(b) == 2 {
		return &register{
			RegisterId:   id,
			RegisterSize: s,
		}
	}

	v := b[2:]
	return &register{
		RegisterId:   id,
		RegisterSize: s,
		Value:        v,
	}
}

func NewRegisterEmpty(id string) Register {
	if len(id) > 2 {
		panic("id is too long")
	}
	idB, err := hex.DecodeString(id)
	if err != nil {
		panic(err)
	}

	return &register{
		RegisterId:   idB[0],
		RegisterSize: byte(0x00),
		Value:        []byte{},
	}
}

func (r *register) GetId() byte {
	return r.RegisterId
}

func (r *register) GetLen() int {
	return int(r.RegisterSize)
}

func (r *register) GetValue() []byte {
	return r.Value
}

func (r *register) Bytes() []byte {
	var out []byte
	out = append(out, r.RegisterId)
	out = append(out, r.RegisterSize)
	out = append(out, r.Value...)
	return out
}

func (r *register) StringTxt(group byte) string {
	var out bytes.Buffer
	ric := color.New(color.FgRed)
	rici := color.New(color.FgRed, color.Italic)
	rig := color.New(color.FgGreen)

	rd := (*RegisterDetails[group])[r.RegisterId]
	t := rd.txt
	if t == "" {
		t = ric.Sprintf("unknown_%02x_%02x", group, r.RegisterId)
	}
	if strings.Index(t, "_UNKNOWN_") >= 0 {
		t = rici.Sprint(t)
	}

	out.WriteString(t)
	out.WriteString(":")
	out.WriteString(rig.Sprintf("%02x", r.GetValue()))
	return out.String()
}

func (r *register) String() string {
	var out bytes.Buffer
	ric := color.New(color.FgRed)
	rig := color.New(color.FgGreen)
	rii := color.New(color.Italic)
	out.WriteString(ric.Sprintf("%02x", r.RegisterId))
	out.WriteString(rii.Sprintf("%02x", r.RegisterSize))
	if r.GetLen() > 0 {
		out.WriteString(rig.Sprintf("%02x", r.Value))
	}
	return out.String()
}

func (r *register) StringRaw() string {
	c := r.Bytes()
	return hex.EncodeToString(c)
}
