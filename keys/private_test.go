package keys

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromRandom(t *testing.T) {
	p := FromRandom()
	assert.Equal(t, 64, len(p.Hex))
}

func TestGenerateWIF(t *testing.T) {
	tests := map[string]struct {
		pkHex string
		args  WIFArgs
		exp   string
	}{
		"mainnet, uncompressed": {
			pkHex: "ef235aacf90d9f4aadd8c92e4b2562e1d9eb97f0df9ba3b508258739cb013db2",
			args: WIFArgs{
				Network: "main",
			},
			exp: "5Kdc3UAwGmHHuj6fQD1LDmKR6J3SwYyFWyHgxKAZ2cKRzVCRETY",
		},
		"mainnet, compressed": {
			pkHex: "ef235aacf90d9f4aadd8c92e4b2562e1d9eb97f0df9ba3b508258739cb013db2",
			args: WIFArgs{
				Network:    "main",
				Compressed: true,
			},
			exp: "L5EZftvrYaSudiozVRzTqLcHLNDoVn7H5HSfM9BAN6tMJX8oTWz6",
		},
		"testnet, uncompressed": {
			pkHex: "ef235aacf90d9f4aadd8c92e4b2562e1d9eb97f0df9ba3b508258739cb013db2",
			args: WIFArgs{
				Network: "test",
			},
			exp: "93QEdCzUrzMRsnbx2YuF6MsNjxQA6iWSrv9e2wX4NM4UmYzUsLn",
		},
		"testnet, compressed": {
			pkHex: "ef235aacf90d9f4aadd8c92e4b2562e1d9eb97f0df9ba3b508258739cb013db2",
			args: WIFArgs{
				Network:    "test",
				Compressed: true,
			},
			exp: "cVbZ8ovhye9AoAHFsqobCf7LxbXDAECy9Kb8TZdfsDYMZGBUyCnm",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			bb, _ := hex.DecodeString(test.pkHex)
			p := FromBytes(bb)
			assert.Equal(t, test.exp, p.GenerateWIF(test.args))
		})
	}
}

func TestGenerateChecksum(t *testing.T) {
	bb := []byte{0x52, 0xfd, 0xfc, 0x7, 0x21, 0x82, 0x65, 0x4f, 0x16, 0x3f, 0x5f, 0xf, 0x9a, 0x62, 0x1d, 0x72, 0x95, 0x66, 0xc7, 0x4d, 0x10, 0x3, 0x7c, 0x4d, 0x7b, 0xbb, 0x4, 0x7, 0xd1, 0xe2, 0xc6, 0x49}
	exp, _ := hex.DecodeString("042da4d4")
	assert.Equal(t, exp, generateChecksum(bb))
}

func TestGeneratePublicKey(t *testing.T) {
	pkHex := "dcf1194267982c38dd3d152dad0f49c1511e00f6aa519b52d48c1dece10ed8ed"
	bb, _ := hex.DecodeString(pkHex)
	p := FromBytes(bb)
	assert.Equal(t, "047b738638440332c6859423348955b455dbd00c228a27e554e38da06aa9acf45f06f334f357886628058e2bd7e42bb95d7830e8db961eeec6c99fa54bc07b3fbe", string(generatePublicKey(p.Bytes)))
}
