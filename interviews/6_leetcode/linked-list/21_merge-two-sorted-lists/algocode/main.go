package main

import (
	"fmt"
	// "math"
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

// func getVal(list *ListNode) int {
// 	if list == nil {
// 		return math.MaxInt
// 	}
// 	return list.Val
// }

func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    for head1 != nil && head2 != nil {
        temp := &ListNode{}
        if head1.Val > head2.Val {
            temp.Val = head2.Val
            head2 = head2.Next
        } else {
            temp.Val = head1.Val
            head1 = head1.Next 
        }
		curr.Next = temp
        curr = curr.Next
    }
	if head1 != nil {
		curr.Next = head1
	}
	if head2 != nil {
		curr.Next = head2
	}
    return dummy.Next
}
