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

	// Список: 2->4->3 Число:342
	// first := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	// // Список: 5->6->4 Число:465
	// second := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	// zero := &ListNode{0, nil}
	// Список: 9->9->9->9->9->9->9 Число: 9999999
	// third := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	// Список: 9->9->9->9 Число: 9999
	// fourth := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}

	third := &ListNode{9, &ListNode{9, nil}}
	fourth := &ListNode{9, &ListNode{}}

	// fmt.Println(addTwoNumbers(first, second)) // Список: 7->0->8  Число: 342+465=807
	// fmt.Println(addTwoNumbers(zero, zero))    // 0
	fmt.Println(addTwoNumbers(third, fourth)) // Список: 8->9->9->9->0->0->0->1 Число: 10009998
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     dummy := &ListNode{} // фиктивный узел
//     current := dummy
//     carry := 0

//     for l1 != nil || l2 != nil || carry != 0 {
//         val1 := 0
//         if l1 != nil {
//             val1 = l1.Val
//             l1 = l1.Next
//         }

//         val2 := 0
//         if l2 != nil {
//             val2 = l2.Val
//             l2 = l2.Next
//         }

//         sum := val1 + val2 + carry
//         carry = sum / 10

//         current.Next = &ListNode{Val: sum % 10}
//         current = current.Next
//     }

//     return dummy.Next
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    carry := 0
    for l1 != nil && l2 != nil {
        sum := l1.Val + l2.Val + carry
        if sum > 9 {
            carry = sum / 10
        } else {
            carry = 0
        }
        // temp := &ListNode{Val: sum % 10}
        head.Next = &ListNode{Val: sum % 10}
        // head.Val = sum % 10
        head = head.Next
        if l1 != nil { //
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }
    return head.Next
}
