package main

import (
	"fmt"
	"sort"
)

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}

	var sum int = 0
	for _, x := range matchsticks {
		sum += x
	}

	if sum%4 != 0 {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
	return depthFirstSearch(sum/4, 0, &matchsticks, &[]int{0, 0, 0, 0})
}

func depthFirstSearch(target, i int, sticks, sides *[]int) bool {
	if i == len(*sticks) {
		return (*sides)[0] == target && (*sides)[1] == target && (*sides)[2] == target && (*sides)[3] == target
	}

	for sideIndex, side := range *sides {
		if side+(*sticks)[i] > target {
			continue
		}

		(*sides)[sideIndex] += (*sticks)[i]
		if depthFirstSearch(target, i+1, sticks, sides) {
			return true
		}

		(*sides)[sideIndex] -= (*sticks)[i]
	}

	return false
}

func main() {
	fmt.Printf("1. [1,1,2,2,2]. expected: true, actual: %t\n", makesquare([]int{1, 1, 2, 2, 2}))
	fmt.Printf("2. [3,3,3,3,4]. expected: false, actual: %t\n", makesquare([]int{3, 3, 3, 3, 4}))
	fmt.Printf("3. [4]. expected: false, actual: %t\n", makesquare([]int{4}))
	fmt.Printf("3. [5,5,5,5,4,4,4,4,3,3,3,3]. expected: true, actual: %t\n", makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}))
	fmt.Printf("3. [1,1,1,1,1,1,1,1,1,1,1,1,1,1,6]. expected: false, actual: %t\n", makesquare([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6}))

}
