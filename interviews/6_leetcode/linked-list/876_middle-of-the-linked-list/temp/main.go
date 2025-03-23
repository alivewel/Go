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
	// list2 := &ListNode{1,
	// 	&ListNode{2,
	// 		&ListNode{3,
	// 			&ListNode{4,
	// 				&ListNode{5,
	// 					&ListNode{6, nil}}}}}}
	list2 := &ListNode{1, nil}

	traverseNode(list1)
	traverseNode(list2)
}

func traverseNode(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}
