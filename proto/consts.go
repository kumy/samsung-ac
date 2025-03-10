package proto

import (
	"fmt"
	"strconv"
)

const ENABLE_0F = "0f"
const DISABLE_F0 = "f0"

const ENABLE_00 = "00"
const DISABLE_FF = "ff"

const ENABLE_22 = "22"
const DISABLE_23 = "23"

func boolToConst0F(b bool) string {
	if b {
		return ENABLE_0F
	}
	return DISABLE_F0
}
func boolToConst00(b bool) string {
	if b {
		return ENABLE_00
	}
	return DISABLE_FF
}
func boolToConst22(b bool) string {
	if b {
		return ENABLE_22
	}
	return DISABLE_23
}

var (
	REGISTER_GROUP_12 = byte(0x12)
	REGISTER_GROUP_13 = byte(0x13)
	REGISTER_GROUP_14 = byte(0x14)
)

type MessageType [2]byte

var (
	MSG_1202 = MessageType{0x12, 0x02}
	MSG_1203 = MessageType{0x12, 0x03}
	MSG_1204 = MessageType{0x12, 0x04}
	MSG_1205 = MessageType{0x12, 0x05}
	MSG_1206 = MessageType{0x12, 0x06}
	MSG_1207 = MessageType{0x12, 0x07}

	MSG_1302 = MessageType{0x13, 0x02}
	MSG_1303 = MessageType{0x13, 0x03}
	MSG_1304 = MessageType{0x13, 0x04}
	MSG_1305 = MessageType{0x13, 0x05}
	MSG_1306 = MessageType{0x13, 0x06}
	MSG_1307 = MessageType{0x13, 0x07}

	MSG_1402 = MessageType{0x14, 0x02}
	MSG_1403 = MessageType{0x14, 0x03}
	MSG_1404 = MessageType{0x14, 0x04}
	MSG_1405 = MessageType{0x14, 0x05}
	MSG_1406 = MessageType{0x14, 0x06}
	MSG_1407 = MessageType{0x14, 0x07}
)

// TODO store all that as byte
const (
	MODE_COOL = "12"
	MODE_DRY  = "22"
	MODE_WIND = "32"
	MODE_HEAT = "42"
	MODE_AUTO = "e2"
)

const (
	COMODE_OFF         = "12"
	COMODE_TURBO       = "22"
	COMODE_SMART       = "32"
	COMODE_SLEEP       = "42"
	COMODE_QUIET       = "52"
	COMODE_SOFT_COOL   = "62"
	COMODE_WIND_MODE_1 = "82"
	COMODE_WIND_MODE_2 = "92"
	COMODE_WIND_MODE_3 = "a2"
)

const (
	WIND_LEVEL_AUTO  = "00"
	WIND_LEVEL_LOW   = "12"
	WIND_LEVEL_MID   = "14"
	WIND_LEVEL_HIGH  = "16"
	WIND_LEVEL_TURBO = "18"
)

const (
	WIND_DIRECTION_OFF      = "12"
	WIND_DIRECTION_INDIRECT = "21"
	WIND_DIRECTION_DIRECT   = "31"
	WIND_DIRECTION_CENTER   = "41"
	WIND_DIRECTION_WIDE     = "51"
	WIND_DIRECTION_LEFT     = "61"
	WIND_DIRECTION_RIGHT    = "71"
	WIND_DIRECTION_LONG     = "81"
	WIND_DIRECTION_SWING_UD = "82"
	WIND_DIRECTION_SWING_LR = "a2"
	WIND_DIRECTION_ROTATION = "b2"
	WIND_DIRECTION_FIXED    = "c2"
)

const (
	WPS_DIRECT  = "20"
	WPS_DEFAULT = "0f"
)

const (
	FILTER_TIME_00  = "00"
	FILTER_TIME_180 = "01"
	FILTER_TIME_300 = "02"
	FILTER_TIME_500 = "03"
	FILTER_TIME_700 = "04"
)

const (
	REGISTER_12_AC_FUN_ENABLE     = 0x01
	REGISTER_12_AC_FUN_POWER      = 0x02
	REGISTER_12_AC_FUN_UNKNOWN_41 = 0x41
	REGISTER_12_AC_FUN_OPMODE     = 0x43
	REGISTER_12_AC_FUN_COMODE     = 0x44
	REGISTER_12_AC_FUN_WINDLEVEL  = 0x62
	REGISTER_12_AC_FUN_DIRECTION  = 0x63
	REGISTER_12_AC_FUN_UNKNOWN_EA = 0xea
	REGISTER_12_AC_FUN_TEMP_SET   = 0x5a
	REGISTER_12_AC_FUN_TEMP_NOW   = 0x5c
	REGISTER_12_AC_FUN_SLEEP      = 0x73
	REGISTER_12_AC_FUN_BIP        = 0x74
	REGISTER_12_AC_FUN_ERROR      = 0xf7
)

