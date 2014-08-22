package merkle

//Merkle ...
type Merkle struct {
	number uint32
	parent *Merkle
}

//New allocates a new ...
func New(root interface{}) (*Merkle, error) {
	return &Merkle{}, nil
}

//Get ...
func (m *Merkle) Get() (*Merkle, error) {
	return m.parent, nil
}

// h := sha1.New()
// io.WriteString(h, "His money is twice tainted:")
// io.WriteString(h, " 'taint yours and 'taint mine.")
// fmt.Printf("% x", h.Sum(nil))
