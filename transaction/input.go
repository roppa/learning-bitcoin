package transaction

import "github.com/roppa/bitcoin/golang/bitcoin/types"

// Input are the funds for the new transaction.
type Input struct {
	TxID          []byte        `json:"txid"` // reversed
	VOut          types.BuInt32 `json:"vout"` // reversed
	ScriptSigSize types.Varint  `json:"length"`
	ScriptSig     []byte        `json:"script"`   // unlocking script
	Sequence      types.BuInt32 `json:"sequence"` // reversed
}

// ToString returns the hex string of the input.
func (i *Input) ToString() string {
	input := ""
	input += string(i.TxID)             // input txid
	input += i.VOut.ToString()          // vector out
	input += i.ScriptSigSize.ToString() // size of the next unlocking script
	input += string(i.ScriptSig)        // script that unlocks the tx
	input += i.Sequence.ToString()      //
	return input
}
