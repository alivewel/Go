package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// для решения используется 2 указателя - быстрый и медленный
// есть второе решение, когда в первом проходе мы считаем длину списка
// а во втором проходе мы проходимся и находим середину списка
// асимптотитчески оба способа будут работать за одинаковое время

func middleNodeValue(head *ListNode) int {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        // fast = slow.Next
		fast = fast.Next.Next
        // if fast.Next != nil {
        //     fast = fast.Next
        // }
    }
    return slow.Val
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
	list3 := &ListNode{1, &ListNode{2,nil}}

	fmt.Println(middleNodeValue(list1))
	fmt.Println(middleNodeValue(list2))
	fmt.Println(middleNodeValue(list3))
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
