package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for slow !

    return false
}

func main() {
	// Список: 1->2->4
	// node := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}

	// Список: 1->2->4
	node := &ListNode{1, nil}

	// Список: 4->2->1
	fmt.Println(hasCycle(node))
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

// проверить алгоритм при списке
// 2, 0, 6, 3, 6, 1, 3, 1
