package main

import (
	"fmt"

	"github.com/daylook/algorithms/searching"
)

func main() {
	fmt.Println("Choose each algorithm's function and run it")

	elements := []int{1, 2, 3, 4, 5, 6, 7}
	ind := searching.ImprovedLinearSearch(elements, 7)
	fmt.Println(ind)
}