const (
	REGISTER_13_AC_ADD_AUTOCLEAN          = 0x32 // RW
	REGISTER_13_AC_ADD_SETKWH             = 0x40 // RW
	REGISTER_13_AC_ADD_CLEAR_FILTER_ALARM = 0x44 // RW
	REGISTER_13_AC_ADD_STARTWPS           = 0x43 // RW
	REGISTER_13_AC_ADD_SPI                = 0x75 // RW
	REGISTER_13_AC_OUTDOOR_TEMP           = 0x76
	REGISTER_13_AC_COOL_CAPABILITY        = 0x77
	REGISTER_13_AC_WARM_CAPABILITY        = 0x78
)

const (
	REGISTER_14_AC_ADD2_UNKNOWN_17    = 0x17
	REGISTER_14_AC_ADD2_UNKNOWN_18    = 0x18
	REGISTER_14_AC_ADD2_UNKNOWN_19    = 0x19
	REGISTER_14_AC_ADD2_USEDWATT      = 0x32 // R
	REGISTER_14_AC_SG_WIFI            = 0x37
	REGISTER_14_AC_SG_INTERNET        = 0x38
	REGISTER_14_AC_ADD2_OPTIONCODE    = 0x39
	REGISTER_14_AC_ADD2_OUT_VERSION   = 0xf3
	REGISTER_14_AC_ADD2_PANEL_VERSION = 0xf4
	REGISTER_14_AC_FUN_MODEL          = 0xf5
	REGISTER_14_AC_ADD2_VERSION       = 0xf6
	REGISTER_14_AC_SG_MACHIGH         = 0xf7
	REGISTER_14_AC_SG_MACMID          = 0xf8
	REGISTER_14_AC_SG_MACLOW          = 0xf9
	REGISTER_14_AC_ADD2_UNKNOWN_FD    = 0xfd
	REGISTER_14_AC_ADD2_USEDPOWER     = 0xe0 // R

	REGISTER_14_AC_ADD2_USEDTIME        = 0xe4 // R
	REGISTER_14_AC_ADD2_FILTER_USE_TIME = 0xe6 // R

	REGISTER_14_AC_ADD2_CLEAR_POWERTIME = 0xe8 // RW
	REGISTER_14_AC_ADD2_FILTERTIME      = 0xe9 // RW
	REGISTER_14_AC_SG_VENDER01          = 0xfa
	REGISTER_14_AC_SG_VENDER02          = 0xfb
	REGISTER_14_AC_SG_VENDER03          = 0xfc
)

// Reset REGISTER_14_AC_ADD2_USEDPOWER 0x1404 0xe80101 (REGISTER_14_AC_ADD2_CLEAR_POWERTIME)
// ap enable 1304 430120

var HEADER = [3]byte{0xd0, 0xc0, 0x02}
var FOOTER = byte(0xe0)
var SEP = byte(0xfe)

var EmptyMessage = message{
	Start:          HEADER,
	PayloadsLength: 0,
	Counter:        nil,
	Separator:      SEP,
	Payload:        NewEmptyPayload(),
	End:            FOOTER,
}

type registerGroups = map[byte]*registerGroup
type registerGroup = map[byte]registerDetails
type registerDetails struct {
	txt string
	//type bool/enum/string/int
}

