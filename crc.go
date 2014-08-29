package merkle

import (
	"encoding/binary"
	"hash"
	"hash/crc64"
)

//Hash ...
type Hash struct {
	hash.Hash64
}

//Pool of CRC tables
const poolSize = 5

var current = 0
var pool = make([]*crc64.Table, poolSize)

//NewHash creates a Hash64 object using a dynamically allocated pool of CRC Tables
func NewHash() Hash {
	var table *crc64.Table
	if pool[current] == nil {
		table = crc64.MakeTable(crc64.ECMA)
		pool[current] = table
	} else {
		table = pool[current]
	}
	current = (current + 1) % poolSize
	return Hash{Hash64: crc64.New(table)}
}

//HashWrite first hashes the input then writes the result
func (h *Hash) HashWrite(obj interface{}) error {
	return binary.Write(h, binary.BigEndian, uint64(Hash(obj)))
}

//SumHashVal is shorthand for sum64 + type conversion
func (h *Hash) SumHashVal(addr *HashVal) HashVal {
	if addr == nil {
		return HashVal(h.Sum64())
	}
	*addr = HashVal(h.Sum64())
	return *addr
}
