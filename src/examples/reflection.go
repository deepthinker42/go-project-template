package main

import (
	"fmt"
)

type blahdoo struct {
    a string
    i int
}

func main() {
    b := blahdoo{a: "this that", i: -1}
    fmt.Printf("%T, %T, %T", b.a, b.i, b)
}

