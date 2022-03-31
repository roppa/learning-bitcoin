package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputToString(t *testing.T) {
	i := Input{
		TxID:          []byte("7967a5185e907a25225574544c31f7b059c1a191d65b53dcc1554d339c4f9efc"),
		VOut:          1,
		ScriptSigSize: 106,
		ScriptSig:     []byte("47304402206a2eb16b7b92051d0fa38c133e67684ed064effada1d7f925c842da401d4f22702201f196b10e6e4b4a9fff948e5c5d71ec5da53e90529c8dbd122bff2b1d21dc8a90121039b7bcd0824b9a9164f7ba098408e63e5b7e3cf90835cceb19868f54f8961a825"),
		Sequence:      0xffffffff,
	}

	assert.Equal(t, "7967a5185e907a25225574544c31f7b059c1a191d65b53dcc1554d339c4f9efc010000006a47304402206a2eb16b7b92051d0fa38c133e67684ed064effada1d7f925c842da401d4f22702201f196b10e6e4b4a9fff948e5c5d71ec5da53e90529c8dbd122bff2b1d21dc8a90121039b7bcd0824b9a9164f7ba098408e63e5b7e3cf90835cceb19868f54f8961a825ffffffff", i.ToString())
}
