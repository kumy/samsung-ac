package proto

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type constsSuite struct {
	suite.Suite
}

func TestConstsSuite(t *testing.T) {
	suite.Run(t, new(constsSuite))
}

func (s *constsSuite) TestAC_FUN_ENABLE() {
	testCases := []struct {
		name      string
		enable    bool
		wantId    byte
		wantLen   int
		wantValue []byte
	}{
		{
			name:      "true",
			enable:    true,
			wantId:    0x01,
			wantLen:   1,
			wantValue: []byte{0x0f},
		},
		{
			name:      "false",
			enable:    false,
			wantId:    0x01,
			wantLen:   1,
			wantValue: []byte{0xf0},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := AC_FUN_ENABLE(tc.enable)
			s.Equal(tc.wantId, got.GetId())
			s.Equal(tc.wantLen, got.GetLen())
			s.Equal(tc.wantValue, got.GetValue())
		})
	}
}

func (s *constsSuite) TestAC_FUN_POWER() {
	testCases := []struct {
		name   string
		enable bool
		want   []byte
	}{
		{
			name:   "true",
			enable: true,
			want:   []byte{0x0f},
		},
		{
			name:   "false",
			enable: false,
			want:   []byte{0xf0},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := AC_FUN_ENABLE(tc.enable)
			s.Equal(tc.want, got.GetValue())
		})
	}
}

func (s *constsSuite) TestAC_FUN_COMODE() {
	testCases := []struct {
		name      string
		mode      string
		wantId    byte
		wantLen   int
		wantValue []byte
	}{
		{
			name:      "Off",
			mode:      COMODE_OFF,
			wantId:    0x44,
			wantLen:   1,
			wantValue: []byte{0x12},
		},
		{
			name:      "TurboMode",
			mode:      COMODE_TURBO,
			wantId:    0x44,
			wantLen:   1,
			wantValue: []byte{0x22},
		},
		{
			name:      "Smart",
			mode:      COMODE_SMART,
			wantId:    0x44,
			wantLen:   1,
			wantValue: []byte{0x32},
		},
		{
			name:      "Sleep",
			mode:      COMODE_SLEEP,
			wantId:    0x44,
			wantLen:   1,
			wantValue: []byte{0x42},
		},
		{
			name:      "Quiet",
			mode:      COMODE_QUIET,
			wantId:    0x44,
			wantLen:   1,
			wantValue: []byte{0x52},
		},
		{
			name:      "SoftCool",
			mode:      COMODE_SOFT_COOL,
			wantId:    0x44,
			wantLen:   1,
			wantValue: []byte{0x62},
		},
		{
			name:      "WindMode1",
			mode:      COMODE_WIND_MODE_1,
			wantId:    0x44,
			wantLen:   1,
			wantValue: []byte{0x82},
		},
		{
			name:      "WindMode2",
			mode:      COMODE_WIND_MODE_2,
			wantId:    0x44,
			wantLen:   1,
			wantValue: []byte{0x92},
		},
		{
			name:      "WindMode3",
			mode:      COMODE_WIND_MODE_3,
			wantId:    0x44,
			wantLen:   1,
			wantValue: []byte{0xa2},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := AC_FUN_COMODE(tc.mode)
			s.Equal(tc.wantId, got.GetId())
			s.Equal(tc.wantLen, got.GetLen())
			s.Equal(tc.wantValue, got.GetValue())
		})
	}
}

func (s *constsSuite) TestAC_FUN_OPMODE() {
	testCases := []struct {
		name      string
		mode      string
		wantId    byte
		wantLen   int
		wantValue []byte
	}{
		{
			name:      "Cool",
			mode:      MODE_COOL,
			wantId:    0x43,
			wantLen:   1,
			wantValue: []byte{0x12},
		},
		{
			name:      "Dry",
			mode:      MODE_DRY,
			wantId:    0x43,
			wantLen:   1,
			wantValue: []byte{0x22},
		},
		{
			name:      "Wind",
			mode:      MODE_WIND,
			wantId:    0x43,
			wantLen:   1,
			wantValue: []byte{0x32},
		},
		{
			name:      "Heat",
			mode:      MODE_HEAT,
			wantId:    0x43,
			wantLen:   1,
			wantValue: []byte{0x42},
		},
		{
			name:      "Heat",
			mode:      MODE_AUTO,
			wantId:    0x43,
			wantLen:   1,
			wantValue: []byte{0xe2},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := AC_FUN_OPMODE(tc.mode)
			s.Equal(tc.wantId, got.GetId())
			s.Equal(tc.wantLen, got.GetLen())
			s.Equal(tc.wantValue, got.GetValue())
		})
	}
}

