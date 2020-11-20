package main

import ("fmt")

func main() {
	var norm [5] int
	alt := [5] int {5,4,3,2,1}
	slice_of_ints := [] int {1,2,3,4,5} 

	norm[2] = 7
	fmt.Println(norm)
	fmt.Println(alt)
	fmt.Println(slice_of_ints)

	// append to slice_of_ints, can't do with normal arrays
	slice_of_ints = append(slice_of_ints,13)
	fmt.Println(slice_of_ints)
}
