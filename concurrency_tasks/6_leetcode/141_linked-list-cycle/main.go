package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func hasCycle(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	}

// 	// slow := head.Next

// 	// fast := slow.Next

// 	slow := &ListNode{}
// 	fast := &ListNode{}

// 	if head.Next != nil {
// 		slow = slow.Next
// 	}

// 	if slow.Next != nil {
// 		fast = slow.Next
// 	}

// 	// fast := slow.Next
// 	// if fast != nil {
// 	// 	fast = fast.Next
// 	// }

// 	// for slow != nil && fast != nil && fast.Next != nil {
// 	for slow != nil && fast != nil {
// 		if slow == fast {
// 			fmt.Println(slow, fast)
// 			return true
// 		}

// 		slow = slow.Next
// 		fast = fast.Next
// 		if fast != nil {
// 			fast = fast.Next
// 		}
// 	}
// 	return false
// }

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow := head.Next
    if slow == nil {
        return false
    }

    fast := slow.Next

    for fast != nil && slow != nil {
        if fast == slow {
            return true
        }

        slow = slow.Next
        
        fast = fast.Next
        if fast != nil {
            fast = fast.Next
        }
    }

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
