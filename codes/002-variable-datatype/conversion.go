package main

import (
	"fmt"
)

type hotdog1 int

var by hotdog1
var bx int

func main() {
	by = 100
	bx = 200
	fmt.Println(by)
	fmt.Printf("%T\n", by)
	bx = int(by)
	fmt.Println(bx)
	fmt.Printf("%T\n", bx)
}
