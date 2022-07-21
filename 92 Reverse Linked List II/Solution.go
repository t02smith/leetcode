package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// empty list
	if head == nil {
		return nil
	}

	// no shuffling
	if left == right {
		return head
	}

	// dummy head at start of list
	newHead := &ListNode{0, head}
	l := newHead

	// find leftmost node to reverse
	for i := 0; i < left-1; i++ {
		l = l.Next
	}

	// iteratively move nodes from the right to the left
	// 1-2-3-4-5 -> 1-3-2-4-5 -> 1-4-3-2-5
	start, end := l.Next, l.Next.Next
	for i := 0; i < right-left; i++ {
		start.Next = end.Next
		end.Next = l.Next
		l.Next = end
		end = start.Next
	}

	return newHead.Next
}

func main() {
	res := reverseBetween(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2, 4)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
	fmt.Println()

	res = reverseBetween(&ListNode{1, nil}, 1, 1)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
	fmt.Println()

	res = reverseBetween(&ListNode{3, &ListNode{5, nil}}, 1, 2)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}
