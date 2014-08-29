# Merkle

A merkle-tree implementation in Go

## Install

```
go get ...
```

## Usage

Embed the `merkle.HashItem` struct in your structs:

``` go
type Foo struct {
	merkle.HashItem
	A uint8
	B bool
	C string
}

func (f *Foo) Hash() HashVal {
	h := merkle.NewMHash()
	h.HashWrite(f.A)
	h.HashWrite(f.B)
	h.HashWrite(f.C)
	return h.SumHashVal(&f.cache)
}
```
 
Now you can `merkle.Hash` your `Foo`s and summarise their contents:

``` go
var foo1 = &Foo{
	A: 42,
	B: true,
	C: "foo",
}

var foo2 = &Foo{
	A: 35,
	B: true,
	C: "foo",
}

println("Hashes are different", merkle.Hash(foo1) != merkle.Hash(foo2))

foo2.A = 42
merkle.Update(foo2)

println("Hashes are the same", merkle.Hash(foo1) == merkle.Hash(foo2))
```

