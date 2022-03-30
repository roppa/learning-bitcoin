package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuInt32ToBytes(t *testing.T) {
	tests := map[string]struct {
		u   BuInt32
		exp []byte
	}{
		"1 should return little endian 1 in bytes": {
			u:   1,
			exp: []byte{1, 0, 0, 0},
		},
		"2 should return little endian 2 in bytes": {
			u:   2,
			exp: []byte{2, 0, 0, 0},
		},
		"256 should return little endian 2 in bytes": {
			u:   256,
			exp: []byte{0, 1, 0, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.exp, test.u.ToBytes())
		})
	}
}

func TestBuInt32ToString(t *testing.T) {
	tests := map[string]struct {
		u   BuInt32
		exp string
	}{
		"1 should return little endian hex string": {
			u:   1,
			exp: "01000000",
		},
		"2 should return little endian hex string": {
			u:   2,
			exp: "02000000",
		},
		"4294967296 should return little endian hex string": {
			u:   4294967295,
			exp: "ffffffff",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.exp, test.u.ToString())
		})
	}
}

func TestBuInt64ToBytes(t *testing.T) {
	tests := map[string]struct {
		u   BuInt64
		exp []byte
	}{
		"1 should return little endian 1 in bytes": {
			u:   1,
			exp: []byte{1, 0, 0, 0, 0, 0, 0, 0},
		},
		"2 should return little endian 2 in bytes": {
			u:   2,
			exp: []byte{2, 0, 0, 0, 0, 0, 0, 0},
		},
		"4294967296 should return little endian 2 in bytes": {
			u:   4294967296,
			exp: []byte{0, 0, 0, 0, 1, 0, 0, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.exp, test.u.ToBytes())
		})
	}
}

func TestBuInt64ToString(t *testing.T) {
	tests := map[string]struct {
		u   BuInt64
		exp string
	}{
		"1 should return little endian hex string": {
			u:   1,
			exp: "0100000000000000",
		},
		"2 should return little endian hex string": {
			u:   2,
			exp: "0200000000000000",
		},
		"4294967296 should return little endian hex string": {
			u:   4294967296,
			exp: "0000000001000000",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.exp, test.u.ToString())
		})
	}
}

func TestParseBuInt32(t *testing.T) {
	tests := map[string]struct {
		hexString string
		err       error
		exp       BuInt32
	}{
		"empty string should return error": {
			hexString: "",
			err:       errors.New("invalid string"),
		},
		"valid version 1 should return 1": {
			hexString: "01000000",
			exp:       1,
		},
		"valid version 2 should return 2": {
			hexString: "02000000",
			exp:       2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			expected, err := ParseBuInt32(test.hexString)
			if test.err != nil {
				assert.Error(t, err, test.err)
				return
			}
			assert.Equal(t, test.exp, expected)
		})
	}
}
