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

	mptr := v.FieldByName("Merkle")

	if mptr.Kind() != reflect.Ptr {
		return errors.New("Must have an embedded 'Merkle' pointer")
	}

	m := &Merkle{}

	mptr.Set(reflect.ValueOf(m))

	m.Container = obj

	//Extract struct fields
	t := v.Type()
	numFields := t.NumField() - 1 //Merkle field is not included

	m.Fields = make([]reflect.StructField, numFields)

	for i := 1; i <= numFields; i++ {
		f := t.Field(i)
		fmt.Printf("Init %s\n", f.Name)
		//TODO: Recursive initialise field
		m.Fields[i-1] = f
	}

	m.Initd = true
	fmt.Printf("Init %+v\n", obj)
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
	if m == nil {
		return errors.New("Must merkle.Init() this object first")
	}

	for _, f := range m.Fields {
		fmt.Printf("Update %s\n", f.Name)
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
