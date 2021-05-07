/*
Divide and Conquer Algorithm
https://www.geeksforgeeks.org/divide-and-conquer-algorithm-introduction/

This technique can be divided into the following three parts:

Divide: This involves dividing the problem into some sub problem.
Conquer: Sub problem by calling recursively until sub problem solved.
Combine: The Sub problem Solved so that we will get find problem solution.

The following are some standard algorithms that follows Divide and Conquer algorithm.

1. Quicksort is a sorting algorithm. The algorithm picks a pivot element, rearranges the array elements in
such a way that all elements smaller than the picked pivot element move to left side of pivot, and all
greater elements move to right side. Finally, the algorithm recursively sorts the subarrays on left and right
of pivot element.

2. Merge Sort is also a sorting algorithm. The algorithm divides the array in two halves, recursively sorts
them and finally merges the two sorted halves.

3. Closest Pair of Points The problem is to find the closest pair of points in a set of points in x-y plane.
The problem can be solved in O(n^2) time by calculating distances of every pair of points and comparing the
distances to find the minimum. The Divide and Conquer algorithm solves the problem in O(nLogn) time.

4. Strassen’s Algorithm is an efficient algorithm to multiply two matrices. A simple method to multiply two
matrices need 3 nested loops and is O(n^3). Strassen’s algorithm multiplies two matrices in O(n^2.8974) time.

5. Cooley–Tukey Fast Fourier Transform (FFT) algorithm is the most common algorithm for FFT. It is a divide
and conquer algorithm which works in O(nlogn) time.

6. Karatsuba algorithm for fast multiplication it does multiplication of two n-digit numbers in at most
single-digit multiplications in general (and exactly when n is a power of 2). It is therefore faster than
the classical algorithm, which requires n^2 single-digit products. If n = 2^10 = 1024, in particular, the exact
counts are 3^10 = 59, 049 and (2^10)^2 = 1, 048, 576, respectively.

What does not qualifies as Divide and Conquer:
Binary Search is a searching algorithm. In each step, the algorithm compares the input element x with the
value of the middle element in array. If the values match, return the index of the middle. Otherwise, if x
is less than the middle element, then the algorithm recurs for left side of the middle element, else recurs
for the right side of the middle element. Contrary to popular belief, this is not an example of Divide and
Conquer because there is only one sub-problem in each step (Divide and conquer requires that there must
be two or more sub-problems) and hence this is a case of Decrease and Conquer.

Divide And Conquer algorithm :

DAC(a, i, j)
{
    if(small(a, i, j))
      return(Solution(a, i, j))
    else
      m = divide(a, i, j)               // f1(n)
      b = DAC(a, i, mid)                 // T(n/2)
      c = DAC(a, mid+1, j)            // T(n/2)
      d = combine(b, c)                 // f2(n)
   return(d)
}

Recurrence Relation for DAC algorithm :
This is recurrence relation for above program.
T(n) =
    O(1) if n is small
    f1(n) + 2T(n/2) + f2(n)

Example: To find the maximum(250) and minimum(12) element in a given array, using
a divide and conquer approach(DAC) which has three steps divide, conquer and combine.
Input: { 70, 250, 50, 80, 140, 12, 14 }

Divide and Conquer (D & C) vs Dynamic Programming (DP)
Both paradigms (D & C and DP) divide the given problem into subproblems and solve subproblems.
How to choose one of them for a given problem? Divide and Conquer should be used when the same
subproblems are not evaluated many times. Otherwise Dynamic Programming or Memoization should
be used. For example, Quicksort is a Divide and Conquer algorithm, we never evaluate the same
subproblems again. On the other hand, for calculating the nth Fibonacci number, Dynamic
Programming should be preferred

*/

package dandc

// We are using the recursive approach to find max/min where we will see that only two
// elements are left and then:
// i.e. if a[index]>a[index+1]

func MaxInSlice(data []int, index int, left int) int {
	var max int

	// to check the condition that there will be two-element in the left
	if index >= left-2 {
		if data[index] > data[index+1] {
			return data[index]
		} else {
			return data[index+1]
		}
	}

	max = MaxInSlice(data, index+1, left)

	if data[index] > max {
		return data[index]
	} else {
		return max
	}
}

func MinInSlice(data []int, index int, left int) int {
	var min int

	// to check the condition that there will be two-element in the left
	if index >= left-2 {
		if data[index] < data[index+1] {
			return data[index]
		} else {
			return data[index+1]
		}
	}

	min = MinInSlice(data, index+1, left)

	if data[index] < min {
		return data[index]
	} else {
		return min
	}
}

/*
func main() {
	data := []int{70, 250, 50, 80, 140, 12, 14}
	max := SliceMax(data, 0, len(data))
	min := SliceMin(data, 0, len(data))

	fmt.Printf("Maximum number in a given slice:%d\n", max)
	fmt.Printf("Minimum number in a given slice:%d\n", min)
}
*/
