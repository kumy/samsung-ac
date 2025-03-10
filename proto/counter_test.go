package proto

import (
	"encoding/hex"
	"github.com/stretchr/testify/suite"
	"testing"
)

type counterSuite struct {
	suite.Suite
}

func TestCounterSuite(t *testing.T) {
	suite.Run(t, new(counterSuite))
}

func (s *counterSuite) TestNewCounterFromBytes() {
	testCases := []struct {
		name string
		msg  string
		want *counter
	}{
		{
			name: "Decode counter 15",
			msg:  "0f",
			want: &counter{Counter: 0x0f},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			h, _ := hex.DecodeString(tc.msg)

			got := NewCounterFromBytes(h[0])
			s.Assert().Equal(tc.want, got)
		})
	}
}
