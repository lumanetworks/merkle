package main

import "git.luma/lumos/merkle"

//Foo is an example struct
type Foo struct {
	merkle.HashItem
	A uint8
	B bool
	C string
}

//      === OPTIONAL ===
//Hash implements the merkle.Hasher interface
func (f *Foo) Hash() merkle.HashVal {
	h := merkle.NewHash()
	h.Write(f.A)
	h.Write(f.B)
	h.Write(f.C)
	return h.SumAndCache(&f.Cache)
}

func main() {

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

	println("foo1 == foo2:", merkle.Hash(foo1) == merkle.Hash(foo2))

	foo2.A = 42
	merkle.Update(foo2)

	println("foo1 == foo2:", merkle.Hash(foo1) == merkle.Hash(foo2))
}
