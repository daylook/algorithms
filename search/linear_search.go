package search

/*
Problem: Given an array arr[] of n elements, write a function to return the index of a given element x in arr[].
and return -1 , if the element was not found

Solution: A simple approach is to do a linear search, i.e

Start from the leftmost element of arr[] and one by one compare x with each element of arr[]
If x matches with an element, return the index.
If x doesn’t match with any of elements, return -1.

Time complexity: O(n)

Improve Linear Search Worst-Case Complexity
If element Found at last  O(n) to O(1)
If element Not found O(n) to O(n/2)

*/

// LinearSearch with range
func LinearSearch(elements []int, x int) int {
	for index, element := range elements {
		if element == x {
			return index
		}
	}
	return -1
}

// LinearSearch2 implemented with common for loop
func LinearSearch2(elements []int, x int) int {
	for i := 0; i < len(elements); i++ {
		if elements[i] == x {
			return i
		}
	}
	return -1
}

// ImprovedLinearSearch as improved linear searching
func ImprovedLinearSearch(elements []int, x int) int {
	position := -1
	right := len(elements) - 1

	for left := 0; left < right; {

		// If search_element is found with left varaible
		if elements[left] == x {
			position = left
			return position
		}

		// If search_element is found with right varaible
		if elements[right] == x {
			position = right
			return position
		}

		left++
		right--
	}
	return -1
}
