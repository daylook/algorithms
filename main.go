package main

import (
	"fmt"

	"github.com/daylook/algorithms/divideConquer"
)

func main() {
	data := []int{70, 250, 50, 80, 140, 12, 14}
	max := divideConquer.MinInSlice(data, 0, len(data))
	min := divideConquer.MaxInSlice(data, 0, len(data))

	fmt.Printf("Maximum number in a given slice:%d\n", max)
	fmt.Printf("Minimum number in a given slice:%d\n", min)
}
