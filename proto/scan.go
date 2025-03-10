package proto

import (
	"bytes"
	"errors"
	"fmt"
)

func ScanMessage(data []byte, atEOF bool) (advance int, token []byte, err error) {
	headerIndex := bytes.Index(data, HEADER[:])
	//logger.Debugf("buffer: %02x", data)
	if headerIndex == -1 {
		return 0, nil, nil
	}

	if len(data) < headerIndex+4 {
		return 0, nil, nil
	}

	frameLen := int(data[headerIndex+3])
	frameEndIndex := headerIndex + 4 + frameLen - 1
	if len(data[headerIndex:]) < 4+frameLen {
		return 0, nil, nil
	}

	// TODO check checksum

	if data[frameEndIndex] != FOOTER {
		frame := data[headerIndex:frameEndIndex]
		return frameEndIndex, nil, errors.New(
			fmt.Sprintf("footer mismatch l:%d s:%d e:%d [%02x]: %02x %02x", frameLen, headerIndex, frameEndIndex, data[frameEndIndex], frame, data))
	}

	if atEOF && len(data) > 0 {
		return len(data), nil, nil
	}

	frame := data[headerIndex:frameEndIndex]
	return len(frame) + 1, frame, nil
}

func ScanRegister(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) < 2 {
		return 0, nil, nil
	}

	frameLen := int(data[1])
	frameEndIndex := 2 + frameLen
	if len(data) < 2+frameLen {
		return 0, nil, nil
	}

	if atEOF && len(data) > 0 {
		return len(data), nil, nil
	}

	frame := data[:frameEndIndex]
	return len(frame), frame, nil
}
