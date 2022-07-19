# 118. Pascal's Triangle

Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown...

## Solution 1

This first solution iteratively works out the tree layer by layer by looking at the parents of each
element.

**Runtime 2ms -> > 42.47%** \
**Memory usage 2.1MB -> < 57.83%**

## Solution 2

A second solution could also be made by using the notion of binomial distribution to work out each element
in the tree where an element in row i, column j is found using:

iCj = (i!)/(j!(i-j)!)

However computing factorials can be quite expensive given that we need the entire tree anyway and can just
use the numbers we've already generated. This formula would be useful if we only needed a single row or element.
