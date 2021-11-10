package main

import (
	"fmt"
)

type sajeet int

var x sajeet

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
