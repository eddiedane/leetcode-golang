package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Printf("%v\n", middleNode(&ListNode{Val: 1}))
}

func middleNode(head *ListNode) *ListNode {
	var mid *ListNode = head
	var end *ListNode = head

	for end != nil && end.Next != nil {
		mid = mid.Next
		end = end.Next.Next
	}

	return mid
}
