package main

import "fmt"

func main() {
	//fmt.Println("Hello Sajeet, It's 17/10/2021, 2:35 pm")
	fmt.Println("Hello World")
	print()
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			print()
			fmt.Println(i)
		}
	}
}
func print() {
	fmt.Println("Hello Sajeet", 19101372, true)
}
