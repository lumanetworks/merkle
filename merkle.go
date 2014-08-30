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
	Cache HashVal
}

// CachedHash implements CachingHasher.
func (h *HashCache) CachedHash() *HashVal {
	return &h.Cache
}

// HashItem implements Item; it implements the most common case:
// the Parent item stored as a struct field with a simple Parent()
// method to retrieve it
type HashItem struct {
	HashCache
	ParentItem Item
}

// Parent retrieves the item's parent
func (i *HashItem) Parent() Item {
	return i.ParentItem
}

// =================

// hash helpers
const trueVal, falseVal = HashVal(uint64(1)), HashVal(uint64(0))

// stored types
var hashCacheType = reflect.TypeOf(HashCache{})
var hashItemType = reflect.TypeOf(HashItem{})

// =================

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
			return trueVal
		}
		return falseVal
	case []byte:
		h := newCrc()
		h.Write(i)
		return HashVal(h.Sum64())
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
		h := NewHash()
		//hash all fields
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			ftype := field.Type()
			//ignore meta data and store cache address
			if ftype == hashCacheType {
				hashCache := field.Interface().(HashCache)
				cacheptr = &hashCache.Cache
				continue
			} else if ftype == hashItemType {
				hashItem := field.Interface().(HashItem)
				cacheptr = &hashItem.Cache
				continue
			}

			h.Write(field.Interface())
		}
		return h.SumAndCache(cacheptr)
	//reflect on slices
	case reflect.Slice:
		h := NewHash()
		for i := 0; i < val.Len(); i++ {
			h.Write(val.Index(i).Interface())
		}
		return h.Sum()
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
