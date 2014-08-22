package merkle

import (
	"errors"
	"reflect"
)

//Merkle ...
type Merkle interface {
	Get(string) (Merkle, error)
	Set(interface{}) error
	Interface() (interface{}, error)
}

//New allocates a new Merkle tree using the provided object
func New(item interface{}) (Merkle, error) {
	b := newBase(nil)

	t := reflect.TypeOf(item)
	switch t.Kind() {
	case reflect.Struct:
		return newStruct(b)
	case reflect.Int:
		return newInt(b)
	}

	return nil, errors.New("Invalid object")
}

// how to hash
// h := sha1.New()
// io.WriteString(h, "His money is twice tainted:")
// io.WriteString(h, " 'taint yours and 'taint mine.")
// fmt.Printf("% x", h.Sum(nil))

//=============

type mBase struct {
	parent Merkle
	val    *interface{}
}

func newBase(parent Merkle) *mBase {
	return &mBase{parent: parent}
}

//Get a child node of this Merkle tree
func (m *mBase) Get(_ string) (Merkle, error) {
	return nil, errors.New("Not implemented")
}

//Set the value of this node
func (m *mBase) Set(_ interface{}) error {
	return errors.New("Not implemented")
}

//Return the underlying Interface
func (m *mBase) Interface() (interface{}, error) {
	return *m.val, nil
}

//============

type mStruct struct {
	*mBase
	fields []Merkle
}

func newStruct(b *mBase) (*mStruct, error) {
	return &mStruct{mBase: b}, nil
}

//Get ...
func (m *mStruct) Get(field string) (Merkle, error) {
	return m.parent, nil
}

//===========

type mInt struct {
	*mBase
	val    *int
	fields []Merkle
}

func newInt(b *mBase) (*mInt, error) {
	return &mInt{mBase: b}, nil
}

//Set ...
func (m *mInt) Set(obj interface{}) error {
	*m.val = obj.(int)
	return nil
}

func (m *mInt) Interface() (interface{}, error) {
	return *m.val, nil
}

//============
