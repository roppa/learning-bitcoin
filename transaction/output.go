package transaction

import "github.com/roppa/bitcoin/golang/bitcoin/types"

// Output is where the payment/s are going.
type Output struct {
	Value            types.BuInt64 `json:"satoshis"` // reversed
	ScriptPubKeySize types.Varint  `json:"length"`
	ScriptPubKey     []byte        `json:"script"` // locking script
}

// ToString returns the hex of the output.
func (o *Output) ToString() string {
	output := ""
	output += o.Value.ToString()            // how many satoshis
	output += o.ScriptPubKeySize.ToString() // the upcoming size of the locking code
	output += string(o.ScriptPubKey)        // the script that locks the output
	return output
}
