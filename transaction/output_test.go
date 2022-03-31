package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputToString(t *testing.T) {
	o := Output{
		Value:            2207563,
		ScriptPubKeySize: 25,
		ScriptPubKey:     []byte("76a914db4d1141d0048b1ed15839d0b7a4c488cd368b0e88ac"),
	}
	assert.Equal(t, "4baf2100000000001976a914db4d1141d0048b1ed15839d0b7a4c488cd368b0e88ac", o.ToString())
}
