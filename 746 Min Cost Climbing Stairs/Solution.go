package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}

	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func main() {
	fmt.Printf("[1,100,1,1,1,100,1,1,100,1]: %d\n", minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	fmt.Printf("[10,15,20]: %d\n", minCostClimbingStairs([]int{10, 15, 20}))
}
