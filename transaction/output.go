package transaction

import (
	"encoding/json"

	"github.com/roppa/bitcoin/golang/bitcoin/types"
)

// Output is where the payment/s are going.
type Output struct {
	Value        types.BuInt64 `json:"satoshis"` // reversed
	ScriptPubKey []byte        `json:"script"`   // locking script
}

// ToString returns the hex of the output.
func (o *Output) ToString() string {
	output := ""                                               // an output is composed of:
	output += o.Value.ToString()                               // how many satoshis
	output += types.Varint(len(o.ScriptPubKey) / 2).ToString() // the upcoming size of the locking code
	output += string(o.ScriptPubKey)                           // the script that locks the output
	return output
}

// MarshalJSON is the custom json marshal for an output.
func (o *Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Value            types.BuInt64 `json:"satoshis"` // reversed
		ScriptPubKeySize types.Varint  `json:"length"`
		ScriptPubKey     []byte        `json:"script"` // locking script
	}{
		Value:            o.Value,
		ScriptPubKeySize: types.Varint(len(o.ScriptPubKey) / 2),
		ScriptPubKey:     o.ScriptPubKey,
	})
}
