# 473. Matchsticks to Square

You are given an integer array matchsticks where matchsticks[i] is the length of the ith matchstick. You want to use all the matchsticks to make one square. You should not break any stick, but you can link them up, and each matchstick must be used exactly one time.

Return true if you can make this square and false otherwise.

## Solution 1 - DFS

This problem is clearly derivative of the bin packing
problem which is known to be NP-Complete as such a non
exponential solution isn't possible. Given that this
problem has the constraint ```1 <= matchsticks.length <= 15```, DFS is an acceptable solution.

Solution 1 uses DFS to look at possible solutions by greedily adding to the first non-full side. \
Initially the algorithm checks for the possible base cases:

- the array has less than 4 items => false
- the sum of the array is not a multiple of 4 => false

**Runtime 175ms -> > 14.29%** \
**Memory usage 2.2MB -> < 46.43%**
