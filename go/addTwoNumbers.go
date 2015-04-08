package main

import (
	"fmt"
)

type ListNode struct {
	val  int
	next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	carry := 0
	sum := &ListNode{0, nil}
	current := sum

	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		val := 0
		if l1 != nil {
			v1 = l1.val
			l1 = l1.next
		}

		if l2 != nil {
			v2 = l2.val
			l2 = l2.next
		}

		val = (v1 + v2 + carry) % 10
		current.next = &ListNode{val, nil}
		current = current.next

		carry = (v1 + v2 + carry) / 10
	}

	if carry > 0 {
		current.next = &ListNode{carry, nil}
	}

	return sum.next
}

func main() {
	l1 := &ListNode{1, &ListNode{8, nil}}
	l2 := &ListNode{0, nil}

	sum := addTwoNumbers(l1, l2)
	for sum != nil {
		fmt.Println(sum.val)
		sum = sum.next
	}
}
