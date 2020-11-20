package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var x int
	var y int = 17
	z := 88
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	if z > 80 {
		fmt.Println("Z is larger than 80")
	} else if z < 90 {
		fmt.Println("Z is smaller than 90")
	} else {
		fmt.Println("Z is out of our range")
	}
}
