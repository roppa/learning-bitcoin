package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	tests := map[string]struct {
		i   uint64
		exp string
	}{
		"0 should return `0`": {
			i:   0,
			exp: "00",
		},
		"106 should return `6a`": {
			i:   106,
			exp: "6a",
		},
		"252 should return `fc`": {
			i:   252,
			exp: "fc",
		},
		"550 should return `fd0226`": {
			i:   550,
			exp: "fd0226",
		},
		"13337 should return `fd3419`": {
			i:   13337,
			exp: "fd3419",
		},
		"998000 should return `fe000f3a70`": {
			i:   998000,
			exp: "fe000f3a70",
		},
		"134250981 should return `fe080081e5`": {
			i:   134250981,
			exp: "fe080081e5",
		},
		"18446744073709551615 should return `ffffffffffffffffff`": {
			i:   18446744073709551615,
			exp: "ffffffffffffffffff",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.exp, Varint(test.i).ToString())
		})
	}
}

func TestVarintLength(t *testing.T) {
	tests := map[string]struct {
		b   byte
		exp int
	}{
		"0 should return 0": {
			b:   0,
			exp: 0,
		},
		"1 should return 0": {
			b:   1,
			exp: 0,
		},
		"252 should return 0": {
			b:   252,
			exp: 0,
		},
		"253 should return 2": {
			b:   253,
			exp: 2,
		},
		"hex 0xfd should return 2": {
			b:   0xfd,
			exp: 2,
		},
		"hex byte 0xfe should return 4": {
			b:   0xfe,
			exp: 4,
		},
		"hex byte 0xff should return 8": {
			b:   0xff,
			exp: 8,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.exp, VarintLength(test.b))
		})
	}
}

func TestVarintValue(t *testing.T) {
	tests := map[string]struct {
		bb  []byte
		exp Varint
	}{
		"0 should return 0": {
			bb:  []byte{0},
			exp: 0,
		},
		"1 should return 1": {
			bb:  []byte{1},
			exp: 1,
		},
		"fd0226 in hex should return 550 little endian": {
			bb:  []byte{0xfd, 0x26, 0x02},
			exp: 550,
		},
		"fe000f3a70 in hex should return 998000 little endian": {
			bb:  []byte{0xfe, 0x70, 0x3a, 0x0f, 0x00},
			exp: 998000,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.exp, VarintValue(test.bb))
		})
	}
}
