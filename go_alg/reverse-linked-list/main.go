package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{}
	next := head.Next
	for next != nil {
		next = head.Next
		head.Next = newHead
		newHead = next
	}
	return newHead
}

func main() {
	// tests
	end := ListNode{
		Val:  5,
		Next: nil,
	}
	f := ListNode{
		Val:  4,
		Next: &end,
	}
	t := ListNode{
		Val:  3,
		Next: &f,
	}
	s := ListNode{
		Val:  2,
		Next: &t,
	}
	head := ListNode{
		Val:  1,
		Next: &s,
	}

	newHea := reverseList(&head)

	for newHea.Next != nil {
		fmt.Println(head.Val)
		head = *head.Next
	}
}
