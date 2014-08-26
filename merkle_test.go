package merkle

import (
	"bytes"
	"testing"
)

type Foo struct {
	*Merkle
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

	if err := /*merkle.*/ Init(foo1); err != nil {
		t.Errorf("Failed to initialise 'foo1': %s", err)
		return
	}

	foo2 := &Foo{
		A: 35,
		B: true,
		C: "foo",
	}

	if err := /*merkle.*/ Init(foo2); err != nil {
		t.Errorf("Failed to initialise 'foo2': %s", err)
		return
	}

	if bytes.Compare(foo1.Hash, foo2.Hash) == 0 {
		t.Error("Hashes should be different")
		return
	}

	foo2.A = 42

	if err := foo2.Update(); err != nil {
		t.Errorf("Failed to update 'foo2': %s", err)
		return
	}

	if bytes.Compare(foo1.Hash, foo2.Hash) != 0 {
		t.Error("Hashes should be equal")
	}
}
