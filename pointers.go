package main

import "fmt"

func main() {
	i := 7 
	j := 10
	// this will not increment i because it has to be passed by address for it to work, which we do after these lines
	increment(i)
	fmt.Println(i)
	incrementRef(&j)
	fmt.Println(j)
}

func increment(x int) {
	x++
}

func incrementRef(x *int) {
	// de-reference the pointer, else it will just increment the memory address itself, which is pointless
	// when you de-reference, you are telling the compiler to access the value AT the memory address and increment it by 1 (++), not add 1 to the memory address value itself.
	*x++
}
