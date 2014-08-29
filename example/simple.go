package main

import "git.luma/lumos/merkle"

//Foo is an example struct
type Foo struct {
	merkle.HashItem
	A uint8
	B bool
	C string
}

//Hash implements
func (f *Foo) Hash() merkle.HashVal {
	h := merkle.NewHash()
	h.HashWrite(f.A)
	h.HashWrite(f.B)
	h.HashWrite(f.C)
	return h.SumHashVal(&f.cache)
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

	println("Hashes are different", merkle.Hash(foo1) != merkle.Hash(foo2))

	foo2.A = 42
	merkle.Update(foo2)

	println("Hashes are the same", merkle.Hash(foo1) == merkle.Hash(foo2))
}