var RegisterDetails = registerGroups{
	REGISTER_GROUP_12: {
		REGISTER_12_AC_FUN_ENABLE:     registerDetails{txt: "FUN_ENABLE"},
		REGISTER_12_AC_FUN_POWER:      registerDetails{txt: "FUN_POWER"},
		REGISTER_12_AC_FUN_UNKNOWN_41: registerDetails{txt: "FUN_UNKNOWN_41"},
		REGISTER_12_AC_FUN_OPMODE:     registerDetails{txt: "FUN_OPMODE"},
		REGISTER_12_AC_FUN_COMODE:     registerDetails{txt: "FUN_COMODE"},
		REGISTER_12_AC_FUN_WINDLEVEL:  registerDetails{txt: "FUN_WINDLEVEL"},
		REGISTER_12_AC_FUN_DIRECTION:  registerDetails{txt: "FUN_DIRECTION"},
		REGISTER_12_AC_FUN_UNKNOWN_EA: registerDetails{txt: "FUN_UNKNOWN_EA"},
		REGISTER_12_AC_FUN_TEMP_SET:   registerDetails{txt: "FUN_TEMP_SET"},
		REGISTER_12_AC_FUN_TEMP_NOW:   registerDetails{txt: "FUN_TEMP_NOW"},
		REGISTER_12_AC_FUN_SLEEP:      registerDetails{txt: "FUN_SLEEP"},
		REGISTER_12_AC_FUN_BIP:        registerDetails{txt: "FUN_BIP"},
		REGISTER_12_AC_FUN_ERROR:      registerDetails{txt: "FUN_ERROR"},
	},
	REGISTER_GROUP_13: {
		REGISTER_13_AC_ADD_AUTOCLEAN:          registerDetails{txt: "AC_ADD_AUTOCLEAN"},
		REGISTER_13_AC_ADD_SETKWH:             registerDetails{txt: "AC_ADD_SETKWH"},
		REGISTER_13_AC_ADD_CLEAR_FILTER_ALARM: registerDetails{txt: "AC_ADD_CLEAR_FILTER_ALARM"},
		REGISTER_13_AC_ADD_STARTWPS:           registerDetails{txt: "AC_ADD_STARTWPS"},
		REGISTER_13_AC_ADD_SPI:                registerDetails{txt: "AC_ADD_SPI"},
		REGISTER_13_AC_OUTDOOR_TEMP:           registerDetails{txt: "AC_OUTDOOR_TEMP"},
		REGISTER_13_AC_COOL_CAPABILITY:        registerDetails{txt: "AC_COOL_CAPABILITY"},
		REGISTER_13_AC_WARM_CAPABILITY:        registerDetails{txt: "AC_WARM_CAPABILITY"},
	},
	REGISTER_GROUP_14: {
		REGISTER_14_AC_ADD2_UNKNOWN_17:      registerDetails{txt: "AC_ADD2_UNKNOWN_17"},
		REGISTER_14_AC_ADD2_UNKNOWN_18:      registerDetails{txt: "AC_ADD2_UNKNOWN_18"},
		REGISTER_14_AC_ADD2_UNKNOWN_19:      registerDetails{txt: "AC_ADD2_UNKNOWN_19"},
		REGISTER_14_AC_ADD2_UNKNOWN_FD:      registerDetails{txt: "AC_ADD2_UNKNOWN_FD"},
		REGISTER_14_AC_ADD2_USEDWATT:        registerDetails{txt: "AC_ADD2_USEDWATT"},
		REGISTER_14_AC_SG_WIFI:              registerDetails{txt: "AC_SG_WIFI"},
		REGISTER_14_AC_SG_INTERNET:          registerDetails{txt: "AC_SG_INTERNET"},
		REGISTER_14_AC_ADD2_OPTIONCODE:      registerDetails{txt: "AC_ADD2_OPTIONCODE"},
		REGISTER_14_AC_ADD2_OUT_VERSION:     registerDetails{txt: "AC_ADD2_OUT_VERSION"},
		REGISTER_14_AC_ADD2_PANEL_VERSION:   registerDetails{txt: "AC_ADD2_PANEL_VERSION"},
		REGISTER_14_AC_FUN_MODEL:            registerDetails{txt: "AC_FUN_MODEL"},
		REGISTER_14_AC_ADD2_VERSION:         registerDetails{txt: "AC_ADD2_VERSION"},
		REGISTER_14_AC_SG_MACHIGH:           registerDetails{txt: "AC_SG_MACHIGH"},
		REGISTER_14_AC_SG_MACMID:            registerDetails{txt: "AC_SG_MACMID"},
		REGISTER_14_AC_SG_MACLOW:            registerDetails{txt: "AC_SG_MACLOW"},
		REGISTER_14_AC_ADD2_USEDPOWER:       registerDetails{txt: "AC_ADD2_USEDPOWER"},
		REGISTER_14_AC_ADD2_USEDTIME:        registerDetails{txt: "AC_ADD2_USEDTIME"},
		REGISTER_14_AC_ADD2_CLEAR_POWERTIME: registerDetails{txt: "AC_ADD2_CLEAR_POWERTIME"},
		REGISTER_14_AC_ADD2_FILTERTIME:      registerDetails{txt: "AC_ADD2_FILTERTIME"},
		REGISTER_14_AC_ADD2_FILTER_USE_TIME: registerDetails{txt: "AC_ADD2_FILTER_USE_TIME"},
		REGISTER_14_AC_SG_VENDER01:          registerDetails{txt: "AC_SG_VENDER01"},
		REGISTER_14_AC_SG_VENDER02:          registerDetails{txt: "AC_SG_VENDER02"},
		REGISTER_14_AC_SG_VENDER03:          registerDetails{txt: "AC_SG_VENDER03"},
	},
}

