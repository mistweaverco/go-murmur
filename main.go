package main

import (
	mm "github.com/mistweaverco/go-murmur/murmur2"
)

func main() {
	h := mm.MurmurHash2([]byte("hello"), 1)
	println(h) // prints 2788266382
}
