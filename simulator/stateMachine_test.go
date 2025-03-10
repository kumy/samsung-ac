package simulator

//
//import (
//	"github.com/stretchr/testify/suite"
//	"testing"
//)
//
//type stateMachineSuite struct {
//	suite.Suite
//}
//
//func TestStateMachineSuite(t *testing.T) {
//	suite.Run(t, new(stateMachineSuite))
//}
//
//func (sms *stateMachineSuite) TestAnswer() {
//	testCases := []struct {
//		name      string
//		cmd       string
//		counter   uint64
//		registers []Register
//		want      string
//	}{
//		{
//			name:      "AC_FUN_ENABLE: enable",
//			cmd:       "1204",
//			registers: []Register{AC_FUN_ENABLE(true), AC_BIP},
//			want:      "d0c0 02 12 000000000000 fe 1205 06 01010f7401f0 65 e0",
//		},
//		{
//			name:      "AC_SG_WIFI__ENABLE: enable",
//			cmd:       "1404",
//			registers: []Register{AC_SG_WIFI(true)},
//			want:      "d0c0 02 0f 000000000000 fe 1405 03 37010f c8 e0",
//		},
//		{
//			name:      "AC_SG_WIFI__ENABLE: enable",
//			cmd:       "1404",
//			counter:   1,
//			registers: []Register{AC_UNKNOWN_17, AC_UNKNOWN_18, AC_UNKNOWN_19, AC_UNKNOWN_FD},
//			want:      "d0c0 02 18 000000000001 fe 1405 0c 170115180105190127fd0102 36 e0",
//		},
//		{
//			name:      "AC_SG_WIFI_VENDOR",
//			cmd:       "1404",
//			counter:   2,
//			registers: []Register{AC_SG_VENDER01, AC_SG_VENDER02, AC_SG_VENDER03, AC_SG_MACHIGH, AC_SG_MACMID, AC_SG_MACLOW},
//			want:      "d0c0 02 1e 000000000002 fe 1405 12 fa01f8fb0104fc012ef701d8f8013ef901a0 6c e0",
//		},
//	}
//	for _, tc := range testCases {
//		sms.Run(tc.name, func() {
//			sm := NewStateMachine()
//			m := NewEmptyMessage(NewCounter(tc.counter))
//			p := NewPayloadFromString(tc.cmd)
//			p.AddRegisters(tc.registers)
//			m.SetPayload(p)
//
//			got := sm.Answer(m)
//
//			sms.Equal(tc.want, got.String())
//		})
//	}
//}
//
//func (sms *stateMachineSuite) TestKnownMessageSequence() {
//	testCases := []struct {
//		name      string
//		cmd       string
//		counter   uint64
//		registers []Register
//		want      string
//	}{
//		{
//			name:      "AC_FUN_ENABLE: enable",
//			cmd:       "1204",
//			registers: []Register{AC_FUN_ENABLE(true), AC_BIP},
//			want:      "d0c0 02 12 000000000000 fe 1205 06 01010f7401f0 65 e0",
//		},
//		{
//			name:      "AC_SG_WIFI__ENABLE: enable",
//			cmd:       "1404",
//			registers: []Register{AC_UNKNOWN_17, AC_UNKNOWN_18, AC_UNKNOWN_19, AC_UNKNOWN_FD},
//			want:      "d0c0 02 18 000000000001 fe 1405 0c 170115180105190127fd0102 36 e0",
//		},
//		{
//			name:      "AC_SG_WIFI_VENDOR",
//			cmd:       "1404",
//			registers: []Register{AC_SG_VENDER01, AC_SG_VENDER02, AC_SG_VENDER03, AC_SG_MACHIGH, AC_SG_MACMID, AC_SG_MACLOW},
//			want:      "d0c0 02 1e 000000000002 fe 1405 12 fa01f8fb0104fc012ef701d8f8013ef901a0 6c e0",
//		},
//		{
//			name:      "AC_SG_WIFI__ENABLE: enable",
//			cmd:       "1404",
//			registers: []Register{AC_UNKNOWN_17, AC_UNKNOWN_18, AC_UNKNOWN_19, AC_UNKNOWN_FD},
//			want:      "d0c0 02 18 000000000003 fe 1405 0c 170115180105190127fd0102 34 e0",
//		},
//		{
//			name:      "AC_SG_WIFI_VENDOR",
//			cmd:       "1404",
//			registers: []Register{AC_SG_VENDER01, AC_SG_VENDER02, AC_SG_VENDER03, AC_SG_MACHIGH, AC_SG_MACMID, AC_SG_MACLOW},
//			want:      "d0c0 02 1e 000000000004 fe 1405 12 fa01f8fb0104fc012ef701d8f8013ef901a0 6a e0",
//		},
//		{
//			name:      "AC_SG_WIFI__ENABLE: enable",
//			cmd:       "1404",
//			registers: []Register{AC_UNKNOWN_17, AC_UNKNOWN_18, AC_UNKNOWN_19, AC_UNKNOWN_FD},
//			want:      "d0c0 02 18 000000000005 fe 1405 0c 170115180105190127fd0102 32 e0",
//		},
//		{
//			name:      "AC_SG_WIFI_VENDOR",
//			cmd:       "1404",
//			registers: []Register{AC_SG_VENDER01, AC_SG_VENDER02, AC_SG_VENDER03, AC_SG_MACHIGH, AC_SG_MACMID, AC_SG_MACLOW},
//			want:      "d0c0 02 1e 000000000006 fe 1405 12 fa01f8fb0104fc012ef701d8f8013ef901a0 68 e0",
//		},
//		{
//			name: "AC_SG_WIFI_VENDOR",
//			cmd:  "1206",
//			registers: []Register{AC_FUN_POWER(false), AC_UNKNOWN_41, AC_FUN_OPMODE(MODE_AUTO), AC_FUN_COMODE(COMODE_OFF),
//				AC_FUN_WIND_LEVEL(WIND_LEVEL_AUTO), AC_FUN_WIND_DIRECTION(WIND_DIRECTION_FIXED), AC_UNKNOWN_EA,
//				AC_FUN_TEMPSET(21), AC_FUN_TEMP_NOW(23), AC_FUN_SLEEP(true), AC_FUN_ERROR("")},
//			want: "d0c0 02 30 000000000007 fe 1207 24 0201f04101324301e24401126201006301c2ea01fe5a01155c0117730100f70400000000 cf e0",
//		},
//		{
//			name: "AC_SG_WIFI_VENDOR",
//			cmd:  "1302",
//			registers: []Register{
//				NewRegisterEmpty(REGISTER_13_AC_ADD_AUTOCLEAN), NewRegisterEmpty(REGISTER_13_AC_ADD_SETKWH),
//				NewRegisterEmpty(REGISTER_13_AC_ADD_CLEAR_FILTER_ALARM), NewRegisterEmpty(REGISTER_13_AC_ADD_STARTWPS),
//				NewRegisterEmpty(REGISTER_13_AC_ADD_SPI), NewRegisterEmpty(REGISTER_13_AC_OUTDOOR_TEMP),
//				NewRegisterEmpty(REGISTER_13_AC_COOL_CAPABILITY), NewRegisterEmpty(REGISTER_13_AC_WARM_CAPABILITY),
//			},
//			want: "d0c0 02 1c 000000000008 fe 1303 10 32004000440043007500760077007800 81 e0",
//		},
//		{
//			name: "AC_SG_WIFI_VENDOR",
//			cmd:  "1402",
//			registers: []Register{
//				NewRegisterEmpty(REGISTER_14_AC_ADD2_USEDWATT), NewRegisterEmpty(REGISTER_14_AC_ADD2_VERSION),
//				NewRegisterEmpty(REGISTER_14_AC_ADD2_PANEL_VERSION), NewRegisterEmpty(REGISTER_14_AC_ADD2_OUT_VERSION),
//				NewRegisterEmpty(REGISTER_14_AC_FUN_MODEL), NewRegisterEmpty(REGISTER_14_AC_ADD2_OPTIONCODE),
//				NewRegisterEmpty(REGISTER_14_AC_ADD2_USEDPOWER), NewRegisterEmpty(REGISTER_14_AC_ADD2_USEDTIME),
//				NewRegisterEmpty(REGISTER_14_AC_ADD2_CLEAR_POWERTIME), NewRegisterEmpty(REGISTER_14_AC_ADD2_FILTERTIME),
//				NewRegisterEmpty(REGISTER_14_AC_ADD2_FILTER_USE_TIME),
//			},
//			want: "d0c0 02 39 000000000009 fe 1403 2d 3201fef601fef403150224f303130709f501053902d0b8e0040000e635e40400040dcbe801fee90103e6022710 93 e0",
//		},
//		{
//			name: "AC_SG_WIFI",
//			cmd:  "1404",
//			registers: []Register{
//				AC_SG_WIFI(true),
//			},
//			want: "d0c0 02 0f 00000000000a fe 1405 03 37010f c2 e0",
//		},
//		{
//			name: "AC_SG_WIFI",
//			cmd:  "1404",
//			registers: []Register{
//				AC_UNKNOWN_17, AC_UNKNOWN_18, AC_UNKNOWN_19, AC_UNKNOWN_FD,
//			},
//			want: "d0c0 02 18 00000000000b fe 1405 0c 170115180105190127fd0102 3c e0",
//		},
//		{
//			name: "AC_SG_WIFI",
//			cmd:  "1402",
//			registers: []Register{
//				NewRegisterEmpty(REGISTER_14_AC_ADD2_USEDWATT), NewRegisterEmpty(REGISTER_14_AC_ADD2_VERSION),
//				NewRegisterEmpty(REGISTER_14_AC_ADD2_PANEL_VERSION), NewRegisterEmpty(REGISTER_14_AC_ADD2_OUT_VERSION),
//				NewRegisterEmpty(REGISTER_14_AC_FUN_MODEL), NewRegisterEmpty(REGISTER_14_AC_ADD2_OPTIONCODE),
//				NewRegisterEmpty(REGISTER_14_AC_ADD2_USEDPOWER), NewRegisterEmpty(REGISTER_14_AC_ADD2_USEDTIME),
//				NewRegisterEmpty(REGISTER_14_AC_ADD2_CLEAR_POWERTIME), NewRegisterEmpty(REGISTER_14_AC_ADD2_FILTERTIME),
//				NewRegisterEmpty(REGISTER_14_AC_ADD2_FILTER_USE_TIME),
//			},
//			want: "d0c0 02 39 00000000000c fe 1403 2d 3201fef601fef403150224f303130709f501053902d0b8e0040000e635e40400040dcbe801fee90103e6022710 96 e0",
//		},
//		{
//			name:      "AC_SG_WIFI_VENDOR",
//			cmd:       "1404",
//			registers: []Register{AC_SG_VENDER01, AC_SG_VENDER02, AC_SG_VENDER03, AC_SG_MACHIGH, AC_SG_MACMID, AC_SG_MACLOW},
//			want:      "d0c0 02 1e 00000000000d fe 1405 12 fa01f8fb0104fc012ef701d8f8013ef901a0 63 e0",
//		},
//		{
//			name: "AC_SG_WIFI",
//			cmd:  "1404",
//			registers: []Register{
//				AC_SG_INTERNET(false),
//			},
//			want: "d0c0 02 0f 00000000000e fe 1405 03 3801f0 36 e0",
//		},
//		{
//			name: "AC_SG_WIFI",
//			cmd:  "1404",
//			registers: []Register{
//				AC_SG_WIFI(true),
//			},
//			want: "d0c0 02 0f 00000000000f fe 1405 03 37010f c7 e0",
//		},
//	}
//	c := NewCounter(0)
//	sm := NewStateMachine()
//	for _, tc := range testCases {
//		sms.Run(tc.name, func() {
//			m := NewEmptyMessage(c)
//			p := NewPayloadFromString(tc.cmd)
//			p.AddRegisters(tc.registers)
//			m.SetPayload(p)
//
//			got := sm.Answer(m)
//
//			sms.Equal(tc.want, got.String())
//			c.Next()
//		})
//	}
//}
