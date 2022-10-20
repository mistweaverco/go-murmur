package main

import (
	mm "github.com/mistweaverco/go-murmur/murmur2"
)

func main() {
	mm.MurmurHash2([]byte("hello"), 1)
	println("Hello, World!")
}
