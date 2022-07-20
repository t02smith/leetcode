# 792. Number of Matching Subsequences

Given a string s and an array of strings words, return the number of words[i] that is a subsequence of s.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

## Solution 1

Solution 1 looks at whether each given subsequence is actually a subsequence of the input string.

It does this by looking at each character in the subseq in order and seeing if:

- it is in the input string
- it comes after the last found letter in the string

**Runtime 2246ms -> > 19.72%** \
**Memory usage 7.3MB -> < 73.24%**

## Solution 2

Solution 2 stores the positions of each character within the input string and will then iterate over the
list of words and see if they are valid subsequences. It does this by iterating over each character and, if
that character is stored in memory, look for an index in the original string that is higher than the last
used one.

**Runtime 328ms -> > 46.48%** \
**Memory usage 9.6MB -> < 16.90%**
