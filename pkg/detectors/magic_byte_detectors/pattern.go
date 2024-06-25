package magic_byte_detectors

import (
	"encoding/hex"
	"strings"
)

type Pattern interface {
	GetMagicBytes() []byte
	GetDetectedFileType() string
}

type pattern struct {
	magicBytesInHex string
	fileType        string
}

func (receiver pattern) GetMagicBytes() []byte {
	data, err := hex.DecodeString(strings.ToUpper(receiver.magicBytesInHex))
	if err != nil {
		panic(err)
	}
	return data
}

func (receiver pattern) GetDetectedFileType() string {
	return receiver.fileType
}
