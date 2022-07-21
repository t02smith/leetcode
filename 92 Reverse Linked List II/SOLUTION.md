# 92. Reverse Linked List II

Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.

## Solution 1

Solution 1 works by reversing a growing sublist within the given list by:

1. Adding a dummy start node incase the first node needs to be shuffled
2. Findind the leftmost node to be shuffled
3. Performing right-left iterations
   1. each iteration expands the reversed subarray and moves the rightmost element to the start
4. return the dummy's next node

**Runtime 0ms -> > 100%** \
**Memory usage 2MB -> < 100%**
