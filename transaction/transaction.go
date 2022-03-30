package transaction

import (
	"crypto/sha256"

	"github.com/roppa/bitcoin/golang/bitcoin/bytes"
	"github.com/roppa/bitcoin/golang/bitcoin/types"
)

type (
	// Input are the funds for the new transaction.
	Input struct {
		TxID          []byte        `json:"txid"` // reversed
		VOut          types.BuInt32 `json:"vout"` // reversed
		ScriptSigSize types.Varint  `json:"length"`
		ScriptSig     []byte        `json:"script"`   // unlocking script
		Sequence      types.BuInt32 `json:"sequence"` // reversed
	}

	// Output is where the payment/s are going.
	Output struct {
		Value            types.BuInt64 `json:"satoshis"` // reversed
		ScriptPubKeySize types.Varint  `json:"length"`
		ScriptPubKey     []byte        `json:"script"` // locking script
	}

	// Transaction is a bitcoin transaction.
	Transaction struct {
		Version     types.BuInt32 `json:"version"` // reversed
		Inputs      []Input       `json:"inputs"`
		InputCount  types.Varint  `json:"inputCount"`
		OutputCount types.Varint  `json:"outputCount"`
		Outputs     []Output      `json:"outputs"`
		Locktime    types.BuInt32 `json:"locktime"` // reversed
	}
)

// TxID returns the reversed double hash of the transaction byte data.
func (t *Transaction) TxID() []byte {
	h := sha256.Sum256([]byte(t.ToString()))
	hh := sha256.Sum256(h[:])
	return bytes.Reverse(hh[:])
}

func (t *Transaction) ToString() string {
	tx := ""
	tx += t.Version.ToString()
	tx += types.Varint(len(t.Inputs)).ToString()

	tx += string(t.Inputs[0].TxID)
	tx += t.Inputs[0].VOut.ToString()
	tx += t.Inputs[0].ScriptSigSize.ToString()
	tx += string(t.Inputs[0].ScriptSig)
	tx += t.Inputs[0].Sequence.ToString()

	tx += types.Varint(len(t.Outputs)).ToString()
	tx += t.Outputs[0].Value.ToString()
	tx += t.Outputs[0].ScriptPubKeySize.ToString()
	tx += string(t.Outputs[0].ScriptPubKey)

	tx += t.Locktime.ToString()
	return tx
}

// New creates a new Transaction.
func New() *Transaction {
	return &Transaction{
		Version: 1, // currently bitcoin is version 1
	}
}