var AC_BIP = NewRegisterFromString(REGISTER_12_AC_FUN_BIP, "f0")

var AC_FUN_ENABLE = func(enable bool) Register {
	return NewRegisterFromString(REGISTER_12_AC_FUN_ENABLE, boolToConst0F(enable))
}
var AC_FUN_POWER = func(enable bool) Register {
	return NewRegisterFromString(REGISTER_12_AC_FUN_POWER, boolToConst0F(enable))
}
var AC_FUN_OPMODE = func(mode string) Register {
	return NewRegisterFromString(REGISTER_12_AC_FUN_OPMODE, mode)
}
var AC_FUN_COMODE = func(mode string) Register {
	return NewRegisterFromString(REGISTER_12_AC_FUN_COMODE, mode)
}
var AC_FUN_TEMPSET = func(temp int) Register {
	if temp < 16 || temp > 28 {
		panic("temp must be between 16 and 28")
	}
	hexTemp := strconv.FormatInt(int64(temp), 16)
	return NewRegisterFromString(REGISTER_12_AC_FUN_TEMP_SET, hexTemp)
}
var AC_FUN_TEMP_NOW = func(temp int) Register {
	if temp < 0 || temp > 50 {
		panic("temp must be between 0 and 50")
	}
	hexTemp := strconv.FormatInt(int64(temp), 16)
	return NewRegisterFromString(REGISTER_12_AC_FUN_TEMP_NOW, hexTemp)
}
var AC_FUN_WIND_LEVEL = func(level string) Register {
	return NewRegisterFromString(REGISTER_12_AC_FUN_WINDLEVEL, level)
}
var AC_FUN_WIND_DIRECTION = func(dir string) Register {
	return NewRegisterFromString(REGISTER_12_AC_FUN_DIRECTION, dir)
}
var AC_FUN_SLEEP = func(enable bool) Register {
	return NewRegisterFromString(REGISTER_12_AC_FUN_SLEEP, boolToConst00(enable))
}
var AC_FUN_ERROR = func(msg string) Register {
	h := fmt.Sprintf("%x", msg)
	if len(msg) == 0 {
		h = fmt.Sprintf("%02x%02x%02x%02x", 0x00, 0x00, 0x00, 0x00)
	}
	return NewRegisterFromString(REGISTER_12_AC_FUN_ERROR, h)
}

var AC_ADD_AUTOCLEAN = func(enable bool) Register {
	return NewRegisterFromString(REGISTER_13_AC_ADD_AUTOCLEAN, boolToConst22(enable))
}
var AC_ADD_SETKWH = func(kwh int) Register {
	hexKwh := fmt.Sprintf("%02x", kwh)
	return NewRegisterFromString(REGISTER_13_AC_ADD_SETKWH, hexKwh)
}
var AC_ADD_CLEAR_FILTER_ALARM = func(enable bool) Register {
	return NewRegisterFromString(REGISTER_13_AC_ADD_CLEAR_FILTER_ALARM, boolToConst00(enable))
}
var AC_ADD_SPI = func(enable bool) Register {
	return NewRegisterFromString(REGISTER_13_AC_ADD_SPI, boolToConst0F(enable))
}

//	var AC_ADD_APMODE_END = func(foo bool) Register {
//		return NewRegisterFromString(REGISTER_13_AC_ADD_APMODE_END, boolToConst0F(foo))
//	}
var AC_ADD_STARTWPS = func(mode string) Register {
	return NewRegisterFromString(REGISTER_13_AC_ADD_STARTWPS, mode)
}
var AC_OUTDOOR_TEMP = func(temp int) Register {
	hexTemp := fmt.Sprintf("%02x", temp)
	return NewRegisterFromString(REGISTER_13_AC_OUTDOOR_TEMP, hexTemp)
}

