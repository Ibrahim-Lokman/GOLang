package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//assigning printed string in a variable by Sprintf
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