func (s *constsSuite) TestAC_FUN_TEMPSET() {
	testCases := []struct {
		name string
		temp int
		want []byte
	}{
		{
			name: "16 min",
			temp: 16,
			want: []byte{0x10},
		},
		{
			name: "28 max",
			temp: 28,
			want: []byte{0x1c},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := AC_FUN_TEMPSET(tc.temp)
			s.Equal(tc.want, got.GetValue())
		})
	}
}

func (s *constsSuite) TestAC_FUN_WINDLEVEL() {
	testCases := []struct {
		name      string
		mode      string
		wantId    byte
		wantLen   int
		wantValue []byte
	}{
		{
			name:      "Auto",
			mode:      WIND_LEVEL_AUTO,
			wantId:    0x62,
			wantLen:   1,
			wantValue: []byte{0x00},
		},
		{
			name:      "Low",
			mode:      WIND_LEVEL_LOW,
			wantId:    0x62,
			wantLen:   1,
			wantValue: []byte{0x12},
		},
		{
			name:      "Mid",
			mode:      WIND_LEVEL_MID,
			wantId:    0x62,
			wantLen:   1,
			wantValue: []byte{0x14},
		},
		{
			name:      "High",
			mode:      WIND_LEVEL_HIGH,
			wantId:    0x62,
			wantLen:   1,
			wantValue: []byte{0x16},
		},
		{
			name:      "Turbo",
			mode:      WIND_LEVEL_TURBO,
			wantId:    0x62,
			wantLen:   1,
			wantValue: []byte{0x18},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := AC_FUN_WIND_LEVEL(tc.mode)
			s.Equal(tc.wantId, got.GetId())
			s.Equal(tc.wantLen, got.GetLen())
			s.Equal(tc.wantValue, got.GetValue())
		})
	}
}

func (s *constsSuite) TestAC_FUN_WIND_DIRECTION() {
	testCases := []struct {
		name      string
		mode      string
		wantId    byte
		wantLen   int
		wantValue []byte
	}{
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_OFF,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0x12},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_INDIRECT,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0x21},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_DIRECT,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0x31},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_CENTER,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0x41},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_WIDE,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0x51},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_LEFT,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0x61},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_RIGHT,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0x71},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_LONG,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0x81},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_SWING_UD,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0x82},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_SWING_LR,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0xa2},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_ROTATION,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0xb2},
		},
		{
			name:      "Auto",
			mode:      WIND_DIRECTION_FIXED,
			wantId:    0x63,
			wantLen:   1,
			wantValue: []byte{0xc2},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := AC_FUN_WIND_DIRECTION(tc.mode)
			s.Equal(tc.wantId, got.GetId())
			s.Equal(tc.wantLen, got.GetLen())
			s.Equal(tc.wantValue, got.GetValue())
		})
	}
}

func (s *constsSuite) TestAC_FUN_SLEEP() {
	testCases := []struct {
		name      string
		mode      bool
		wantId    byte
		wantLen   int
		wantValue []byte
	}{
		{
			name:      "true",
			mode:      true,
			wantId:    0x73,
			wantLen:   1,
			wantValue: []byte{0x00},
		},
		{
			name:      "false",
			mode:      false,
			wantId:    0x73,
			wantLen:   1,
			wantValue: []byte{0xff},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := AC_FUN_SLEEP(tc.mode)
			s.Equal(tc.wantId, got.GetId())
			s.Equal(tc.wantLen, got.GetLen())
			s.Equal(tc.wantValue, got.GetValue())
		})
	}
}

func (s *constsSuite) TestAC_FUN_ERROR() {
	testCases := []struct {
		name      string
		msg       string
		wantId    byte
		wantLen   int
		wantValue []byte
	}{
		{
			name:      "NULL",
			msg:       "NULL",
			wantId:    0xf7,
			wantLen:   4,
			wantValue: []byte{0x4e, 0x55, 0x4c, 0x4c},
		},
		{
			name:      "TEST",
			msg:       "TEST",
			wantId:    0xf7,
			wantLen:   4,
			wantValue: []byte{0x54, 0x45, 0x53, 0x54},
		},
		{
			name:      "MY_ERROR_MESSAGE",
			msg:       "MY ERROR MESSAGE",
			wantId:    0xf7,
			wantLen:   16,
			wantValue: []byte{0x4d, 0x59, 0x20, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x20, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45},
		},
		{
			name:      "0123456789012345678901234567890123456789012345678901234567890123456789",
			msg:       "0123456789012345678901234567890123456789012345678901234567890123456789",
			wantId:    0xf7,
			wantLen:   70,
			wantValue: []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := AC_FUN_ERROR(tc.msg)
			s.Equal(tc.wantId, got.GetId())
			s.Equal(tc.wantLen, got.GetLen())
			s.Equal(tc.wantValue, got.GetValue())
		})
	}
}
