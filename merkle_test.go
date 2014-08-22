package merkle

import (
	"fmt"
	"testing"
)

type Foo struct {
	A int
	B bool
	C string
	D map[string]*Bazz
	E []*Bar
}

type Bazz struct {
	S int
	T int
	U int
}

type Bar struct {
	X int
	Y int
	Z int
}

func Test_Foo(t *testing.T) {

	foo := &Foo{}

	//create root of tree
	root, err := New(foo)
	if err != nil {
		panic(err)
	}

	//set values
	root.Set("A", 42)
	root.Set("B", false)
	root.Set("C", "ping")

	//bind events
	root.On("C", func() {
		fmt.Println("C changed")
	})

	root.(*MerkleStruct)

	//get underlying object
	fmt.Printf("=> %+v", root.Interface().(*Foo))
}
