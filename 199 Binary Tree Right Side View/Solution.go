package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SOLUTION 1 - DFS

func rightSideView1(root *TreeNode) []int {
	var view *[]int = &[]int{}
	rightSideViewRec(root, view, 0)

	return *view
}

func rightSideViewRec(root *TreeNode, view *[]int, depth int) {
	if root == nil {
		return
	}

	if len(*view) == depth {
		*view = append(*view, root.Val)
	}

	rightSideViewRec(root.Right, view, depth+1)
	rightSideViewRec(root.Left, view, depth+1)
}

// SOLUTION 2 - Iterative

func rightSideView2(root *TreeNode) []int {
	var view []int = []int{}
	if root == nil {
		return view
	}

	var level []*TreeNode = []*TreeNode{root}
	for len(level) > 0 {
		view = append(view, level[len(level)-1].Val)

		var newLevel []*TreeNode = []*TreeNode{}
		for _, t := range level {
			if t.Left != nil {
				newLevel = append(newLevel, t.Left)
			}

			if t.Right != nil {
				newLevel = append(newLevel, t.Right)
			}
		}

		level = newLevel
	}

	return view
}

func main() {

	fmt.Printf("1. [1,2,3,null,5,null,4]. expected: [1,3,4] correct: %t\n",
		reflect.DeepEqual(rightSideView2(t1), []int{1, 3, 4}))

	fmt.Printf("2. [1,null,3]. expected: [1,3] correct: %t\n",
		reflect.DeepEqual(rightSideView2(t2), []int{1, 3}))

	fmt.Printf("[]. expected: [] correct: %t\n",
		reflect.DeepEqual(rightSideView2(nil), []int{}))

	fmt.Printf("[1,2]. expected: [1,2] correct: %t\n",
		reflect.DeepEqual(rightSideView2(t3), []int{1, 2}))
}

var t1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val:  3,
		Left: nil,
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	},
}

var t2 = &TreeNode{
	Val:  1,
	Left: nil,
	Right: &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	},
}

var t3 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	},
	Right: nil,
}
