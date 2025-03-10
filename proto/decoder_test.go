package proto

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type decodeMessageSuite struct {
	suite.Suite
}

func TestDecodeMessageSuite(t *testing.T) {
	suite.Run(t, new(decodeMessageSuite))
}

func (s *decodeMessageSuite) TestDecodeMessage() {
	testCases := []struct {
		name string
		msg  string
		//want string
	}{
		{
			name: "Empty",
			msg:  "d0c002 0c 000000000000 fe 1404 00 f0 e0",
		},
		{
			name: "1404",
			msg:  "d0c002 0f 000000000001 fe 1404 03 37010f c8 e0",
		},
		{
			name: "1206",
			msg: "d0c002 30 000000000006 fe 1206 24 " +
				"0201f0 " +
				"410132 " +
				"4301e2 " +
				"440112 " +
				"620100 " +
				"6301c2 " +
				"ea01fe " +
				"5a0115 " +
				"5c0117 " +
				"730100 " +
				"f70400000000 " +
				"cf e0",
		},
		{
			name: "1402",
			msg:  "d0c002 22 000000000009 fe 1402 16 3200 f600 f400 f300 f500 3900 e000 e400 e800 e900 e600 2b e0",
		},
		{
			name: "1403",
			msg: "d0c002 39 000000000009 fe 1403 2d " +
				"3201fe " +
				"f601fe " +
				"f403150224 " +
				"f303130709 " +
				"f50105 " +
				"3902d0b8 " +
				"e0040000e635 " +
				"e40400040dcb " +
				"e801fe " +
				"e90103 " +
				"e6022710 " +
				"93 e0",
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			d := NewDecoder()
			got := d.DecodeFromString(tc.msg)
			fmt.Println(got.StringRaw())
			s.Assert().Equal(tc.msg, got.StringRaw())
		})
	}
}
