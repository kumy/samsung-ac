package simulator

import (
	"fmt"
	"github.com/kumy/samsung-ac/proto"
)

type stateMachine struct {
	Groups map[byte]*registerGroupVal `jsonapi:"relation,groups"`
}

type registerGroupVal struct {
	Register map[byte][]byte `jsonapi:"attr,register"`
}

type StateMachine interface {
	Get(group byte, key byte) ([]byte, error)
	Set(group byte, key byte, value []byte) error
	SetFromRegister(group byte, register proto.Register) error
	Answer(msg proto.Message) proto.Message
	BuildMessageFromRegisters(t proto.MessageType, group byte, registers []byte) proto.Message
}

func NewStateMachine() StateMachine {
	return &stateMachine{
		Groups: map[byte]*registerGroupVal{
			proto.REGISTER_GROUP_12: &registerGroupVal{
				Register: map[byte][]byte{
					proto.REGISTER_12_AC_FUN_ENABLE:     {0xf0}, // Enable/Disable
					proto.REGISTER_12_AC_FUN_POWER:      {0xf0}, // On Off
					proto.REGISTER_12_AC_FUN_UNKNOWN_41: {0x32}, // AC_FUN_SUPPORTED?
					proto.REGISTER_12_AC_FUN_OPMODE:     {0x42}, // MODE_COOL/MODE_DRY/...
					proto.REGISTER_12_AC_FUN_COMODE:     {0x12}, // COMODE_OFF/COMODE_TURBO/...
					proto.REGISTER_12_AC_FUN_WINDLEVEL:  {0x12}, // WIND_LEVEL_AUTO/WIND_LEVEL_LOW/...
					proto.REGISTER_12_AC_FUN_DIRECTION:  {0xc2}, // WIND_DIRECTION_OFF/WIND_DIRECTION_INDIRECT/...
					proto.REGISTER_12_AC_FUN_UNKNOWN_EA: {0xfe},
					proto.REGISTER_12_AC_FUN_TEMP_SET:   {0x15}, // to int
					proto.REGISTER_12_AC_FUN_TEMP_NOW:   {0x13}, // to int
					proto.REGISTER_12_AC_FUN_SLEEP:      {0x00}, // to int?
					proto.REGISTER_12_AC_FUN_BIP:        {0xf0},
					proto.REGISTER_12_AC_FUN_ERROR:      {0x00, 0x00, 0x00, 0x00},
				},
			},
			proto.REGISTER_GROUP_13: &registerGroupVal{
				Register: map[byte][]byte{
					proto.REGISTER_13_AC_ADD_AUTOCLEAN:          {0x22}, // On/Off
					proto.REGISTER_13_AC_ADD_SETKWH:             {0xff}, // 255 to int?
					proto.REGISTER_13_AC_ADD_CLEAR_FILTER_ALARM: {0x00}, // 240
					proto.REGISTER_13_AC_ADD_STARTWPS:           {0x0f},
					proto.REGISTER_13_AC_ADD_SPI:                {0xf0},
					proto.REGISTER_13_AC_OUTDOOR_TEMP:           {0x44},
					proto.REGISTER_13_AC_COOL_CAPABILITY:        {0x00, 0x44},
					proto.REGISTER_13_AC_WARM_CAPABILITY:        {0x00, 0x50},
				},
			},
			proto.REGISTER_GROUP_14: &registerGroupVal{
				Register: map[byte][]byte{
					proto.REGISTER_14_AC_ADD2_UNKNOWN_17:      {0x15},
					proto.REGISTER_14_AC_ADD2_UNKNOWN_18:      {0x05},
					proto.REGISTER_14_AC_ADD2_UNKNOWN_19:      {0x27},
					proto.REGISTER_14_AC_ADD2_UNKNOWN_FD:      {0x02},
					proto.REGISTER_14_AC_ADD2_USEDWATT:        {0xfe},
					proto.REGISTER_14_AC_SG_WIFI:              {0x0f},
					proto.REGISTER_14_AC_SG_INTERNET:          {0xf0},
					proto.REGISTER_14_AC_ADD2_VERSION:         {0xfe},
					proto.REGISTER_14_AC_ADD2_PANEL_VERSION:   {0x15, 0x02, 0x24},
					proto.REGISTER_14_AC_ADD2_OUT_VERSION:     {0x13, 0x07, 0x09},
					proto.REGISTER_14_AC_FUN_MODEL:            {0x05},
					proto.REGISTER_14_AC_ADD2_OPTIONCODE:      {0xd0, 0xb8},
					proto.REGISTER_14_AC_SG_MACHIGH:           {0xd8},
					proto.REGISTER_14_AC_SG_MACMID:            {0x3e},
					proto.REGISTER_14_AC_SG_MACLOW:            {0xa0},
					proto.REGISTER_14_AC_ADD2_USEDPOWER:       {0x00, 0x00, 0xe6, 0x35},
					proto.REGISTER_14_AC_ADD2_USEDTIME:        {0x00, 0x04, 0x0d, 0xcb},
					proto.REGISTER_14_AC_ADD2_CLEAR_POWERTIME: {0xfe},
					proto.REGISTER_14_AC_ADD2_FILTERTIME:      {0x03},
					proto.REGISTER_14_AC_ADD2_FILTER_USE_TIME: {0x27, 0x10},
					proto.REGISTER_14_AC_SG_VENDER01:          {0x27, 0xf8},
					proto.REGISTER_14_AC_SG_VENDER02:          {0x27, 0x04},
					proto.REGISTER_14_AC_SG_VENDER03:          {0x27, 0x2e},
				},
			},
		},
	}
}

