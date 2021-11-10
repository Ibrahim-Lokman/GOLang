package main

import (
	"fmt"
)

type sajeet int

var x sajeet
var y int

func main() {
	x = 42
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
