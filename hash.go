package verkle

import (
	"github.com/zeebo/blake3"
)

var hash *blake3.Hasher
var _hash *blake3.Hasher

func init() {
	hash = blake3.New()
	_hash = blake3.New()
}

func hash256(data []byte) [32]byte {
	var out [32]byte
	_hash.Reset()
	_hash.Write(data[:])
	_hash.Digest().Read(out[:])
	return out
}
