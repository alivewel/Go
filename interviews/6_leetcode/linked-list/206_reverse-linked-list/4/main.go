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
    var prev *ListNode
    curr := head
    for curr != nil {
        nextNode := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextNode
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
