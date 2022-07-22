# 86. Partition List

Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

## Solution 1

Solution 1 splits the list into two separate lists:

- elements that are < x
- elements that are >= x

It will iterate over the linked list and add elements to the correct
list such that elements keep their relative order but are separated
by the value x.

The two lists are then joined together.

**Runtime 5ms -> > 28.70%** \
**Memory usage 2.4MB -> < 18.52%**
