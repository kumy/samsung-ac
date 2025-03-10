package proto

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type registerSuite struct {
	suite.Suite
}

func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(registerSuite))
}

func (s *registerSuite) TestNewRegister() {
	testCases := []struct {
		name     string
		register byte
		value    string
		want     string
	}{
		{
			name:     "Empty: enable",
			register: REGISTER_12_AC_FUN_ENABLE,
			value:    "",
			want:     "0100",
		},
		{
			name:     "REGISTER_12_AC_FUN_ENABLE: enable",
			register: REGISTER_12_AC_FUN_ENABLE,
			value:    ENABLE_0F,
			want:     "01010f",
		},
		{
			name:     "REGISTER_12_AC_FUN_ENABLE: enable",
			register: REGISTER_12_AC_FUN_ENABLE,
			value:    DISABLE_F0,
			want:     "0101f0",
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := NewRegisterFromString(tc.register, tc.value)
			s.Equal(tc.want, got.String())
		})
	}
}

type payloadSuite struct {
	suite.Suite
}

func TestPayloadSuite(t *testing.T) {
	suite.Run(t, new(payloadSuite))
}

func (s *payloadSuite) TestNewPayload() {
	testCases := []struct {
		name      string
		registers []Register
		want      string
	}{
		{
			name:      "AC_FUN_ENABLE: enable",
			registers: []Register{AC_FUN_ENABLE(true)},
			want:      "12040301010f",
		},
		{
			name:      "AC_FUN_ENABLE: enable with commit",
			registers: []Register{AC_FUN_ENABLE(true), AC_BIP},
			want:      "12040601010f7401f0",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			p := NewPayloadFromString("1204")
			p.AddRegisters(tc.registers)
			s.Equal(tc.want, p.(*payload).String())
		})
	}
}

type messageSuite struct {
	suite.Suite
}

func TestMessageSuite(t *testing.T) {
	suite.Run(t, new(messageSuite))
}

func (s *messageSuite) TestNewMessage() {
	testCases := []struct {
		name      string
		cmd       string
		counter   uint64
		registers []Register
		want      string
	}{
		{
			name:      "Empty",
			cmd:       "1204",
			registers: []Register{},
			want:      "d0c002 0c 000000000000 fe 1204 00 f6 e0",
		},
		{
			name:      "AC_FUN_ENABLE: enable",
			cmd:       "1204",
			registers: []Register{AC_FUN_ENABLE(true), AC_BIP},
			want:      "d0c002 12 000000000000 fe 1204 06 01010f 7401f0 64 e0",
		},
		{
			name:      "AC_SG_WIFI__ENABLE: enable",
			cmd:       "1404",
			registers: []Register{AC_SG_WIFI(true)},
			want:      "d0c002 0f 000000000000 fe 1404 03 37010f c9 e0",
		},
		{
			name:      "AC_SG_WIFI__ENABLE: enable",
			cmd:       "1404",
			counter:   1,
			registers: []Register{AC_UNKNOWN_17, AC_UNKNOWN_18, AC_UNKNOWN_19, AC_UNKNOWN_FD},
			want:      "d0c002 18 000000000001 fe 1404 0c 170115 180105 190127 fd0102 37 e0",
		},
		{
			name:      "AC_SG_WIFI_VENDOR",
			cmd:       "1404",
			counter:   2,
			registers: []Register{AC_SG_VENDER01, AC_SG_VENDER02, AC_SG_VENDER03, AC_SG_MACHIGH, AC_SG_MACMID, AC_SG_MACLOW},
			want:      "d0c002 1e 000000000002 fe 1404 12 fa01f8 fb0104 fc012e f701d8 f8013e f901a0 6d e0",
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			m := NewEmptyMessage(NewCounter(byte(tc.counter)))
			p := NewPayloadFromString(tc.cmd)
			p.AddRegisters(tc.registers)
			m.SetPayload(p)
			s.Equal(tc.want, m.(*message).StringRaw())
		})
	}
}

func (s *messageSuite) TestSetRegistersFromString() {
	testCases := []struct {
		name      string
		cmd       string
		counter   uint64
		registers string
		want      string
	}{
		{
			name:      "Empty",
			cmd:       "1204",
			registers: "",
			want:      "d0c002 0c 000000000000 fe 1204 00 f6 e0",
		},
		{
			name:      "AC_FUN_ENABLE: enable",
			cmd:       "1204",
			registers: "01010f 7401f0",
			want:      "d0c002 12 000000000000 fe 1204 06 01010f 7401f0 64 e0",
		},
		{
			name:      "AC_SG_WIFI__ENABLE: enable",
			cmd:       "1404",
			registers: "37010f",
			want:      "d0c002 0f 000000000000 fe 1404 03 37010f c9 e0",
		},
		{
			name:      "AC_SG_WIFI__ENABLE: enable",
			cmd:       "1404",
			counter:   1,
			registers: "170115 180105 190127 fd0102",
			want:      "d0c002 18 000000000001 fe 1404 0c 170115 180105 190127 fd0102 37 e0",
		},
		{
			name:      "AC_SG_WIFI_VENDOR",
			cmd:       "1404",
			counter:   2,
			registers: "fa01f8 fb0104 fc012e f701d8 f8013e f901a0",
			want:      "d0c002 1e 000000000002 fe 1404 12 fa01f8 fb0104 fc012e f701d8 f8013e f901a0 6d e0",
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			m := NewEmptyMessage(NewCounter(byte(tc.counter)))
			p := NewPayloadFromString(tc.cmd)
			p.SetRegistersFromString(tc.registers)
			m.SetPayload(p)
			s.Equal(tc.want, m.(*message).StringRaw())
		})
	}
}

//func (s *messageSuite) TestAnswerFromAC() {
//	testCases := []struct {
//		name      string
//		cmd       string
//		counter   uint64
//		registers []Register
//		want      string
//	}{
//		{
//			name:    "AC_SG_WIFI_VENDOR",
//			cmd:     "1302",
//			counter: 8,
//			registers: []Register{
//				NewRegisterEmpty(REGISTER_13_AC_ADD_AUTOCLEAN), NewRegisterEmpty(REGISTER_13_AC_ADD_SETKWH),
//				NewRegisterEmpty(REGISTER_13_AC_ADD_CLEAR_FILTER_ALARM), NewRegisterEmpty(REGISTER_13_AC_ADD_STARTWPS),
//				NewRegisterEmpty(REGISTER_13_AC_ADD_SPI), NewRegisterEmpty(REGISTER_13_AC_OUTDOOR_TEMP),
//				NewRegisterEmpty(REGISTER_13_AC_COOL_CAPABILITY), NewRegisterEmpty(REGISTER_13_AC_WARM_CAPABILITY),
//			},
//			want: "d0c002 1c 000000000008 fe 1303 1a 3201234001ff4401f043010f7501f0760144770244780250 32 e0",
//		},
//	}
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			m := NewEmptyMessage(NewCounter(tc.counter))
//			p := NewPayloadFromString(tc.cmd)
//			p.AddRegisters(tc.registers)
//			m.SetPayload(p)
//
//			got := m
//
//			s.Equal(tc.want, got.String())
//		})
//	}
//}
