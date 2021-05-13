package branchbound

/*
Branch-and-bound
Branch and bound method is used when we can evaluate cost of visiting each node by a utility
functions. At each step, we choose the node with lowest cost to proceed further. Branch-and
bound algorithms are implemented using a priority queue. In branch and bound, we traverse the
nodes in breadth-first manner.

A* Algorithm
A* is sort of an elaboration on branch-and-bound. In branch-and-bound, at each iteration we
expand the shortest path that we have found so far. In A*, instead of just picking the path with
the shortest length so far, we pick the path with the shortest estimated total length from start
to goal, where the total length is estimated as length traversed so far plus a heuristic estimate
of the remaining distance from the goal.

Branch-and-bound will always find an optimal solution, which is shortest path. A* will always
find an optimal solution if the heuristic is correct. Choosing a good heuristic is the most
important part of A* algorithm.

*/
