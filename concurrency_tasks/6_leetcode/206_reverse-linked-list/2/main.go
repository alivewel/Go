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
	// Создание первого списка [1, 2, 3, 4, 5]
	list := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil}}}}}

	// Создание второго списка [1, 2, 3, 5]
	// res := &ListNode{1,
	// 	&ListNode{2,
	// 		&ListNode{3,
	// 			&ListNode{5, nil}}}}

	fmt.Println(reverseList(list)) // Список: 8->9->9->9->0->0->0->1 Число: 10009998
	// fmt.Println(res)
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	prev := &ListNode{}
	for curr != nil {
		tmp := curr
		curr = curr.Next
		tmp.Next = prev
		prev = tmp
	}
	return prev
}

// Можно изменить изначальный список, а можно создать новый и с ним уже играться.
