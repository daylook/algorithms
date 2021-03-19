package search

/*
Given a sorted array arr[] of n elements, return the index of a given element x in arr[].

Binary Search: Search a sorted array by repeatedly dividing the search interval in half.
Begin with an interval covering the whole array. If the value of the search key is less
than the item in the middle of the interval, narrow the interval to the lower half.
Otherwise narrow it to the upper half. Repeatedly check until the value is found or the
interval is empty.

time complexity: O(Log n).

*/

// RecursiveBinarySearch as Recursive implementation of Binary Search
// li : left index
// ri : right index
func RecursiveBinarySearch(elements []int, x, li, ri int) int {
	if ri >= li {
		mid := li + (ri-li)/2
		if elements[mid] == x {
			return mid
		} else if elements[mid] > x {
			return RecursiveBinarySearch(elements, x, li, mid-1)
		}
		return RecursiveBinarySearch(elements, x, mid+1, ri)
	}
	return -1
}

func IterativeBinarySearch(elements []int, x int) int {
	left := 0
	right := len(elements) - 1

	for left <= right {
		mid := left + (right-left)/2
		if elements[mid] == x {
			return mid
		} else if elements[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
