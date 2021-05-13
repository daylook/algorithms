package backtracking

/*
Backtracking
In real life, let us suppose someone gave you a lock with a number (three digit lock, number
range from 1 to 9). Moreover, you do not have the exact password key for the lock. You need to
test every combination until you got the right one. Obviously, you need to test starting from
something like “111”, then “112” and so on. You will get your key before you reach “999”. Therefore,
what you are doing is backtracking.

Suppose the lock produce some sound “click” correct digit is selected for any level. If we can
listen to this sound such intelligence/ heuristics will help you to reach your goal much faster.
These functions are called Pruning function or bounding functions.

Backtracking is a method by which solution is found by exhaustively searching through large but finite
number of states, with some pruning or bounding function that will narrow down our search.

For all the problems like (NP hard problems) for which there does not exist any other method we
use backtracking.

Backtracking problems have the following components:
1. Initial state
2. Target / Goal state
3. Intermediate states
4. Path from the initial state to the target / goal state
5. Operators to get from one state to another
6. Pruning function (optional)

The solving process of backtracking algorithm starts with the construction of state’s tree, whose
nodes represents the states. The root node is the initial state and one or more leaf node will be
our target state. Each edge of the tree represents some operation. The solution is obtained by
searching the tree until a Target state is found.

Backtracking uses depth-first search:
1) Store the initial state in a stack
2) While the stack is not empty, repeat:
3) Read a node from the stack.
4) While there are available operators, do:
a. Apply an operator to generate a child
b. If the child is a goal state – return solution
c. If it is a new state, and pruning function does not discard it push the child into the stack.

There are three monks and three demons at one side of a river. We want to move all of them to the
other side using a small boat. The boat can carry only two persons at a time. Given if on any shore
the number of demons will be more than monks then they will eat the monks. How can we move all of
these people to the other side of the river safely?

Same as the above problem there is a farmer who has a goat, a cabbage and a wolf. If the farmer
leaves, goat with cabbage, goat will eat the cabbage. If the farmer leaves wolf alone with goat,
wolf will kill the goat. How can the farmer move all his belongings to the other side of the river?

You are given two jgs, a 4-gallon one and a 3-gallon one. There are no measuring markers on jugs.
A tap can be used to fill the jugs with water. How can you get 2 gallons of water in the 4-gallon
jug?
*/
