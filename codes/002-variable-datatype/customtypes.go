package main

import (
	"fmt"
)

type hotdog int

var bz hotdog

func main() {
	bz = 100
	fmt.Println(bz)
	fmt.Printf("%T\n", bz)
}
