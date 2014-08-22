package merkle

import (
	"fmt"
	"testing"
)

type Root struct {
	A int
	B bool
	C string
	D map[string]*Foo
	E []*Bar
	M *Merkle
}

type Foo struct {
	S int
	T int
	U int
	M *Merkle
}

type Bar struct {
	X int
	Y int
	Z int
	M *Merkle
}

func Test_Foo(t *testing.T) {

	r := &Root{}
	m, _ := New(r)

	// m.Get("D").Get("foo1").Get()

	r.D["something"].S

	subm, _ := m.Get()

	// m.Put("")

	fmt.Printf("result = %s", result)
}
