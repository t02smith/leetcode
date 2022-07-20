package main

import (
	"fmt"
)

// solution 1

func numMatchingSubseq(s string, words []string) int {
	count := 0
	for _, ss := range words {
		if isSubseq(s, ss) {
			count++
		}
	}

	return count
}

func isSubseq(str, subStr string) bool {
	if len(subStr) > len(str) {
		return false
	}

	strPointer := 0
	for ssPointer := 0; ssPointer < len(subStr); ssPointer++ {
		for strPointer < len(str) && subStr[ssPointer] != str[strPointer] {
			strPointer++
		}

		if strPointer == len(str) {
			return false
		}

		if subStr[ssPointer] == str[strPointer] {
			strPointer++
		}
	}

	return true
}

// solution 2

func numMatchingSubseq2(s string, words []string) int {
	mem := make(map[byte][]int)
	for c := 0; c < len(s); c++ {
		if _, ok := mem[s[c]]; ok {
			mem[s[c]] = append(mem[s[c]], c)
		} else {
			mem[s[c]] = []int{c}
		}
	}

	count := 0

	for _, str := range words {
		if isSubseq2(&mem, str) {
			count++
		}
	}

	return count
}

func isSubseq2(chars *map[byte][]int, ss string) bool {
	lastUsed, tmpLastUsed := -1, -1

	for c := 0; c < len(ss); c++ {
		tmpLastUsed = lastUsed

		// char not found
		if _, ok := (*chars)[ss[c]]; !ok {
			return false
		}

		for _, v := range (*chars)[ss[c]] {
			if v > lastUsed {
				lastUsed = v
				break
			}
		}

		if lastUsed == tmpLastUsed {
			return false
		}

	}

	return true
}

func isSubsetTests() {
	fmt.Println("isSubset Tests -------------------------------")
	fmt.Printf("1. input (abcde, ace), expected: true, actual: %t\n", isSubseq("abcde", "ace"))
	fmt.Printf("2. input (dsahjpjauf, dhjjf), expected: true, actual: %t\n", isSubseq("dsahjpjauf", "dhjjf"))
	fmt.Printf("3. input (dsahjpjauf, aaa), expected: false, actual: %t\n", isSubseq("dsahjpjauf", "aaa"))
	fmt.Printf("4. input (dsahjpjauf, pjx), expected: false, actual: %t\n", isSubseq("abcde", "pjx"))
	fmt.Printf("5. input (a, pjx), expected: false, actual: %t\n", isSubseq("a", "pjx"))
	fmt.Printf("6. input (abcdef, ), expected: true, actual: %t\n", isSubseq("abcdef", ""))
}

func numMatchingSubseqTests() {
	fmt.Println("numMatchingSubseq Tests -------------------------------")
	fmt.Printf("1. input abcde, [a,bb,acd,ace], expected: 3, actual: %d\n", numMatchingSubseq2("abcde", []string{"a", "bb", "acd", "ace"}))
	fmt.Printf("2. input dsahjpjauf, [ahjpjau,ja,ahbwzgqnuk,tnmlanowax], expected: 2, actual: %d\n", numMatchingSubseq2("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}))
}

func main() {
	//isSubsetTests()
	fmt.Println()
	numMatchingSubseqTests()

}
