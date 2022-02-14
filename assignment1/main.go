package main

import (
	"fmt"
)

func main() {
	ints := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	for i := range ints {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
