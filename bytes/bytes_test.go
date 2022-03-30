package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := map[string]struct {
		bytes []byte
		exp   []byte
	}{
		"empty": {
			bytes: []byte{},
			exp:   []byte{},
		},
		"two bytes": {
			bytes: []byte{1, 2},
			exp:   []byte{2, 1},
		},
		"three bytes": {
			bytes: []byte{1, 2, 3},
			exp:   []byte{3, 2, 1},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.exp, Reverse(test.bytes))
		})
	}
}
