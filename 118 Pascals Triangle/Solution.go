package main

import (
	"fmt"
	"reflect"
)

func generate(numRows int) [][]int {
	output := [][]int{{1}}
	var row []int

	for i := 1; i < numRows; i++ {
		row = []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, output[i-1][j]+output[i-1][j-1])
			}
		}

		output = append(output, row)
	}

	return output
}

func main() {
	fmt.Printf("1. 5. expected: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]] correct: %t\n",
		reflect.DeepEqual(generate(5), [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}))

	fmt.Printf("1. 1. expected: [[1]] correct: %t\n",
		reflect.DeepEqual(generate(1), [][]int{{1}}))
}
