package transformconquer

/*
Reduction / Transform-and-Conquer: First, the problem is transformed into a known problem for
which we know optimal solution and then, the problem is solved.

The most common types of transformation are sort of a list. For example, Given a list of numbers
finds the two closest number. The brute force solution for this problem will take distance between
each element in the list and will try to keep the minimum distance pair; this approach will have
a Time Complexity of O(n^2 )

Transform and conquer solution, will be first sort the list in O(nlogn) time and then find the
closest number by scanning the list in another O(n). Which will give the total Time Complexity
of O(nlogn). Examples:
· Gaussian elimination
· Heaps and Heapsort

*/
