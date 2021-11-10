package main

import (
	"fmt"
)

var ya = 16

func main() {
	fmt.Println("Hello, playground")
	fmt.Printf("%T\n", ya)
	fmt.Printf("%b\n", ya)
	fmt.Printf("%x\n", ya)
	fmt.Printf("%#x", ya)
	fmt.Printf("%#x\t%b\t%x", ya, ya, ya)

}
