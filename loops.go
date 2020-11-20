package main

import "fmt"

func main() {
	// the classic 'for' loop
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	// in go, there is no 'while' loop, we just modify the for loop to make it a while loop
	j := 3
	for j >= 0  {
		fmt.Println(j)
		j--
	}

	// declare an array 'arr' and use the for range syntax
	arr := [] string {"a","b","c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	// declare a map and use the same syntax again (for ... range)
	mmm := make(map[string]string)
	mmm["a"] = "alpha"
	mmm["b"] = "bravo"
	mmm["c"] = "charlie"

	for key, value := range arr {
		fmt.Println("key:", key, "value:", value)
	}
}
