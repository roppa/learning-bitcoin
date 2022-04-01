package transaction

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/roppa/bitcoin/golang/bitcoin/bytes"
	"github.com/roppa/bitcoin/golang/bitcoin/types"
)

type (
	// Transaction is a bitcoin transaction.
	Transaction struct {
		Version  types.BuInt32 `json:"version"` // reversed
		Inputs   []Input       `json:"inputs"`
		Outputs  []Output      `json:"outputs"`
		Locktime types.BuInt32 `json:"locktime"` // reversed
	}
)

// TxID returns the reversed double hash of the transaction byte data.
func (t *Transaction) TxID() string {
	raw, _ := hex.DecodeString(t.ToString())
	h := sha256.Sum256(raw)
	hh := sha256.Sum256(h[:])
	return hex.EncodeToString(bytes.Reverse(hh[:]))
}

// ToString returns the hex encoded data that makes up the transaction.
func (t *Transaction) ToString() string {
	tx := ""
	tx += t.Version.ToString()
	tx += types.Varint(len(t.Inputs)).ToString()

	for i := 0; i < len(t.Inputs); i++ {
		tx += t.Inputs[i].ToString()
	}

	tx += types.Varint(len(t.Outputs)).ToString()
	for j := 0; j < len(t.Outputs); j++ {
		tx += t.Outputs[j].ToString()
	}

	tx += t.Locktime.ToString()
	return tx
}

// MarshalJSON is the custom json marshal for a transaction.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	tx := t.ToString()
	txid := t.TxID()
	return json.Marshal(&struct {
		Version     uint32   `json:"version"`
		Locktime    uint32   `json:"locktime"`
		TxID        string   `json:"txid"`
		Size        int      `json:"size"`
		InputCount  int      `json:"vincount"`
		Inputs      []Input  `json:"vin"`
		OutputCount int      `json:"voutcount"`
		Outputs     []Output `json:"vout"`
	}{
		Version:     uint32(t.Version),
		Locktime:    uint32(t.Locktime),
		TxID:        txid,
		Size:        len(tx) / 2, // size of hex bytes
		InputCount:  len(t.Inputs),
		Outputs:     t.Outputs,
		OutputCount: len(t.Outputs),
		Inputs:      t.Inputs,
	})
}

// New creates a new Transaction.
func New() *Transaction {
	return &Transaction{
		Version: 1, // currently bitcoin is version 1
	}
}
