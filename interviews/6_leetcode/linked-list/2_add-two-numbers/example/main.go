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

func createList(n int) *ListNode {
    if n == 0 {
        return nil
    }

    head := &ListNode{Val: 1}
    current := head

    for i := 2; i <= n; i++ {
        current.Next = &ListNode{Val: i}
        current = current.Next
    }

    return head
}

func printList(head *ListNode) {
    for node := head; node != nil; node = node.Next {
        fmt.Println(node.Val)
    }
}


func main() {
    head := createList(5)
    printList(head)
}

