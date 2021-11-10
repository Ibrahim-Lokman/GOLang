package main

import (
	"fmt"
)

var y = 43
var z = "Ibrahim Lokman Sajeet"
var a = `Sajeet said, "Blah Blah Blah"`
var zerovalue string

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) //type is int

	fmt.Println(z)
	fmt.Printf("%T\n", z) //type is String

	// z = 43 is not possible, cause z is declared string

	fmt.Println(a)
	fmt.Printf("%T\n", a) //type is String

	fmt.Println("value at y is ", zerovalue, ".")
	zerovalue = "mew mew mew"
	fmt.Println("value at y is ", zerovalue, ".")
}