var AC_COOL_CAPABILITY = func(foo int) Register {
	hexFoo := fmt.Sprintf("%02x", foo)
	return NewRegisterFromString(REGISTER_13_AC_COOL_CAPABILITY, hexFoo)
}

var AC_WARM_CAPABILITY = func(foo int) Register {
	hexFoo := fmt.Sprintf("%02x", foo)
	return NewRegisterFromString(REGISTER_13_AC_WARM_CAPABILITY, hexFoo)
}

var AC_ADD2_USEDWATT = func(kwh int) Register {
	hexKwh := fmt.Sprintf("%02x", kwh)
	return NewRegisterFromString(REGISTER_14_AC_ADD2_USEDWATT, hexKwh)
}
var AC_ADD2_VERSION = func(version int) Register {
	hexV := fmt.Sprintf("%02x", version)
	return NewRegisterFromString(REGISTER_14_AC_ADD2_VERSION, hexV)
}
var AC_ADD2_PANEL_VERSION = func(version int) Register {
	hexV := fmt.Sprintf("%02x", version)
	return NewRegisterFromString(REGISTER_14_AC_ADD2_PANEL_VERSION, hexV)
}
var AC_ADD2_OUT_VERSION = func(version int) Register {
	hexV := fmt.Sprintf("%02x", version)
	return NewRegisterFromString(REGISTER_14_AC_ADD2_OUT_VERSION, hexV)
}
var AC_FUN_MODEL = func(version int) Register {
	hexV := fmt.Sprintf("%02x", version)
	return NewRegisterFromString(REGISTER_14_AC_FUN_MODEL, hexV)
}
var AC_ADD2_OPTIONCODE = func(opCode string) Register {
	return NewRegisterFromString(REGISTER_14_AC_ADD2_OPTIONCODE, opCode)
}
var AC_ADD2_USEDPOWER = func(kwh int) Register {
	hexKwh := strconv.FormatInt(int64(kwh), 16)
	return NewRegisterFromString(REGISTER_14_AC_ADD2_USEDPOWER, hexKwh)
}
var AC_ADD2_USEDTIME = func(hours int) Register {
	hexHours := fmt.Sprintf("%02x", hours)
	return NewRegisterFromString(REGISTER_14_AC_ADD2_USEDTIME, hexHours)
}
var AC_ADD2_CLEAR_POWERTIME = func(foo int) Register {
	hexFoo := fmt.Sprintf("%02x", foo)
	return NewRegisterFromString(REGISTER_14_AC_ADD2_CLEAR_POWERTIME, hexFoo)
}
var AC_ADD2_FILTERTIME = func(filterTime string) Register {
	return NewRegisterFromString(REGISTER_14_AC_ADD2_FILTERTIME, filterTime)
}
var AC_ADD2_FILTER_USE_TIME = func(filterUseTime string) Register {
	return NewRegisterFromString(REGISTER_14_AC_ADD2_FILTER_USE_TIME, filterUseTime)
}

var AC_SG_WIFI = func(enable bool) Register {
	return NewRegisterFromString(REGISTER_14_AC_SG_WIFI, boolToConst0F(enable))
}
var AC_SG_INTERNET = func(enable bool) Register {
	return NewRegisterFromString(REGISTER_14_AC_SG_INTERNET, boolToConst0F(enable))
}

var AC_UNKNOWN_17 = NewRegisterFromString(0x17, "15")
var AC_UNKNOWN_18 = NewRegisterFromString(0x18, "05")
var AC_UNKNOWN_19 = NewRegisterFromString(0x19, "27")
var AC_UNKNOWN_41 = NewRegisterFromString(0x41, "32")
var AC_UNKNOWN_EA = NewRegisterFromString(0xea, "fe")
var AC_UNKNOWN_FD = NewRegisterFromString(0xfd, "02")

var AC_SG_VENDER01 = NewRegisterFromString(REGISTER_14_AC_SG_VENDER01, "f8")
var AC_SG_VENDER02 = NewRegisterFromString(REGISTER_14_AC_SG_VENDER02, "04")
var AC_SG_VENDER03 = NewRegisterFromString(REGISTER_14_AC_SG_VENDER03, "2e")
var AC_SG_MACHIGH = NewRegisterFromString(REGISTER_14_AC_SG_MACHIGH, "d8")
var AC_SG_MACMID = NewRegisterFromString(REGISTER_14_AC_SG_MACMID, "3e")
var AC_SG_MACLOW = NewRegisterFromString(REGISTER_14_AC_SG_MACLOW, "a0")
