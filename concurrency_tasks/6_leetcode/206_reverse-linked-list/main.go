package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil // чтобы избавиться от пустого значения в начале списка
	current := head
	for current != nil {
		tempNext := current.Next // предыдущий узел
		current.Next = prev
		prev = current
		current = tempNext
	}
	return prev
}

func main() {
	// Список: 1->2->4
	node := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}

	// Список: 4->2->1
	fmt.Println(reverseList(node))
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
