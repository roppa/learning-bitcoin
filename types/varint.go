package types

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

type Varint uint64

// ToString returns the hex value of the varint.
func (v Varint) ToString() string {
	h := strconv.FormatUint(uint64(v), 16)
	if v < 253 {
		return fmt.Sprintf("%02s", h)
	} else if v < 65536 {
		return fmt.Sprintf("fd%04s", h)
	} else if v < 4294967296 {
		return fmt.Sprintf("fe%08s", h)
	}
	return fmt.Sprintf("ff%016s", h)
}

// VarintLength given a number of bytes, it tells you how big the varint is.
func VarintLength(b byte) int {
	if b == 0xff {
		return 8
	} else if b == 0xfe {
		return 4
	} else if b == 0xfd {
		return 2
	}
	return 0
}

// VarintValue returns the actual number, or size, of inputs/outputs.
func VarintValue(bb []byte) Varint {
	r := make([]byte, 8)
	l := VarintLength(bb[0])
	if l == 0 {
		r[0] = bb[0]
	}
	for i := 0; i < l; i++ {
		r[i] = bb[i+1]
	}
	return Varint(binary.LittleEndian.Uint64(r))
}
