package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	fmt.Println(sum(482, 100))

	// calling this, we get two values and they are set into 'result' and 'err'
	// if err is != to nil that means something went wrong (we passed a negative function to sqrt) else we print the result (in this case, 8 is the sqrt of 64)
	// try change the 64 to -16 and you'll see it print the error message
	result, err := sqrt(64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sum (x int, y int) int {
	return x + y
}

// this function can return TWO values (flaot64 and error (built-in types)
// if x is less than 0, so negative, it will return 0 as the answer and the error message "Undefined for negative numbers"
// otherwise it will return the answer + 'nil' for the error (so that we can check for that later and print an error message if the returned error is != nil
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
