# go-murmur
Go Package for MurmurHash2

## Example

```go
package main

import (
	mm "github.com/mistweaverco/go-murmur/murmur2"
)

func main() {
	h := mm.MurmurHash2([]byte("hello"), 1)
	println(h)
}
```
