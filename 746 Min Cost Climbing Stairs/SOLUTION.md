# 746. Min Cost Climbing Stairs - 10/7/2022

## problem

You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.

## Solution 1

Find the minimum cost to get to each index i by looking at i-1 and i-2. The result will be stored at index |list|.

Base cases:

- 0 == 0
- 1 == 0

**Runtime 101ms -> > 36.32%** \
**Memory usage 13.9MB -> < 75.93%**

## Solution 2

Same method as Solution 1 but changes items in the cost array in-place to be the cost to take a step from that index.

e.g. [10, 15, 20] -> [10, 15, 30] => min(15, 30) = 15

**Runtime 77ms -> > 69.60%** \
**Memory usage 13.9MB -> < 75.93%**

### Golang

A version of this solution written in Golang was added.

**Runtime 2ms -> > 92.45%** \
**Memory usage 2.9MB -> < 87.08%**
