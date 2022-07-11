# 199. Binary Tree Right Side View - 11/7/2022

Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

## Solution 1

This first solution is a greedy DFS search which will find the rightmost child at each depth by recursively
searching down the rightmost path. At each level if there is no current value stored for that depth the current
node's value will be used.

**Runtime 0ms -> > 100%** \
**Memory usage 2.3MB -> < 14.37%**

## Solution 2

Solution 2 looks at each layer in the tree and will choose the rightmost node's value and add it to the list.

The algorithm maintains a list of nodes at the current level and at each iteration, each node is expanded into
its left and right child if they exist.

**Runtime 6ms -> > 16.06%** \
**Memory usage 2.3MB -> < 76.62%**

*This solution is nowhere near as fast as Solution 1 and uses basically the same memory*
