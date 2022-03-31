package transaction

import (
	"testing"

	"github.com/roppa/bitcoin/golang/bitcoin/types"
	"github.com/stretchr/testify/assert"
)

var tx string = `01000000017967a5185e907a25225574544c31f7b059c1a191d65b53dcc1554d339c4f9efc010000006a47304402206a2eb16b7b92051d0fa38c133e67684ed064effada1d7f925c842da401d4f22702201f196b10e6e4b4a9fff948e5c5d71ec5da53e90529c8dbd122bff2b1d21dc8a90121039b7bcd0824b9a9164f7ba098408e63e5b7e3cf90835cceb19868f54f8961a825ffffffff014baf2100000000001976a914db4d1141d0048b1ed15839d0b7a4c488cd368b0e88ac00000000`

func TestNew(t *testing.T) {
	tx := New()
	assert.Equal(t, tx, &Transaction{
		Version: 1,
	})
}

func TestTxID(t *testing.T) {
	tx := New()
	var sequence types.BuInt32 = 0xffffffff
	var value types.BuInt64 = 2207563
	tx.InputCount = types.Varint(1)
	tx.Inputs = append(tx.Inputs, Input{
		TxID:          []byte("7967a5185e907a25225574544c31f7b059c1a191d65b53dcc1554d339c4f9efc"),
		VOut:          1,
		ScriptSigSize: 106,
		ScriptSig:     []byte("47304402206a2eb16b7b92051d0fa38c133e67684ed064effada1d7f925c842da401d4f22702201f196b10e6e4b4a9fff948e5c5d71ec5da53e90529c8dbd122bff2b1d21dc8a90121039b7bcd0824b9a9164f7ba098408e63e5b7e3cf90835cceb19868f54f8961a825"),
		Sequence:      sequence,
	})

	tx.OutputCount = types.Varint(1)
	tx.Outputs = append(tx.Outputs, Output{
		Value:            value,
		ScriptPubKeySize: 25,
		ScriptPubKey:     []byte("76a914db4d1141d0048b1ed15839d0b7a4c488cd368b0e88ac"),
	})

	tx.Locktime = 0
	assert.Equal(t, "c1b4e695098210a31fe02abffe9005cffc051bbe86ff33e173155bcbdc5821e3", tx.TxID())
}

func TestToString(t *testing.T) {
	tx := New()
	var sequence types.BuInt32 = 4294967295
	var value types.BuInt64 = 627215281
	tx.InputCount = types.Varint(1)
	tx.Inputs = append(tx.Inputs, Input{
		TxID:          []byte("cbf4b80d954b5607be6ddfb5806c7126f77295836e0e44110d780ac59a683207"),
		VOut:          0,
		ScriptSigSize: 108,
		ScriptSig:     []byte("c3e21b0fb785f0844cbc6f93bd2a87d019ff6336be3366da33d76b246cdd17fa"),
		Sequence:      sequence,
	})

	tx.OutputCount = types.Varint(1)
	tx.Outputs = append(tx.Outputs, Output{
		Value:            value,
		ScriptPubKeySize: 25,
		ScriptPubKey:     []byte("76a914fc0319ac35bd66511bae0015e937d159c6bf7ccd88ac"),
	})

	tx.Locktime = 0

	assert.Equal(t, "0100000001cbf4b80d954b5607be6ddfb5806c7126f77295836e0e44110d780ac59a683207000000006cc3e21b0fb785f0844cbc6f93bd2a87d019ff6336be3366da33d76b246cdd17faffffffff01b18b6225000000001976a914fc0319ac35bd66511bae0015e937d159c6bf7ccd88ac00000000", tx.ToString())
}

func TestToString2(t *testing.T) {
	tx := New()
	var sequence types.BuInt32 = 0xffffffff
	var value types.BuInt64 = 2207563
	tx.InputCount = types.Varint(1)
	tx.Inputs = append(tx.Inputs, Input{
		TxID:          []byte("7967a5185e907a25225574544c31f7b059c1a191d65b53dcc1554d339c4f9efc"),
		VOut:          1,
		ScriptSigSize: 106,
		ScriptSig:     []byte("47304402206a2eb16b7b92051d0fa38c133e67684ed064effada1d7f925c842da401d4f22702201f196b10e6e4b4a9fff948e5c5d71ec5da53e90529c8dbd122bff2b1d21dc8a90121039b7bcd0824b9a9164f7ba098408e63e5b7e3cf90835cceb19868f54f8961a825"),
		Sequence:      sequence,
	})

	tx.OutputCount = types.Varint(1)
	tx.Outputs = append(tx.Outputs, Output{
		Value:            value,
		ScriptPubKeySize: 25,
		ScriptPubKey:     []byte("76a914db4d1141d0048b1ed15839d0b7a4c488cd368b0e88ac"),
	})

	tx.Locktime = 0

	assert.Equal(t, "01000000017967a5185e907a25225574544c31f7b059c1a191d65b53dcc1554d339c4f9efc010000006a47304402206a2eb16b7b92051d0fa38c133e67684ed064effada1d7f925c842da401d4f22702201f196b10e6e4b4a9fff948e5c5d71ec5da53e90529c8dbd122bff2b1d21dc8a90121039b7bcd0824b9a9164f7ba098408e63e5b7e3cf90835cceb19868f54f8961a825ffffffff014baf2100000000001976a914db4d1141d0048b1ed15839d0b7a4c488cd368b0e88ac00000000", tx.ToString())
}
