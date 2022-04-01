package keys

import (
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"

	"github.com/btcsuite/btcutil/base58"
)

type PrivateKey struct {
	Hex   string
	Bytes []byte
}

type WIFArgs struct {
	Network    string
	Compressed bool
}

var network = map[string]byte{
	"main": 0x80,
	"test": 0xEF,
}

func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func generateChecksum(bb []byte) []byte {
	h := sha256.Sum256(bb)
	hh := sha256.Sum256(h[:])
	return hh[0:4]
}

func generatePublicKey(r []byte) string {
	c := elliptic.P256()
	params := c.Params()
	px, py := c.ScalarMult(params.Gx, params.Gy, r)
	pubKey := string(px.Bytes()) + string(py.Bytes())
	return "04" + hex.EncodeToString([]byte(pubKey))
}

// GenerateWIF returns the WIF format of the private key.
func (p *PrivateKey) GenerateWIF(args WIFArgs) string {
	wif := []byte{network[args.Network]}
	wif = append(wif, p.Bytes...)
	if args.Compressed {
		wif = append(wif, 0x01)
	}
	wif = append(wif, generateChecksum(wif)...)
	return base58.Encode(wif)
}

// FromRandom creates a PrivateKey by generating 32 random bytes.
func FromRandom() *PrivateKey {
	bb := generateRandomBytes(32)
	return &PrivateKey{
		Hex:   hex.EncodeToString(bb),
		Bytes: bb,
	}
}

// FromBytes creates a private key from a bite slice.
func FromBytes(bb []byte) *PrivateKey {
	return &PrivateKey{
		Hex:   hex.EncodeToString(bb),
		Bytes: bb,
	}
}
