package main

import "fmt"

var a = 120
var z int // data type is of int

func main() {
	var y = 43 * 2
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(z)
	print()
}
func print() {
	// fmt.Println(y) not accessible
	fmt.Println(a)
}
