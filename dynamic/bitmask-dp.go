/*
Bitmasking and Dynamic Programming

Consider the below problems statement.
There are 100 different types of caps each having a unique id from 1 to 100. Also, there
are ‘n’ persons each having a collection of a variable number of caps. One day all of these
persons decide to go in a party wearing a cap but to look unique they decided that none of
them will wear the same type of cap. So, count the total number of arrangements or ways
such that none of them is wearing the same type of cap.

Constraints: 1 <= n <= 10 Example:

The first line contains the value of n, next n lines contain collections of all the n persons.

Input:
3
5 100 1     // Collection of the first person.
2           // Collection of the second person.
5 100       // Collection of the third person.

Output:
4
Explanation: All valid possible ways are (5, 2, 100),  (100, 2, 5), (1, 2, 5) and  (1, 2, 100)


What is Bitmasking?
Suppose we have a collection of elements which are numbered from 1 to N. If we want to represent
a subset of this set then it can be encoded by a sequence of N bits (we usually call this sequence
a “mask”). In our chosen subset the i-th element belongs to it if and only if the i-th bit
of the mask is set i.e., it equals to 1. For example, the mask 10000101 means that the subset
of the set [1… 8] consists of elements 1, 3 and 8. We know that for a set of N elements there
are total 2^N subsets thus 2^N masks are possible, one representing each subset. Each mask is,
in fact, an integer number written in binary notation.

Our main methodology is to assign a value to each mask (and, therefore, to each subset) and thus
calculate the values for new masks using values of the already computed masks. Usually our main
target is to calculate value/solution for the complete set i.e., for mask 11111111. Normally, to
find the value for a subset X we remove an element in every possible way and use values for
obtained subsets X’1, X’2… ,X’k to compute the value/solution for X. This means that the values
for X’i must have been computed already, so we need to establish an ordering in which masks will
be considered. It’s easy to see that the natural ordering will do: go over masks in increasing
order of corresponding numbers. Also, We sometimes, start with the empty subset X and we add
elements in every possible way and use the values of obtained subsets X’1, X’2… ,X’k to compute
the value/solution for X.

We mostly use the following notations/operations on masks:
bit(i, mask) – the i-th bit of mask
count(mask) – the number of non-zero bits in the mask
first(mask) – the number of the lowest non-zero bit in the mask
set(i, mask) – set the ith bit in the mask
check(i, mask) – check the ith bit in the mask

How is this problem solved using Bitmasking + DP?
The idea is to use the fact that there are upto 10 persons. So we can use an integer variable
as a bitmask to store which person is wearing a cap and which is not.

Let i be the current cap number (caps from 1 to i-1 are already processed). Let integer variable
mask indicates that the persons wearing and not wearing caps.  If i'th bit is set in mask, then
i'th person is wearing a cap, else not.

// consider the case when ith cap is not included in the arrangement
countWays(mask, i) = countWays(mask, i+1) +

// when ith cap is included in the arrangement so, assign this cap to all possible persons
// one by one and recur for remaining persons.
∑ countWays(mask | (1 << j), i+1)
    for every person j that can wear cap i

Note that the expression "mask | (1 << j)" sets j'th bit in mask. And a person can wear cap i if
it is there in the person's cap list provided as input.

If we draw the complete recursion tree, we can observe that many subproblems are solved again
and again. So we use Dynamic Programming. A table dp[][] is used such that in every entry
dp[i][j], i is mask and j is cap number.

Since we want to access all persons that can wear a given cap, we use an array of vectors,
capList[101]. A value capList[i] indicates the list of persons that can wear cap i.

Below is the implementation of above idea.



*/

package dynamic
