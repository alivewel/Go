package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	for list1 != nil && list2 != nil {
		dummy := &ListNode{}
		if list2.Val > list1.Val {
			dummy = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			dummy = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		res.Next = dummy
		res = res.Next
	}
	if list1 != nil {
		res.Next = list1
	}
	if list2 != nil {
		res.Next = list2
	}

	return head.Next
}

func main() {
	// Список: 1->2->4
	first := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	// Список: 1->3->4
	second := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	fmt.Println(mergeTwoLists(first, second)) // Список: 1->1->2->3->4->4
}

func (it *ListNode) String() (result string) {
	for cur := it; cur != nil; cur = cur.Next {
		if cur == cur.Next {
			return "CYCLE DETECTED"
		}
		result += strconv.Itoa(cur.Val) + " "
	}
	return result
}
