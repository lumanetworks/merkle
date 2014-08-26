package merkle

import (
	"errors"
	"fmt"
	"reflect"
)

//Init any struct with an embedded merkle object
func Init(obj interface{}) error {

	vptr := reflect.ValueOf(obj)
	v := vptr.Elem()

	if v.Kind() != reflect.Struct {
		return errors.New("Must be struct [pointer]")
	}

	mfield := v.FieldByName("Merkle")

	if !mfield.IsValid() {
		return errors.New("Must have a 'Merkle' property")
	}

	cont := mfield.FieldByName("Container")
	cont.Set(vptr)

	initd := mfield.FieldByName("Initd")
	initd.SetBool(true)

	return nil
}

//Merkle state
type Merkle struct {
	Initd     bool
	Container interface{}
	Fields    []reflect.StructField
	Hash      []byte
}

//Update this Merkle tree
func (m *Merkle) Update() error {

	if !m.Initd {
		return errors.New("Must merkle.Init() this object first")
	}

	t := reflect.TypeOf(m.Container)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	name := t.Name()

	numFields := t.NumField()

	m.Fields = make([]reflect.StructField, numFields)

	for i := 0; i < numFields; i++ {
		f := t.Field(i)
		//TODO
		fmt.Printf("%s: %s\n", name, f.Name)
		m.Fields[i] = f
	}

	return nil
}

//Set the value of this node
// func (m *Merkle) Set(_ interface{}) error {
// 	return errors.New("Not implemented")
// }

//Return the underlying Interface
// func (m *Merkle) Interface() (interface{}, error) {
// 	return *m.val, nil
// }

// Hashing example
// h := sha1.New()
// io.WriteString(h, "His money is twice tainted:")
// io.WriteString(h, " 'taint yours and 'taint mine.")
// fmt.Printf("% x", h.Sum(nil))
