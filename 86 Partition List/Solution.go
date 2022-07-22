package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	// 1. create two lists for elems <x and >=x

	dummySmall := &ListNode{
		Val:  0,
		Next: nil,
	}
	dummySmallPointer := dummySmall

	dummyLarge := &ListNode{
		Val:  0,
		Next: nil,
	}
	dummyLargePointer := dummyLarge

	// 2. iteratively add elements to their respective list

	curr := head
	for curr != nil {
		if curr.Val < x {
			dummySmall.Next = curr
			dummySmall = dummySmall.Next
		} else {
			dummyLarge.Next = curr
			dummyLarge = dummyLarge.Next
		}

		curr = curr.Next
	}

	// 3. join both lists
	dummySmall.Next = dummyLargePointer.Next
	dummyLarge.Next = nil

	return dummySmallPointer.Next
}

func main() {
	fmt.Print("1. input: [1,4,3,5,2] 3, expected [1,2,2,4,3,5], actual: ")
	res := partition(&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}, 3)
	for res != nil {
		fmt.Printf("%d, ", res.Val)
		res = res.Next
	}
	fmt.Println()

	fmt.Print("2. input: [2, 1] 2, expected [1,2], actual: ")
	res = partition(&ListNode{2, &ListNode{1, nil}}, 2)
	for res != nil {
		fmt.Printf("%d, ", res.Val)
		res = res.Next
	}
	fmt.Println()
}
