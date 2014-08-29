package merkle

import (
	"fmt"
	"reflect"
)

// HashVal ...
type HashVal uint64

// stringer!
func (val HashVal) String() string {
	return fmt.Sprintf("[HashVal:%x]", uint64(val))
}

// HashZero represents an uninitialised HashVal
const HashZero HashVal = 0

// Hasher provides a custom hash implementation for a type. Not
// everything needs to implement it, but doing so can speed
// updates.
type Hasher interface {
	Hash() HashVal
}

// CachingHasher is the interface for items that cache a hash value.
// Normally implemented by embedding HashCache.
type CachingHasher interface {
	CachedHash() *HashVal
}

// Item is a node in the Merkle tree, which must know how to find its
// parent Item (the root node should return nil) and should usually
// embed HashCache for efficient updates. To avoid using reflection,
// Items might benefit from being Hashers as well.
type Item interface {
	CachingHasher
	Parent() Item
}

// HashCache implements CachingHasher; it's meant to be embedded in your
// structs to make updating hash trees more efficient.
type HashCache struct {
	cache HashVal
}

// CachedHash implements CachingHasher.
func (h *HashCache) CachedHash() *HashVal {
	return &h.cache
}

// HashItem ...
type HashItem struct {
	HashCache
	parent Item
}

// Parent ...
func (i *HashItem) Parent() Item {
	return i.parent
}

// hash helpers
var trueByte = []byte{0xff}
var falseByte = []byte{0x00}

var hashCacheType = reflect.TypeOf(HashCache{})
var hashItemType = reflect.TypeOf(HashItem{})

// Hash returns something's hash, using a cached hash or Hash() method if
// available.
func Hash(i interface{}) HashVal {

	// fmt.Printf("hashing: %+v\n", i)
	if cachingHasher, ok := i.(CachingHasher); ok {
		if cached := *cachingHasher.CachedHash(); cached != HashZero {
			// fmt.Printf("retrieved cache from %+v\n", i)
			return cached
		}
	}

	//faster hashes
	switch i := i.(type) {
	case Hasher:
		return i.Hash()
	case uint8:
		return HashVal(uint64(i))
	case uint16:
		return HashVal(uint64(i))
	case uint32:
		return HashVal(uint64(i))
	case uint64:
		return HashVal(i)
	case string:
		return Hash([]byte(i))
	case bool:
		if i {
			return Hash(trueByte)
		}
		return Hash(falseByte)
	case []byte:
		h := NewMHash()
		h.Write(i)
		return h.SumHashVal(nil)
	}

	//===================
	//object is not a hasher,
	//we must use slow reflection!

	val := reflect.ValueOf(i)
	//deref if necessary
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	//we can reflect-hash *certain* types
	switch val.Kind() {
	//reflect on struct
	case reflect.Struct:
		var cacheptr *HashVal
		h := NewMHash()
		//hash all fields
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			ftype := field.Type()
			//ignore meta data and store cache address
			if ftype == hashCacheType {
				hashCache := field.Interface().(HashCache)
				cacheptr = &hashCache.cache
				continue
			} else if ftype == hashItemType {
				hashItem := field.Interface().(HashItem)
				cacheptr = &hashItem.cache
				continue
			}

			h.HashWrite(field.Interface())
		}
		return h.SumHashVal(cacheptr)
	//reflect on slices
	case reflect.Slice:
		h := NewMHash()
		for i := 0; i < val.Len(); i++ {
			h.HashWrite(val.Index(i).Interface())
		}
		return h.SumHashVal(nil)
	}

	panic("Unable to hash type: " + val.Kind().String())
}

// Update updates the chain of items between i and the root, given the
// leaf node that may have been changed.
func Update(i Item) {
	for i != nil {
		cached := i.CachedHash()
		*cached = HashZero // invalidate
		*cached = Hash(i)
		i = i.Parent()
	}
}