func (sm *stateMachine) Get(group byte, key byte) ([]byte, error) {
	ret, ok := sm.Groups[group]
	if !ok {
		return nil, fmt.Errorf("unknown group %d", group)
	}
	//k := hex.EncodeToString([]byte{key})
	val, ok := ret.Register[key]
	if !ok {
		return nil, fmt.Errorf("unknown register key %x", key)
	}
	return val, nil
}

func (sm *stateMachine) Answer(msg proto.Message) proto.Message {
	if msg.IsAck() {
		return nil
	}
	sm.receive(msg)

	var newRegisters []proto.Register
	payload := msg.GetPayload()
	switch payload.GetType() {
	case proto.MSG_1202:
		for _, r := range payload.GetRegisters() {
			val, err := sm.Get(payload.GetType()[0], r.GetId())
			if err != nil {
				return nil
			}
			p := proto.NewRegister(r.GetId(), val)
			newRegisters = append(newRegisters, p)
		}
	case proto.MSG_1302:
		for _, r := range payload.GetRegisters() {
			val, err := sm.Get(payload.GetType()[0], r.GetId())
			if err != nil {
				return nil
			}
			p := proto.NewRegister(r.GetId(), val)
			newRegisters = append(newRegisters, p)
		}
	case proto.MSG_1402:
		for _, r := range payload.GetRegisters() {
			val, err := sm.Get(payload.GetType()[0], r.GetId())
			if err != nil {
				return nil
			}
			p := proto.NewRegister(r.GetId(), val)
			newRegisters = append(newRegisters, p)
		}
	default:
		newRegisters = payload.GetRegisters()
	}

	r := proto.NewEmptyMessage(msg.GetCounter())
	t := payload.GetType()
	newPayload := proto.NewPayload(proto.MessageType(t[:]))
	newPayload.SetRegisters(newRegisters)
	r.SetPayload(newPayload.GetAck())
	return r
}

func (sm *stateMachine) Set(group byte, key byte, value []byte) error {
	sm.Groups[group].Register[key] = value
	return nil
}

func (sm *stateMachine) SetFromRegister(group byte, register proto.Register) error {
	return sm.Set(group, register.GetId(), register.GetValue())
}

func (sm *stateMachine) receive(msg proto.Message) {
	payload := msg.GetPayload()
	for _, reg := range payload.GetRegisters() {
		if reg.GetLen() == 0 {
			continue
		}
		_ = sm.SetFromRegister(payload.GetType()[0], reg)
	}
}

func (sm *stateMachine) BuildMessageFromRegisters(t proto.MessageType, group byte, registers []byte) proto.Message {
	p := proto.NewPayload(t)
	for _, reg := range registers {
		val, err := sm.Get(group, reg)
		if err != nil {
			continue
		}
		r := proto.NewRegister(reg, val)
		p.AddRegister(r)
	}
	m := proto.NewEmptyMessage(nil)
	m.SetPayload(p)
	return m
}
