package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
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

func main() {
	// Создание списка [1, 2, 3, 4, 5]
	list1 := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil}}}}}

	// Создание списка [1, 2, 3, 4, 5, 6]
	list2 := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5,
						&ListNode{6, nil}}}}}}

	fmt.Println(mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	res := dummy
	for list1 != nil && list2 != nil {
		temp := &ListNode{}
		if list1.Val > list2.Val {
			temp = &ListNode{Val: list2.Val}
			list2 = list2.Next
		} else {
			temp = &ListNode{Val: list1.Val}
			list1 = list1.Next
		}
		dummy.Next = temp
		dummy = dummy.Next
	}
	if list1 != nil {
		dummy.Next = list1
	}
	if list2 != nil {
		dummy.Next = list2
	}
	return res.Next
}
