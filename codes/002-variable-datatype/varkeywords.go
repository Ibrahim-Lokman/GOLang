package main

import "fmt"

var a = 120

func main() {
	var y = 43 * 2
	fmt.Println(y)
	fmt.Println(a)
	print()
}
func print() {
	// fmt.Println(y) not accessable
	fmt.Println(a)
}
