package main

import (
	"fmt"

	"github.com/daylook/algorithms/search"
)

func main() {
	fmt.Println("Select algorithm's function and run it")

	elements := []int{1, 2, 3, 4, 5, 6, 7}
	ind := search.RecursiveBinarySearch(elements, 1, 0, 6)
	fmt.Println(ind)
}
