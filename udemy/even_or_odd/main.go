package main

import "fmt"

func main() {
	sliceInt := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, x := range sliceInt {
		if x%2 == 0 {
			fmt.Printf("%d is even\n", x)
		} else {
			fmt.Printf("%d is odd\n", x)
		}
	}
}
