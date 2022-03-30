package types

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

type (
	// BuInt32 is a bitcoin little endian wrapper around uint32
	BuInt32 uint32

	// BuInt64 is a bitcoin little endian wrapper around uint64
	BuInt64 uint64
)

func (b BuInt64) ToBytes() []byte {
	vb := make([]byte, 8)
	binary.LittleEndian.PutUint64(vb, uint64(b))
	return vb
}

func (b BuInt64) ToString() string {
	return fmt.Sprintf("%x", b.ToBytes())
}

func (b BuInt32) ToBytes() []byte {
	vb := make([]byte, 4)
	binary.LittleEndian.PutUint32(vb, uint32(b))
	return vb
}

func (b BuInt32) ToString() string {
	return fmt.Sprintf("%x", b.ToBytes())
}

func ParseBuInt32(str string) (BuInt32, error) {
	if len(str) != 8 {
		return 0, errors.New("invalid string")
	}
	decoded, err := hex.DecodeString(str)
	if err != nil {
		return 0, err
	}

	return BuInt32(binary.LittleEndian.Uint32(decoded)), nil
}
