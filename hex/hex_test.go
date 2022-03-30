package hex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexToUint64(t *testing.T) {
	tests := map[string]struct {
		hex string
		exp uint64
		err error
	}{
		"empty string should return 0": {
			hex: "",
			exp: 0,
		},
		"01 in hex should return 1": {
			hex: "01",
			exp: 1,
		},
		"02 in hex should return 2": {
			hex: "02",
			exp: 2,
		},
		"fc in hex should return 252": {
			hex: "fc",
			exp: 252,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			i, err := HexToUint64(test.hex)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
				return
			}
			assert.Equal(t, test.exp, i)
		})
	}
}
