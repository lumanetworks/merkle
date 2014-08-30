package merkle

import "testing"

//==============================

type Foo struct {
	HashItem
	A uint8
	B bool
	C string
}

func (f *Foo) Hash() HashVal {
	h := NewHash()
	h.Write(f.A)
	h.Write(f.B)
	h.Write(f.C)
	return h.SumAndCache(&f.Cache)
}

func Test_Simple(t *testing.T) {

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

	if h1, h2 := Hash(foo1), Hash(foo2); h1 == h2 {
		t.Errorf("Hashes should be different (%v == %v)", h1, h2)
		return
	}

	//make equal then rehash
	foo2.A = 42
	Update(foo2)

	if h1, h2 := Hash(foo1), Hash(foo2); h1 != h2 {
		t.Errorf("Hashes should be equal (%v != %v)", h1, h2)
		return
	}
}

//==============================

type Ping struct {
	HashItem
	D uint8
	E bool
	F []*Pong
}

// func (p *Ping) Hash() HashVal {
// 	h := NewHash()
// 	h.Write(p.D)
// 	h.Write(p.E)
// 	for _, f := range p.F {
// 		h.Write(f)
// 	}
// 	return h.SumAndCache(&p.Cache)
// }

type Pong struct {
	HashItem
	X uint8
	Y uint8
}

// func (p *Pong) Hash() HashVal {
// 	h := NewHash()
// 	h.Write(p.X)
// 	h.Write(p.Y)
// 	v := h.SumAndCache(&p.Cache)
// 	return v
// }

func Test_Nested(t *testing.T) {

	//create ping1
	var ping1 = &Ping{
		D: 89,
		E: false,
	}

	ping1.F = []*Pong{
		&Pong{X: 2, Y: 3, HashItem: HashItem{ParentItem: ping1}},
		&Pong{X: 4, Y: 5, HashItem: HashItem{ParentItem: ping1}},
	}

	//create ping2
	var ping2 = &Ping{
		D: 90,
		E: true,
	}

	ping2.F = []*Pong{
		&Pong{X: 2, Y: 3, HashItem: HashItem{ParentItem: ping2}},
		&Pong{X: 5, Y: 5, HashItem: HashItem{ParentItem: ping2}},
	}

	//hash em
	if h1, h2 := Hash(ping1), Hash(ping2); h1 == h2 {
		t.Errorf("Hashes should be different (%v == %v)", h1, h2)
		return
	}

	ping2.D = 89
	ping2.E = false
	ping2.F[1].X = 4
	Update(ping2.F[1])

	if h1, h2 := Hash(ping1), Hash(ping2); h1 != h2 {
		t.Errorf("Hashes should be equal (%v != %v)", h1, h2)
		return
	}
}

//==============================
