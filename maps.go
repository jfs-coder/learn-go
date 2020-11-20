package main

import "fmt"

func main() {
	vertices := make (map [string] int)

	// this is how you add map entries
	vertices["triangle" ] = 2
	vertices["square"   ] = 3
	vertices["dodecagon"] = 12

	// prints everything in the map
	fmt.Println(vertices)
	// only prints the selected entry "triangles" associated value which is 2
	fmt.Println(vertices["triangle"])

	// use delete to remove something from the map
	delete(vertices, "square")
	fmt.Println(vertices)

}
