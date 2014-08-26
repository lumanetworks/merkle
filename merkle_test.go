package merkle

import (
	"bytes"
	"testing"
)

type Foo struct {
	Merkle
	A int
	B bool
	C string
}

func Test_Foo(t *testing.T) {

	foo1 := &Foo{
		A: 42,
		B: true,
		C: "foo",
	}

	/*merkle.*/ Init(foo1)

	foo2 := &Foo{
		A: 35,
		B: true,
		C: "foo",
	}

	/*merkle.*/ Init(foo2)

	if bytes.Compare(foo1.Hash, foo2.Hash) == 0 {
		t.Error("Hashes should be different")
	}

	foo2.A = 42
	foo2.Update()

	if bytes.Compare(foo1.Hash, foo2.Hash) != 0 {
		t.Error("Hashes should be equal")
	}
}
