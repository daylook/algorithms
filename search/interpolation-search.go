package search

/*
The Interpolation Search is an improvement over Binary Search for instances, where the values
in a sorted array are uniformly distributed. Binary Search always goes to the middle element
to check. On the other hand, interpolation search may go to different locations according to
the value of the key being searched. For example, if the value of the key is closer to the
last element, interpolation search is likely to start search toward the end side.

To find the position to be searched, it uses following formula.

The idea of formula is to return higher value of pos
when element to be searched is closer to arr[hi]. And
smaller value when closer to arr[lo]

pos = lo + [ (x-arr[lo])*(hi-lo) / (arr[hi]-arr[Lo]) ]

arr[] ==> Array where elements need to be searched
x     ==> Element to be searched
lo    ==> Starting index in arr[]
hi    ==> Ending index in arr[]

The formula for pos can be derived as follows.

Let's assume that the elements of the array are linearly distributed.

General equation of line : y = m*x + c.
y is the value in the array and x is its index.

Now putting value of lo,hi and x in the equation
arr[hi] = m*hi+c ----(1)
arr[lo] = m*lo+c ----(2)
x = m*pos + c     ----(3)

m = (arr[hi] - arr[lo] )/ (hi - lo)

subtracting eqxn (2) from (3)
x - arr[lo] = m * (pos - lo)
lo + (x - arr[lo])/m = pos
pos = lo + (x - arr[lo]) *(hi - lo)/(arr[hi] - arr[lo])


*/
