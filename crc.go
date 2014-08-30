package merkle

import (
	"encoding/binary"
	"hash"
	"hash/crc64"
)

//Pool of CRC tables
const poolSize = 5

var current = 0
var pool = make([]*crc64.Table, poolSize)

func newCrc() hash.Hash64 {
	var table *crc64.Table
	if pool[current] == nil {
		table = crc64.MakeTable(crc64.ECMA)
		pool[current] = table
	} else {
		table = pool[current]
	}
	current = (current + 1) % poolSize
	return crc64.New(table)
}

//Hash64 ...
type Hash64 struct {
	crc hash.Hash64
}

//NewHash creates a Hash64 object using a dynamically allocated pool of CRC Tables
func NewHash() Hash64 {
	return Hash64{crc: newCrc()}
}

//Write first hashes the input then writes the result
func (h *Hash64) Write(obj interface{}) error {
	return binary.Write(h.crc, binary.BigEndian, uint64(Hash(obj)))
}

//Sum collates the result of this hash
func (h *Hash64) Sum() HashVal {
	return HashVal(h.crc.Sum64())
}

//SumAndCache both sums and caches this result at the address provided
func (h *Hash64) SumAndCache(addr *HashVal) HashVal {
	*addr = h.Sum()
	return *addr
}
