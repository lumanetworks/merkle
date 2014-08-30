# Merkle

#### :warning: Incomplete

A merkle-tree implementation in Go

## Install

```
go get ...
```

## Usage

Embed the `merkle.HashItem` struct in your structs:

``` go
package main

import "git.luma/lumos/merkle"

//Foo is an example struct
type Foo struct {
	merkle.HashItem
	A uint8
	B bool
	C string
}

```
 
Now you can `merkle.Hash()` your `Foo`s and summarise their contents:

``` go
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

	println("foo1 == foo2:", merkle.Hash(foo1) == merkle.Hash(foo2))

	foo2.A = 42
	merkle.Update(foo2)

	println("foo1 == foo2:", merkle.Hash(foo1) == merkle.Hash(foo2))
}
```

Which will output:

``` sh
$ go run example/simple.go
foo1 == foo2: false
foo1 == foo2: true
```

This works via reflecting over `Foo`s fields, though performance can be improved if `Foo` implements the `merkle.Hasher` interface:

``` go
func (f *Foo) Hash() merkle.HashVal {
	h := merkle.NewHash()
	h.Write(f.A)
	h.Write(f.B)
	h.Write(f.C)
	return h.SumAndCache(&f.Cache)
}
```
## API

See GoDoc

#### MIT License

Copyright © 2014 Luma Networks <contact@lumanetworks.com>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

