package main

import "fmt"

var aa = 120
var zz int // data type is of int

func main() {
	var y = 43 * 2
	fmt.Println(y)
	fmt.Println(aa)
	fmt.Println(zz)
	print()
}
func print() {
	// fmt.Println(y) not accessible
	fmt.Println(a)
}
